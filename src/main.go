package main // import "golang"
import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	// configuration
	config()
	checkFolderExist()

	//initialize a blockchain
	BC := BlockChain{}
	MP := MEMPool{}

	// create users and miner
	bmm1Address := CreateNewAccount("bmm1")
	bmm2Address := CreateNewAccount("bmm2")
	miner1Address := CreateNewAccount("miner1")
	log.Println("User and miner accounts generated")

	CreateGenesisBlock(bmm1Address, &BC)

	// create transactions and send to the pending transactions in the memory pool
	Tx1 := InitTransaction(bmm1Address, bmm2Address, 100.0, 10.0,
		"Brain, Mind and Markets 2021")
	// sign the transaction
	Tx1.SignTransaction("./accounts/private_bmm1.pem")
	Tx1.SendTransactionToPool(&MP)

	Tx2 := InitTransaction(bmm1Address, bmm2Address, 25.0, 1.0,
		"Bazzinga")
	Tx2.SignTransaction("./accounts/private_bmm1.pem")
	Tx2.SendTransactionToPool(&MP)

	Tx3 := InitTransaction(bmm2Address, bmm1Address, 17.0, 5.0,
		"shadowlands")
	Tx3.SignTransaction("./accounts/private_bmm2.pem")
	Tx3.SendTransactionToPool(&MP)

	// miner initialize a new block
	newBlock := InitializeNewBlock(&BC)

	// miner pick and verify the validity of the transactions
	PickTxAndVerifyValidity(&newBlock, MP)
	log.Println("Finish picking up transactions and complete verification process")

	//mining (solve the math problem)
	log.Println("Start mining")
	minedBlock := Mining(miner1Address, newBlock, &BC, 2)

	MP.updateMEMPool(minedBlock.SelectedTransactionList)
	minedBlock.blockConfirmationUpdateSystemAccount()

	//(if mining successful and no new block in the blockchain, requires async receiving msg from the network)
	//broadcast to everyone in the network

	// save the blockchain in a
	file, _ := json.MarshalIndent(&BC, "", " ")
	_ = ioutil.WriteFile("./blockchain.json", file, os.ModePerm)

	log.Println("Coinbase balance: ", BalanceCheck("coinbase"))
	log.Println(bmm1Address, "balance: ", BalanceCheck(bmm1Address))
	log.Println(bmm2Address, "balance: ", BalanceCheck(bmm2Address))
	log.Println(miner1Address, " (miner) balance: ", BalanceCheck(miner1Address))
}
