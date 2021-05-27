package main

type AccountInfo struct {
	hashID           string
	noOfTransactions int
	totalReceive     float64
	totalSend        float64
	balance          float64
	transactions     []Transaction
}
