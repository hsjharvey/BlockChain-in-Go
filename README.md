<h1 align="center">
 A mini blockchain example
</h1>

## Description

- This is a synchronous mini example of an ethereum based blockchain system, implemented in Go.

<br>

## Requirement

- See ```go.mod```.
- GCC is required.

<br>

## How-to-use

- Terminal run ```go mod tidy```
- Windows: terminal run ```go run ./src/.```
- Mac: terminal run ```go run ./src/*.go```

<br>

## Output

- ```blockchain.json```
- terminal output:

```
2021/05/31 19:39:18.161889 User and miner accounts generated
2021/05/31 19:39:18.162406 Genesis block generated
2021/05/31 19:39:18.169130 New block initialized
2021/05/31 19:39:18.169683 Finish picking up transactions and complete verification process
2021/05/31 19:39:18.169683 Start mining
2021/05/31 19:39:18.170199 Target hash header to be matched in mining: 01
2021/05/31 19:39:18.171776 New block (01Cb/eOCmKbYCnY9S9Ws/wtxlgu69Q25QBvDHSXMucc=) mined by (0xBB302989c963e8c6451399CAb4102046f2334bfd)
2021/05/31 19:39:18.172303 ----------------- Balance check -----------------
2021/05/31 19:39:18.172303 Coinbase:  3.141592653589793e+25
2021/05/31 19:39:18.172303 0x43A128525DAD70133ffDBBa7CB1F1657f8dc8d11 :  4864
2021/05/31 19:39:18.172831 0xd6c1dE015F764a33fdD07c160d01c40C4BD1504D :  125
2021/05/31 19:39:18.172831 0xBB302989c963e8c6451399CAb4102046f2334bfd  (miner):  61
2021/05/31 19:39:18.173361 -------------------------------------------------
```

## Reference:

- [Nonce and difficulty](https://medium.com/verifyas/what-you-should-know-about-nonces-and-difficulty-8c4ce499a766)
