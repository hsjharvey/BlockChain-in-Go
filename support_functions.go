package main

import (
	"fmt"
	"os"
	"time"
)

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
func getCurrentUnixTime() string {
	return fmt.Sprintf("%v", time.Now().Unix())
}
