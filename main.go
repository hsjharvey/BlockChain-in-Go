package main // import "golang"
import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// initialize a blockchain
	//pendingTransactions := MEMPool{}
	BC := BlockChain{}

	// create users
	bmm1Address := CreateNewAccount("bmm1")
	bmm2Address := CreateNewAccount("bmm2")
	fmt.Println("user accounts generated")

	CreateGenesisBlock(bmm1Address, &BC)

	// create a transaction
	var TxList []Transaction
	genesisTx := InitTransaction(bmm1Address, bmm2Address, 100.0, 10.0,
		"Brain, Mind and Markets 2021")
	TxList = append(TxList, genesisTx)

	// sign transactions

	// send transaction to pending transactions memory pool (called MEMPOOL)

	// create new miner

	// miner intialize a new

	// miner pick up and verify the validity of the transactions

	// mining (solve the math problem)

	// (if mining successful and no new block in the blockchain, requires async receiving msg from the network)
	// broadcast to everyone in the network

	file, _ := json.MarshalIndent(&BC, "", " ")
	_ = ioutil.WriteFile("BlockChain.json", file, os.ModePerm)
}
