package main

import (
	"crypto/sha256"
	"encoding/base64"
	"log"
	"strconv"
	"strings"
)

type BlockChain struct {
	Chain []Block `json:"Chain"`
}

type Block struct {
	Index                   int           `json:"Index"`
	PreviousBlockAddress    string        `json:"PreviousBlockAddress"`
	TimeStamp               int64         `json:"TimeStamp"`
	MinerAddress            string        `json:"MinerAddress"`
	Nonce                   int           `json:"Nonce"`
	Difficulty              int           `json:"Difficulty"`
	BlockHash               string        `json:"BlockHash"`
	SelectedTransactionList []Transaction `json:"SelectedTransactionList"`
}

func (BC *BlockChain) InitNewBlock(minerID string) Block {
	lastBlock := BC.Chain[len(BC.Chain)-1]
	newBLK := Block{}
	newBLK.Index = lastBlock.Index + 1
	newBLK.PreviousBlockAddress = lastBlock.BlockHash
	newBLK.Nonce = 0
	CoinBaseTransaction(minerID, &newBLK)

	newBLK.BlockHashCalculation()

	log.Println("New block initialized")

	return newBLK
}
func Mining(minerHashID string, BLK Block, BC *BlockChain, difficulty int) Block {
	// determine difficulty level
	var diff []string
	BLK.Difficulty = difficulty

	for i := 0; i < difficulty; i++ {
		diff = append(diff, strconv.Itoa(i))
	}
	problemToBeSolved := strings.Join(diff[:], "")

	log.Println("Target hash header to be matched in mining: " + problemToBeSolved)

	// solve the problem
	for {
		BLK.TimeStamp = getCurrentUnixTime()

		if BLK.BlockHash[0:difficulty] != problemToBeSolved {
			BLK.Nonce += 1
			BLK.BlockHashCalculation()
		} else {
			BLK.MinerAddress = minerHashID
			BC.Chain = append(BC.Chain, BLK)

			log.Println("New block (" + BLK.BlockHash + ") mined by (" + BLK.MinerAddress + ")")
			break
		}
	}
	return BLK
}

func (BLK *Block) BlockHashCalculation() {
	var selectedTxsHash string
	for _, eachTx := range BLK.SelectedTransactionList {
		selectedTxsHash += eachTx.TxHash
	}

	hashString := strconv.FormatInt(BLK.TimeStamp, 10) + selectedTxsHash +
		BLK.PreviousBlockAddress + strconv.Itoa(BLK.Nonce)
	h := sha256.Sum256([]byte(hashString))
	BLK.BlockHash = base64.StdEncoding.EncodeToString(h[:])
}

func createGenesisBlock(genesisUser string, BC *BlockChain) {
	createCoinbaseAccount()
	var TList []Transaction
	genesisT := InitTx("coinbase", genesisUser, 5000.0, 0.0,
		"hello world")

	genesisT.Signature = []byte("signed by God")
	genesisT.Accepted = true

	TList = append(TList, genesisT)

	b := Block{
		Index:                   0,
		PreviousBlockAddress:    "BigBang",
		TimeStamp:               getCurrentUnixTime(),
		SelectedTransactionList: TList,
		Nonce:                   2021,
		MinerAddress:            "Harvey Huang",
	}

	b.BlockHashCalculation()

	BC.Chain = append(BC.Chain, b)

	genesisT.updateAccount()

	log.Println("Genesis block generated")
}
