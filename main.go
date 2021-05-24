package main // import "golang"

import (
	"crypto/sha256"
	"fmt"
)

type Configs struct {
	blockChain          []Block
	pendingTransactions []Transaction
}

type BlockChain struct {
	chain               []Block
	pendingTransactions []Transaction
}

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

type Block struct {
	index                   int
	previousBlockID         string
	time                    string
	hash                    string
	selectedTransactionList []Transaction
	nonce                   int
}

func mineNewBlock(difficulty int) {

}

func (T *Transaction) transactionHashCalculation() {
	hashString := T.senderID + T.receiverID + fmt.Sprintf("%f", T.amount) + T.time
	h := sha256.New()
	h.Write([]byte(hashString))
	T.hash = fmt.Sprintf("%v", h.Sum(nil))
	fmt.Printf("%x", T)
}

func (BC *Block) blockHashCalculation() {
	hashT := ""
	for _, eachT := range BC.selectedTransactionList {
		hashT += eachT.hash
	}

	hashString := BC.time + hashT + BC.previousBlockID + fmt.Sprintf("%v", BC.index)

	h := sha256.New()
	h.Write([]byte(hashString))
	BC.hash = fmt.Sprintf("%v", h.Sum(nil))
	fmt.Printf("%v", BC)

}

func main() {

	fmt.Println("test end")
}
