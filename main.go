package main // import "golang"

import (
	"fmt"
)

type Configs struct {
	blockChain          []Block
	pendingTransactions []Transaction
}

func main() {
	pendingTransactions := make(chan []Transaction)
	fmt.Println(pendingTransactions)
	fmt.Println("test end")
}
