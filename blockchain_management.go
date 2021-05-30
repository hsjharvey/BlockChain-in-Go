package main

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"
)

type BlockChain struct {
	Chain []Block `json:"Chain"`
}

type Block struct {
	Index                   int           `json:"Index"`
	PreviousBlockAddress    string        `json:"PreviousBlockAddress"`
	TimeStamp               string        `json:"TimeStamp"`
	SelectedTransactionList []Transaction `json:"SelectedTransactionList"`
	MinerAddress            string        `json:"MinerAddress"`
	Nonce                   int           `json:"Nonce"`
	Hash                    string        `json:"HashString"`
}

func InitializeNewBlock(BC *BlockChain) Block {
	lastBlock := BC.Chain[len(BC.Chain)-1]
	newBLK := Block{}
	newBLK.Index = lastBlock.Index + 1
	newBLK.PreviousBlockAddress = lastBlock.Hash

	fmt.Println("New Block initialized")
	return newBLK
}
func Mining(minerHashID string, BLK Block, BC *BlockChain, difficulty int) {
	// determine difficulty level
	var diff []string
	BLK.Nonce = 0

	for i := 0; i <= difficulty; i++ {
		diff = append(diff, strconv.Itoa(i))
	}
	problemToBeSolved := strings.Join(diff[:], ",")

	// solve the problem
	for {
		BLK.TimeStamp = getCurrentUnixTime()

		if BLK.Hash[0:difficulty] != problemToBeSolved {
			BLK.Nonce += 1
			BLK.BlockHashCalculation()
		} else {
			BLK.MinerAddress = minerHashID
			CoinBaseTransaction(minerHashID, &BLK)

			BC.Chain = append(BC.Chain, BLK)

			fmt.Println("-------------------------")
			fmt.Println("New Block " + BLK.Hash + " mined by " + BLK.MinerAddress)
			fmt.Println("-------------------------")
			break
		}
	}
}

func (BLK *Block) BlockHashCalculation() {
	hashTx := ""
	for _, eachTx := range BLK.SelectedTransactionList {
		hashTx += eachTx.HashString
	}

	hashString := BLK.TimeStamp + hashTx + BLK.PreviousBlockAddress + fmt.Sprintf("%v", BLK.Nonce)
	h := sha256.Sum256([]byte(hashString))
	BLK.Hash = base64.StdEncoding.EncodeToString(h[:])
}

func CreateGenesisBlock(genesisUser string, BC *BlockChain) {
	createCoinbaseAccount()
	var TList []Transaction
	genesisT := InitTransaction("coinbase", genesisUser, 5000.0, 0.0,
		"hello world")

	genesisT.Signature = []byte("signed by God")
	genesisT.Accepted = true

	TList = append(TList, genesisT)

	b := Block{
		Index:                   0,
		PreviousBlockAddress:    "BigBang",
		TimeStamp:               getCurrentUnixTime(),
		SelectedTransactionList: TList,
		Nonce:                   1992,
		MinerAddress:            "Harvey",
	}

	b.BlockHashCalculation()

	BC.Chain = append(BC.Chain, b)

	genesisT.updateAccount()

	fmt.Println("-------------------------")
	fmt.Println("Genesis block generated")
	fmt.Println("-------------------------")
}
