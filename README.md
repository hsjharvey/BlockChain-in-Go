<h1 align="center">
 A mini blockchain example
</h1>

## Description

- This is a Synchronous mini example of an Ethereum based blockchain system, implemented in Go.
- See the [wiki page](https://github.com/hsjharvey/BlockChain-in-Go/wiki/A-not-so-technical-explanation-on-BlockChain-system) of this project for a detailed guide on blockchain system.
<p align="center">
    <img src="https://github.com/hsjharvey/Presentations/blob/master/2021_blockchain_illustration.png">
</p>


## Requirement

- See ```go.mod```.
- GCC is required.


## How-to-use

- Terminal run ```go mod tidy```.
- Windows: terminal run ```go run ./src/.```
- Mac: terminal run ```go run ./src/*.go```.

## Output

- [```blockchain.json```](./blockchain.json)
- terminal output:

```
2021/06/01 15:15:28.171236 User and miner accounts generated
2021/06/01 15:15:28.172268 Genesis block generated
2021/06/01 15:15:28.179805 New block initialized
2021/06/01 15:15:28.179843 Finish picking up transactions and complete verification process
2021/06/01 15:15:28.179843 Start mining
2021/06/01 15:15:28.180382 Target hash header to be matched in mining: 01
2021/06/01 15:15:28.185864 New block (01ARr0wI68FG142A8L7e1Ej6uPHUBn1rUm6YAzYQoX8=) mined by (0x65da0e4e0d08Bf39846b88b49df10662c870143e)
2021/06/01 15:15:28.186425 ----------------- Balance check -----------------
2021/06/01 15:15:28.186425 Coinbase:  3.141592653589793e+25
2021/06/01 15:15:28.186425 0x6273D800918fCDbD9729F75c069b9fBB4D410193 :  4864
2021/06/01 15:15:28.186928 0xEfd0AB09a4A5d5241615b5E6e15493C064a3E161 :  125
2021/06/01 15:15:28.186956 0x65da0e4e0d08Bf39846b88b49df10662c870143e  (miner):  61
2021/06/01 15:15:28.186956 -------------------------------------------------
```
