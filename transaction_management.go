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
)

type Transaction struct {
	From      string  `json:"From"`
	To        string  `json:"To"`
	Amount    float64 `json:"Amount"`
	Fee       float64 `json:"Fee"`
	Time      string  `json:"Time"`
	Hash      string  `json:"Hash"`
	Message   string  `json:"Message"`
	Signature []byte  `json:"Signature"`
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

func (MP *MEMPool) SendTransactionToPool(T Transaction) {
	MP.pendingTransactions = append(MP.pendingTransactions, T)
}

func CoinBaseTransaction(minerID string) Transaction {
	T := Transaction{
		From:    "coinbase",
		To:      minerID,
		Amount:  50,
		Fee:     0,
		Time:    getCurrentUnixTime(),
		Message: "reward for block miner",
	}

	T.transactionHashCalculation()
	return T
}

func (MP *MEMPool) PickTxAndVerifyValidity() {

}
