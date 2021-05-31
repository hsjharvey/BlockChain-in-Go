package main

import (
	"os"
	"time"
)

func checkError(err error) {
	if err != nil {
		//fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
func getCurrentUnixTime() string {
	return string(time.Now().Unix())
}
