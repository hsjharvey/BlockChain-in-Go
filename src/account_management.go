package main

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/crypto"
	"os"
)

type AccountInfo struct {
	TxCount      int           `json:"TxCount"` //this is the account nonce
	TotalReceive float64       `json:"TotalReceive"`
	TotalSend    float64       `json:"TotalSend"`
	Transactions []Transaction `json:"Transactions"`
}

var AccountManagement = make(map[string]*AccountInfo)

func CreateNewAccount(accountID string) string {
	privateKey, err := crypto.GenerateKey()
	checkError(err)

	publicKey := privateKey.PublicKey

	savePrivateKey("./accounts/private_"+accountID+".pem", privateKey)

	accountAddress := hashAndSaveAddress("./accounts/encoded_address_"+accountID+".txt", publicKey)

	newAccount := AccountInfo{
		TxCount:      0,
		TotalReceive: 0.0,
		TotalSend:    0.0,
		Transactions: []Transaction{},
	}
	AccountManagement[accountAddress] = &newAccount

	return accountAddress
}

func createCoinbaseAccount() {
	newAccount := AccountInfo{
		TxCount:      0,
		TotalReceive: 31415926535897932384626433,
		TotalSend:    0.0,
		Transactions: []Transaction{},
	}
	AccountManagement["coinbase"] = &newAccount
}

func (BLK Block) blockConfirmationUpdateAccounts() {
	for _, eachT := range BLK.SelectedTransactionList {
		eachT.updateAccount()
	}
}

func (T Transaction) updateAccount() {
	if T.Accepted {
		// update sender
		AccountManagement[T.From].TxCount += 1
		AccountManagement[T.From].TotalSend += T.Amount + T.Fee
		AccountManagement[T.From].Transactions = append(AccountManagement[T.From].Transactions, T)

		// update receiver
		AccountManagement[T.To].TxCount += 1
		AccountManagement[T.To].TotalReceive += T.Amount
		AccountManagement[T.To].Transactions = append(AccountManagement[T.To].Transactions, T)
	}
}

func BalanceCheck(AccountAddress string) float64 {
	return AccountManagement[AccountAddress].TotalReceive - AccountManagement[AccountAddress].TotalSend
}

func hashAndSaveAddress(fileName string, pubKey ecdsa.PublicKey) string {
	outFile, err := os.Create(fileName)
	checkError(err)
	defer outFile.Close()

	hashString := crypto.PubkeyToAddress(pubKey)
	_, err = outFile.WriteString(hashString.String())
	checkError(err)

	return hashString.String()
}

func savePrivateKey(fileName string, key *ecdsa.PrivateKey) {
	outFile, err := os.Create(fileName)
	checkError(err)
	defer outFile.Close()

	err = crypto.SaveECDSA(fileName, key)
	checkError(err)
}
