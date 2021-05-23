package main // import "golang"

import (
	"crypto/sha256"
	"fmt"
)

type Transaction struct {
	senderID          string
	receiverID        string
	transactionAmount float64
	transactionFee    float64
	time              string
	hash              string
	message           string
}

type Block struct {
	index                   int
	previousBlock           string
	time                    string
	hash                    string
	selectedTransactionList []Transaction
}

var blockChain []Block
var pendingTransactions []Transaction

func AddNewBlocks() {

}

func (T *Transaction) HashCalculation() {
	hashString := T.senderID + T.receiverID + fmt.Sprintf("%f", T.transactionAmount) + T.time
	h := sha256.New()
	h.Write([]byte(hashString))
	fmt.Printf("%x", h.Sum(nil))


}

func main() {

	fmt.Println("test end")
}
