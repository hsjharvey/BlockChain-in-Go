<h1 align="center">
 A mini blockchain example
</h1>

## Description

- This is a mini example of an Ethereum based blockchain system, implemented in Go.
- See the [wiki page](https://github.com/hsjharvey/BlockChain-in-Go/wiki/A-not-so-technical-explanation-on-an-example-BlockChain-system) of this project for a detailed guide on blockchain system.
<p align="center">
    <img src="https://github.com/hsjharvey/Notes-and-Presentations/blob/master/Presentations/2021_blockchain_illustration.png">
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
2021/06/02 12:04:51.007887 User and miner accounts generated
2021/06/02 12:04:51.008406 Genesis block generated
2021/06/02 12:04:51.014111 New block initialized
2021/06/02 12:04:51.014111 Finish picking up transactions and complete verification process
2021/06/02 12:04:51.014111 Start mining
2021/06/02 12:04:51.014660 Target hash header to be matched in mining: 01
2021/06/02 12:04:51.015167 New block (01cdGK6DiobLVeNXi8x9C+PY9+w3j+aL6UEo4cdwkfA=) mined by (0xc218cC68F0B1039BBDFCF19b92671Ed1a54FDbf1)
2021/06/02 12:04:51.015681 ------------------- Balance check -------------------
2021/06/02 12:04:51.015681 Coinbase:  2.9999994939e+10
2021/06/02 12:04:51.016200 0xD701b4EEbDBCD9d89a2AD0bA8cb41d5ecb2d25a7 :  4875
2021/06/02 12:04:51.016200 0xCdB7820b55992FDC9168846bf9B90a5a69AC2E53 :  114
2021/06/02 12:04:51.016714 0xc218cC68F0B1039BBDFCF19b92671Ed1a54FDbf1  (miner):  61
2021/06/02 12:04:51.016714 -----------------------------------------------------
```
