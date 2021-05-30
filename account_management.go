package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"os"
)

type AccountInfo struct {
	TxCount      int           `json:"TxCount"`
	TotalReceive float64       `json:"TotalReceive"`
	TotalSend    float64       `json:"TotalSend"`
	Transactions []Transaction `json:"Transactions"`
}

var AccountManagement = make(map[string]*AccountInfo)

func CreateNewAccount(accountID string) string {
	reader := rand.Reader
	p256 := elliptic.P256()

	privateKey, err := ecdsa.GenerateKey(p256, reader)
	checkError(err)

	publicKey := privateKey.PublicKey

	savePrivatePEMKey("./accounts/private_"+accountID+".pem", privateKey)
	savePublicPEMKey("./accounts/public_"+accountID+".pem", &publicKey)

	accountAddress := hashAndSaveAddress("./accounts/sha256_base64_encoded_address_"+accountID+".txt", &publicKey)

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

func (T Transaction) updateAccount() {
	if T.Accepted {
		// update sender
		AccountManagement[T.From].TxCount += 1
		AccountManagement[T.From].TotalSend += T.Amount
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

func hashAndSaveAddress(fileName string, pubKey *ecdsa.PublicKey) string {
	outFile, err := os.Create(fileName)
	checkError(err)
	defer outFile.Close()

	ms, err := x509.MarshalPKIXPublicKey(pubKey)
	checkError(err)

	h := sha256.Sum256(ms)
	hashString := base64.StdEncoding.EncodeToString(h[:])
	_, err = outFile.WriteString(hashString)
	checkError(err)

	return hashString
}

func savePrivatePEMKey(fileName string, key *ecdsa.PrivateKey) {
	outFile, err := os.Create(fileName)
	checkError(err)
	defer outFile.Close()

	ms, err := x509.MarshalECPrivateKey(key)
	checkError(err)

	var privateKey = &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: ms,
	}

	err = pem.Encode(outFile, privateKey)
	checkError(err)
}

func savePublicPEMKey(fileName string, key *ecdsa.PublicKey) {
	outFile, err := os.Create(fileName)
	checkError(err)
	defer outFile.Close()

	ms, err := x509.MarshalPKIXPublicKey(key)
	checkError(err)

	var publicKey = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: ms,
	}

	err = pem.Encode(outFile, publicKey)
	checkError(err)
}
