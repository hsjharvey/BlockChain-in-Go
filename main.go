package main // import "golang"

func main() {
	testSignatureFc()

	////initialize a blockchain
	//BC := BlockChain{}
	//MP := MEMPool{}
	//
	//// create users and miner
	//bmm1Address := CreateNewAccount("bmm1")
	//bmm2Address := CreateNewAccount("bmm2")
	////miner1Address := CreateNewAccount("miner1")
	//fmt.Println("user and miner accounts generated")
	//
	//CreateGenesisBlock(bmm1Address, &BC)
	//
	//// create transactions and send to the pending transactions in the memory pool
	//Tx1 := InitTransaction(bmm1Address, bmm2Address, 100.0, 10.0,
	//	"Brain, Mind and Markets 2021")
	//// sign the transaction
	//Tx1.SignTransaction("./accounts/private_bmm1.pem")
	//Tx1.SendTransactionToPool(&MP)
	//
	//Tx2 := InitTransaction(bmm2Address, bmm1Address, 25.0, 1.0,
	//	"Bazzinga")
	//Tx2.SignTransaction("./accounts/private_bmm2.pem")
	//Tx2.SendTransactionToPool(&MP)
	//
	//// miner initialize a new block
	//newBlock := InitializeNewBlock(&BC)
	//
	//// miner pick and verify the validity of the transactions
	//PickTxAndVerifyValidity(&newBlock, MP)
	//
	////mining (solve the math problem)
	////Mining(miner1Address, newBlock, &BC, 2)
	//
	////(if mining successful and no new block in the blockchain, requires async receiving msg from the network)
	////broadcast to everyone in the network
	//
	//file, _ := json.MarshalIndent(&BC, "", " ")
	//_ = ioutil.WriteFile("BlockChain.json", file, os.ModePerm)
}
