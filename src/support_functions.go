package main

import (
	"log"
	"os"
	"path/filepath"
	"time"
)

func checkError(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}
func getCurrentUnixTime() int64 {
	return time.Now().UnixNano()
}

func config() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
}

func checkFolderExist() {
	np := filepath.Join(".", "accounts")
	err := os.MkdirAll(np, os.ModePerm)
	checkError(err)
}
