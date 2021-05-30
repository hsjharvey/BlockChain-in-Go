package main

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"io/ioutil"
	"sort"
)

type Transaction struct {
	From       string  `json:"From"`
	To         string  `json:"To"`
	Amount     float64 `json:"Amount"`
	Fee        float64 `json:"Fee"`
	Time       string  `json:"TimeStamp"`
	HashString string  `json:"HashString"`
	Message    string  `json:"Message"`
	Signature  []byte  `json:"Signature"`
	Accepted   bool    `json:"Accepted"`
}

type MEMPool struct {
	pendingTransactions []Transaction
}

func (T *Transaction) transactionHashCalculation() {
	hashString := T.From + T.To + fmt.Sprintf("%f", T.Amount) + T.Time
	h := sha256.Sum256([]byte(hashString))
	T.HashString = base64.StdEncoding.EncodeToString(h[:])
}

func InitTransaction(senderAddress string, receiverAddress string, amount float64, fee float64, message string) Transaction {
	T := Transaction{
		From:    senderAddress,
		To:      receiverAddress,
		Amount:  amount,
		Fee:     fee,
		Time:    getCurrentUnixTime(),
		Message: message,
	}

	T.transactionHashCalculation()
	return T
}

func getPrivateKey(fileLocation string) (*ecdsa.PrivateKey, error) {
	key, err := ioutil.ReadFile(fileLocation)
	checkError(err)

	fileBytes, _ := pem.Decode(key)
	der, err := x509.ParseECPrivateKey(fileBytes.Bytes)
	checkError(err)

	return der, err
}

func (T *Transaction) SignTransaction(privateKeyLocation string) {
	privateKey, err := getPrivateKey(privateKeyLocation)
	checkError(err)

	decodedVal, err := base64.StdEncoding.DecodeString(T.HashString)
	checkError(err)

	signature, err := crypto.Sign(decodedVal, privateKey)
	checkError(err)

	T.Signature = signature
}

func (T Transaction) SendTransactionToPool(MP *MEMPool) {
	MP.pendingTransactions = append(MP.pendingTransactions, T)
}

func CoinBaseTransaction(minerID string, NB *Block) {
	totalGas := 0.0
	for _, eachTx := range NB.SelectedTransactionList {
		totalGas += eachTx.Fee
	}

	T := Transaction{
		From:    "coinbase",
		To:      minerID,
		Amount:  50 + totalGas,
		Fee:     0,
		Time:    getCurrentUnixTime(),
		Message: "reward for block miner",
	}

	T.transactionHashCalculation()
	NB.SelectedTransactionList = append(NB.SelectedTransactionList, T)
}

func PickTxAndVerifyValidity(NB *Block, MP MEMPool) {
	// sort the pending transactions based on fees
	sort.Slice(MP.pendingTransactions, func(i, j int) bool {
		return MP.pendingTransactions[i].Fee > MP.pendingTransactions[j].Fee
	})

	idx := 0
	for {
		if len(NB.SelectedTransactionList) < 2 {
			// check the balance and the validity of the transaction
			MP.pendingTransactions[idx].TxValidityCheck()
			if MP.pendingTransactions[idx].Accepted {
				NB.SelectedTransactionList = append(NB.SelectedTransactionList)
			}
		} else {
			break
		}
	}
}

func (T *Transaction) TxValidityCheck() {
	// Step 1: check whether has enough balance
	Bal := BalanceCheck(T.From)
	if Bal >= T.Fee {
		// Step 2: verify signature
		decodedVal, err := base64.StdEncoding.DecodeString(T.HashString)
		checkError(err)

		pubKey, err := crypto.Ecrecover(decodedVal, T.Signature)
		checkError(err)

		h := sha256.Sum256(pubKey)
		hashString := base64.StdEncoding.EncodeToString(h[:])
		fmt.Println(hashString)
		fmt.Println(T.From)

		if hashString == T.From {
			T.Accepted = true
		} else {
			T.Accepted = false
		}

	} else {
		T.Accepted = false
	}
}

func testSignatureFc() {
	hs := "helloworld"

	data := []byte(hs)
	decodedVal := sha256.Sum256(data)
	de, _ := pem.Decode([]byte(pbk))
	x, err := x509.ParsePKIXPublicKey(de.Bytes)
	readPubKey, _ := x.(*ecdsa.PublicKey)
	checkError(err)
	//fmt.Printf("Got a %T, readPubKey)
	fmt.Println("Public key")
	fmt.Println(readPubKey)
	de, _ = pem.Decode([]byte(prk))
	readPrivateKey, err := x509.ParseECPrivateKey(de.Bytes)

	fmt.Println("Private key")
	fmt.Println(readPrivateKey)

	signature, err := crypto.Sign(decodedVal[:], readPrivateKey)
	checkError(err)

	recoverPubKeyBytes, err := crypto.Ecrecover(decodedVal[:], signature)
	fmt.Println(recoverPubKeyBytes)
	//checkError(err)
	//fmt.Println("public key")
	//fmt.Println(recoverPubKeyBytes)
	//
	//fmt.Println(readPubKey == recoverPubKeyBytes)

	r, s, err := ecdsa.Sign(rand.Reader, readPrivateKey, decodedVal[:])
	if err != nil {
		return
	}
	matches := ecdsa.Verify(readPubKey, decodedVal[:], r, s)
	fmt.Println(matches)
}
