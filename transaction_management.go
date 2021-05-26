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
	senderAddress   string
	receiverAddress string
	amount          float64
	fee             float64
	time            string
	hash            string
	message         string
	signature       []byte
}

type MEMPool struct {
	pendingTransactions []Transaction
}

func (T *Transaction) transactionHashCalculation() {
	hashString := T.senderAddress + T.receiverAddress + fmt.Sprintf("%f", T.amount) + T.time
	h := sha256.Sum256([]byte(hashString))
	T.hash = base64.StdEncoding.EncodeToString(h[:])
	fmt.Printf("%x", T)
}

func InitTransaction(senderAddress string, receiverAddress string, amount float64, fee float64) Transaction {
	T := Transaction{
		senderAddress:   senderAddress,
		receiverAddress: receiverAddress,
		amount:          amount,
		fee:             fee,
		time:            getCurrentUnixTime(),
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

func (T *Transaction) SignTransaction(fileLocation string) {
	privateKey, _ := getPrivateKey(fileLocation)
	signature, _ := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, []byte(T.hash))
	T.signature = signature
}

func (MP *MEMPool) SendTransactionToPool(T Transaction) {
	// Step 1: verify this is a valid transaction

	// Step 2: send the transaction to the Pending MEMPOOL
	MP.pendingTransactions = append(MP.pendingTransactions, T)
}

func CoinBaseTransaction(minerID string) Transaction {
	T := Transaction{
		senderAddress:   "coinbase",
		receiverAddress: minerID,
		amount:          50,
		fee:             0,
		time:            getCurrentUnixTime(),
		message:         "reward for block miner",
	}

	T.transactionHashCalculation()
	return T
}
