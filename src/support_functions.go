package main

import (
	"log"
	"os"
	"path/filepath"
	"time"
)

func checkError(err error) {
	if err != nil {
		//fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
func getCurrentUnixTime() int64 {
	return time.Now().UnixNano()
}

func config() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
}

func checkFolderExist() {
	newpath := filepath.Join(".", "Accounts")
	os.MkdirAll(newpath, os.ModePerm)
}
