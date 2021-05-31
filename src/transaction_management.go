package main

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"io/ioutil"
	"sort"
	"strconv"
)

type Transaction struct {
	From       string  `json:"From"`
	To         string  `json:"To"`
	Amount     float64 `json:"Amount"`
	Fee        float64 `json:"Fee"`
	TimeStamp  int64   `json:"TimeStamp"`
	HashString string  `json:"HashString"`
	Message    string  `json:"Message"`
	Signature  []byte  `json:"Signature"`
	Accepted   bool    `json:"Accepted"`
}

type MEMPool struct {
	pendingTransactions []Transaction
}

func (MP *MEMPool) updateMEMPool(Ts []Transaction) {
	for _, Tval := range Ts {
		for j, pendingVal := range MP.pendingTransactions {
			if pendingVal.HashString == Tval.HashString {
				MP.pendingTransactions = append(MP.pendingTransactions[0:j], MP.pendingTransactions[j+1:]...)
			}
		}
	}
}

func (T *Transaction) txHashCalculation() {
	hashString := T.From + T.To + fmt.Sprintf("%f", T.Amount) + strconv.FormatInt(T.TimeStamp, 10)
	h := sha256.Sum256([]byte(hashString))
	T.HashString = base64.StdEncoding.EncodeToString(h[:])
}

func InitTx(senderAddress string, receiverAddress string, amount float64, fee float64, message string) Transaction {
	T := Transaction{
		From:      senderAddress,
		To:        receiverAddress,
		Amount:    amount,
		Fee:       fee,
		TimeStamp: getCurrentUnixTime(),
		Message:   message,
	}

	T.txHashCalculation()
	return T
}

func getPrivateKey(fileLocation string) (*ecdsa.PrivateKey, error) {
	key, err := ioutil.ReadFile(fileLocation)
	checkError(err)

	key1, err := crypto.HexToECDSA(string(key))
	checkError(err)

	return key1, err
}

func (T *Transaction) SignTx(privateKeyLocation string) {
	privateKey, err := getPrivateKey(privateKeyLocation)
	checkError(err)

	decodedVal := crypto.Keccak256([]byte(T.HashString))
	checkError(err)

	signature, err := crypto.Sign(decodedVal, privateKey)
	checkError(err)

	T.Signature = signature
}

func (T Transaction) SendTxtoMEMPool(MP *MEMPool) {
	MP.pendingTransactions = append(MP.pendingTransactions, T)
}

func CoinBaseTransaction(minerID string, NB *Block) {
	totalGas := 0.0
	for _, eachTx := range NB.SelectedTransactionList {
		totalGas += eachTx.Fee
	}

	T := Transaction{
		From:      "coinbase",
		To:        minerID,
		Amount:    50 + totalGas,
		Fee:       0,
		TimeStamp: getCurrentUnixTime(),
		Signature: []byte("signed by God"),
		Message:   "reward for block miner",
		Accepted:  true,
	}

	T.txHashCalculation()
	NB.SelectedTransactionList = append(NB.SelectedTransactionList, T)
}

func PickTxAndVerifyValidity(NB *Block, MP MEMPool) {
	// sort the pending transactions based on fees
	sort.Slice(MP.pendingTransactions, func(i, j int) bool {
		return MP.pendingTransactions[i].Fee > MP.pendingTransactions[j].Fee
	})

	for _, eachT := range MP.pendingTransactions {
		if len(NB.SelectedTransactionList) < 2 {
			// check the balance and the validity of the transaction
			eachT.TxValidityCheck()

			if eachT.Accepted {
				// add this valid transaction to the new block
				NB.SelectedTransactionList = append(NB.SelectedTransactionList, eachT)
			}
		} else {
			break
		}
	}
}

func (T *Transaction) TxValidityCheck() {
	// Step 1: check whether has enough balance
	Bal := BalanceCheck(T.From)
	if Bal >= T.Amount+T.Fee {
		// Step 2: verify signature
		recoveredPub, err := crypto.SigToPub(crypto.Keccak256([]byte(T.HashString)), T.Signature)
		checkError(err)
		recoveredAddr := crypto.PubkeyToAddress(*recoveredPub)

		if recoveredAddr.String() == T.From {
			T.Accepted = true
		} else {
			T.Accepted = false
		}
	} else {
		T.Accepted = false
	}
}
