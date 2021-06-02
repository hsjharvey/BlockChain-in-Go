package main

import (
	"log"
	"os"
	"path/filepath"
	"time"
)

var (
	accountMgt = make(map[string]*AccountInfo)
)

const (
	MiningDifficultyLv  int     = 2
	TxLimitPerBlock     int     = 2
	CoinBaseInitBalance float64 = 3e10
	MiningReward        float64 = 50
)

func configLog() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}
func getCurrentUnixTime() int64 {
	return time.Now().UnixNano()
}

func checkAccountFolderExist() {
	np := filepath.Join(".", "accounts")
	err := os.MkdirAll(np, os.ModePerm)
	checkError(err)
}
