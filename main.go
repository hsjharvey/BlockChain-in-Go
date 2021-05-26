package main // import "golang"

import "fmt"

type Configs struct {
	blockChain          []Block
	pendingTransactions []Transaction
}

func main() {
	// create users
	bmm1Address := CreateNewAccount("bmm1")
	bmm2Address := CreateNewAccount("bmm2")
	fmt.Println("user accounts generated")
	fmt.Println(bmm1Address)
	fmt.Println(bmm2Address)
	CreateGenesisBlock(bmm1Address)

	// create new transaction

	// sign transactions

	// send transaction to pending transactions memory pool (called MEMPOOL)

	// create new miner

	// mining

	// validate the transaction and broadcast to everyone in the network

	// pendingTransactions := make(chan []Transaction)
}
