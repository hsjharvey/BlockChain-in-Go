package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/gob"
	"encoding/pem"
	"fmt"
	"os"
)

func CreateNewAccount(accountID string) string {
	reader := rand.Reader
	bitSize := 2048

	key, err := rsa.GenerateKey(reader, bitSize)
	checkError(err)

	publicKey := key.PublicKey

	saveGobKey("./accounts/private_"+accountID+".key", key)
	savePEMKey("./accounts/private_"+accountID+".pem", key)

	saveGobKey("./accounts/public_"+accountID+".key", publicKey)
	savePublicPEMKey("./accounts/public_"+accountID+".pem", &publicKey)

	accountAddress := hashAndSaveAddress("./accounts/sha256_base64_encoded_address_"+accountID+".txt", &publicKey)
	return accountAddress
}

func hashAndSaveAddress(fileName string, pubKey *rsa.PublicKey) string {
	outFile, err := os.Create(fileName)
	checkError(err)
	defer outFile.Close()

	h := sha256.Sum256(x509.MarshalPKCS1PublicKey(pubKey))
	hashString := base64.StdEncoding.EncodeToString(h[:])
	_, err = outFile.WriteString(hashString)
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
	checkError(err)

	return hashString
}

func saveGobKey(fileName string, key interface{}) {
	outFile, err := os.Create(fileName)
	checkError(err)
	defer outFile.Close()

	encoder := gob.NewEncoder(outFile)
	err = encoder.Encode(key)
	checkError(err)
}

func savePEMKey(fileName string, key *rsa.PrivateKey) {
	outFile, err := os.Create(fileName)
	checkError(err)
	defer outFile.Close()

	var privateKey = &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(key),
	}

	err = pem.Encode(outFile, privateKey)
	checkError(err)
}

func savePublicPEMKey(fileName string, key *rsa.PublicKey) {
	var pemkey = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: x509.MarshalPKCS1PublicKey(key),
	}

	pemfile, err := os.Create(fileName)
	checkError(err)
	defer pemfile.Close()

	err = pem.Encode(pemfile, pemkey)
	checkError(err)
}
