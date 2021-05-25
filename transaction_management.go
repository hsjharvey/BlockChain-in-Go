package main

import (
	"crypto/sha256"
	"fmt"
	"time"
)

type Transaction struct {
	senderID   string
	receiverID string
	amount     float64
	fee        float64
	time       string
	hash       string
	message    string
	signature  string
}

func (T *Transaction) transactionHashCalculation() {
	hashString := T.senderID + T.receiverID + fmt.Sprintf("%f", T.amount) + T.time
	h := sha256.New()
	h.Write([]byte(hashString))
	T.hash = fmt.Sprintf("%v", h.Sum(nil))
	fmt.Printf("%x", T)
}

func InitTransaction(senderID string, receiverID string, amount float64, fee float64) Transaction {
	T := Transaction{
		senderID:   senderID,
		receiverID: receiverID,
		amount:     amount,
		fee:        fee,
		time:       fmt.Sprintf("%v", time.Now().Unix()),
	}

	T.transactionHashCalculation()
	return T
}
