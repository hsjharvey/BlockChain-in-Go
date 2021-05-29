package main

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"sort"

	"github.com/ethereum/go-ethereum/crypto"
)

type Transaction struct {
	From      string  `json:"From"`
	To        string  `json:"To"`
	Amount    float64 `json:"Amount"`
	Fee       float64 `json:"Fee"`
	Time      string  `json:"TimeStamp"`
	Hash      string  `json:"Hash"`
	Message   string  `json:"Message"`
	Signature []byte  `json:"Signature"`
	Accepted  bool    `json:"Accepted"`
}

type MEMPool struct {
	pendingTransactions []Transaction
}

func (T *Transaction) transactionHashCalculation() {
	hashString := T.From + T.To + fmt.Sprintf("%f", T.Amount) + T.Time
	h := sha256.Sum256([]byte(hashString))
	T.Hash = base64.StdEncoding.EncodeToString(h[:])
}

func InitTransaction(senderAddress string, receiverAddress string, amount float64, fee float64, message string) Transaction {
	T := Transaction{
		From:    senderAddress,
		To:      receiverAddress,
		Amount:  amount,
		Fee:     fee,
		Time:    getCurrentUnixTime(),
		Message: message,
	}

	T.transactionHashCalculation()
	return T
}

func getPrivateKey(fileLocation string) (*ecdsa.PrivateKey, error) {
	key, err := ioutil.ReadFile(fileLocation)
	checkError(err)

	block, _ := pem.Decode(key)
	der, err := x509.ParseECPrivateKey(block.Bytes)
	checkError(err)

	return der, err
}

func (T *Transaction) SignTransaction(privateKeyLocation string) {
	privateKey, _ := getPrivateKey(privateKeyLocation)

	signature, err := crypto.Sign([]byte(T.Hash), privateKey)
	checkError(err)

	T.Signature = signature
}

func (T Transaction) SendTransactionToPool(MP *MEMPool) {
	MP.pendingTransactions = append(MP.pendingTransactions, T)
}

func CoinBaseTransaction(minerID string, NB *Block) {
	totalGas := 0.0
	for _, eachTx := range NB.SelectedTransactionList {
		totalGas += eachTx.Fee
	}

	T := Transaction{
		From:    "coinbase",
		To:      minerID,
		Amount:  50 + totalGas,
		Fee:     0,
		Time:    getCurrentUnixTime(),
		Message: "reward for block miner",
	}

	T.transactionHashCalculation()
	NB.SelectedTransactionList = append(NB.SelectedTransactionList, T)
}

func PickTxAndVerifyValidity(NB *Block, MP MEMPool) {
	// sort the pending transactions based on fees
	sort.Slice(MP.pendingTransactions, func(i, j int) bool {
		return MP.pendingTransactions[i].Fee > MP.pendingTransactions[j].Fee
	})

	idx := 0
	for {
		if len(NB.SelectedTransactionList) < 2 {
			// check the balance and the validity of the transaction
			MP.pendingTransactions[idx].TxValidityCheck()
			if MP.pendingTransactions[idx].Accepted {
				NB.SelectedTransactionList = append(NB.SelectedTransactionList)
			}
		} else {
			break
		}
	}
}

func (T *Transaction) TxValidityCheck() {
	// Step 1: check whether has enough balance
	Bal := BalanceCheck(T.From)
	if Bal >= T.Fee {
		// Step 2: verify signature
		sigPublicKey, err := crypto.Ecrecover([]byte(T.Hash), T.Signature)
		checkError(err)

		ms, err := x509.MarshalPKIXPublicKey(sigPublicKey)
		checkError(err)

		h := sha256.Sum256(ms)
		hashString := base64.StdEncoding.EncodeToString(h[:])

		if hashString == T.From {
			T.Accepted = true
		} else {
			T.Accepted = false
		}

	} else {
		T.Accepted = false
	}
}
