package main

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"
)

type BlockChain struct {
	chain []Block
}

type Block struct {
	index                   int
	previousBlockID         string
	time                    string
	hash                    string
	selectedTransactionList []Transaction
	minerID                 string
	nonce                   int
}

func MineNewBlock(minerHashID string, BLK Block, BC *BlockChain, difficulty int) {
	// determine difficulty level
	var diff []string
	for i := 0; i <= difficulty; i++ {
		diff = append(diff, strconv.Itoa(i))
	}
	problemToBeSolved := strings.Join(diff[:], ",")

	// solve the problem
	for {
		if BLK.hash[0:difficulty] != problemToBeSolved {
			BLK.nonce += 1
			BLK.BlockHashCalculation()
		} else {
			BLK.minerID = minerHashID
			coinbaseT := CoinBaseTransaction(minerHashID)
			BLK.selectedTransactionList = append(BLK.selectedTransactionList, coinbaseT)

			BC.chain = append(BC.chain, BLK)

			fmt.Println("New Block " + BLK.hash + " by " + BLK.minerID)
			break
		}
	}
}

func (BLK *Block) BlockHashCalculation() {
	hashT := ""
	for _, eachT := range BLK.selectedTransactionList {
		hashT += eachT.hash
	}

	hashString := BLK.time + hashT + BLK.previousBlockID + fmt.Sprintf("%v", BLK.nonce)
	h := sha256.Sum256([]byte(hashString))
	BLK.hash = base64.StdEncoding.EncodeToString(h[:])
	fmt.Printf("%v", BLK)
}

func CreateGenesisBlock(genesisUser string) {
	var TList []Transaction
	genesisT := InitTransaction("Harvey", genesisUser, 5000.0, 0.0)
	TList = append(TList, genesisT)

	b := Block{
		index:                   0,
		previousBlockID:         "BigBang",
		time:                    getCurrentUnixTime(),
		selectedTransactionList: TList,
		nonce:                   1992,
		minerID:                 "Harvey",
	}

	b.BlockHashCalculation()

	fmt.Println("-------------------------")
	fmt.Println("Genesis block generated")
	fmt.Println(b)
	fmt.Println("-------------------------")

}
