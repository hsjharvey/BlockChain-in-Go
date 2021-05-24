package main

import (
	"crypto/sha256"
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

	h := sha256.New()
	h.Write([]byte(hashString))
	BLK.hash = fmt.Sprintf("%v", h.Sum(nil))
	fmt.Printf("%v", BLK)
}
