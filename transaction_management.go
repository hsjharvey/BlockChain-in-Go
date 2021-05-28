package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"sort"
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

func getPrivateKey(fileLocation string) (*rsa.PrivateKey, error) {
	key, err := ioutil.ReadFile(fileLocation)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(key)
	der, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return der, err
}

func (T *Transaction) SignTransaction(privateKeyLocation string) {
	privateKey, _ := getPrivateKey(privateKeyLocation)
	signature, _ := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, []byte(T.Hash))
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
	sort.Slice(MP.pendingTransactions, func(i, j int) bool {
		return MP.pendingTransactions[i].Fee > MP.pendingTransactions[j].Fee
	})

	idx := 0
	for {
		if len(NB.SelectedTransactionList) < 2 {
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
	T.Accepted = true
}
