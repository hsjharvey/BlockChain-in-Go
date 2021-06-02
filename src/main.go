package main // import "golang"
import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	// configuration
	configLog()
	checkAccountFolderExist()

	//initialize a blockchain
	BC := BlockChain{}
	MP := MEMPool{}

	// create users and miner
	bmm1Address := CreateNewAccount("bmm1")
	bmm2Address := CreateNewAccount("bmm2")
	miner1Address := CreateNewAccount("miner1")
	log.Println("User and miner accounts generated")

	createGenesisBlock(bmm1Address, &BC)

	// create transactions and send to the pending transactions in the memory pool
	Tx1 := InitTx(bmm1Address, bmm2Address, 100.0, 10.0,
		"Brain, Mind and Markets 2021")
	// sign the transaction
	Tx1.SignTx("./accounts/private_bmm1.pem")
	Tx1.SendTxtoMEMPool(&MP)

	Tx2 := InitTx(bmm1Address, bmm2Address, 25.0, 1.0,
		"Bazzinga")
	Tx2.SignTx("./accounts/private_bmm1.pem")
	Tx2.SendTxtoMEMPool(&MP)

	// Note: this one should be rejected (Accept=false) as bmm2 does not have enough balance
	Tx3 := InitTx(bmm2Address, bmm1Address, 17.0, 5.0,
		"shadowlands")
	Tx3.SignTx("./accounts/private_bmm2.pem")
	Tx3.SendTxtoMEMPool(&MP)

	// miner initialize a new block
	newBlock := BC.InitNewBlock(miner1Address)

	// miner pick and verify the validity of the transactions
	PickTxAndVerifyValidity(MP, &newBlock)
	log.Println("Finish picking up transactions and complete verification process")

	//mining (solve the math problem)
	log.Println("Start mining")
	minedBlock := Mining(miner1Address, newBlock, &BC, MiningDifficultyLv)

	MP.updateMEMPool(minedBlock.SelectedTransactionList)
	minedBlock.blockConfirmationUpdateAccounts()

	//(if mining successful and no new block in the blockchain, requires async receiving msg from the network)
	//broadcast to everyone in the network

	// save the blockchain in json file
	file, _ := json.MarshalIndent(&BC, "", "  ")
	_ = ioutil.WriteFile("./blockchain.json", file, os.ModePerm)

	log.Println("------------------- Balance check -------------------")
	log.Println("Coinbase: ", BalanceCheck("coinbase"))
	log.Println(bmm1Address, ": ", BalanceCheck(bmm1Address))
	log.Println(bmm2Address, ": ", BalanceCheck(bmm2Address))
	log.Println(miner1Address, " (miner): ", BalanceCheck(miner1Address))
	log.Println("-----------------------------------------------------")
}
