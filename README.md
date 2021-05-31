# A mini blockchain example implemented in Go

## Description

- This is a synchronous mini example of an ethereum based blockchain system.
- See ```block_chain.json``` for the chain output.

<br>

## Requirement

- See ```go.mod```.
- GCC is required.

<br>

## How-to

- Windows: terminal run ```go run ./src/.```
- Mac: terminal run ```go run ./src/*.go```

<br>

## Output
-  ```blockchain.json```
- terminal output:
```
2021/05/31 17:10:41.544364 User and miner accounts generated
2021/05/31 17:10:41.544545 Genesis block generated
2021/05/31 17:10:41.549445 New Block initialized
2021/05/31 17:10:41.549689 Finish picking up transactions and complete verification process
2021/05/31 17:10:41.549695 Start mining
2021/05/31 17:10:41.549701 Target problem to be matched in mining: 01
2021/05/31 17:10:41.567613 New Block (01K40da6V03/7a/KcOHgeYTsVWVljIY2uO8EntiMHxU=) mined by (0xa3b2b0067e4B4b5d1f31afA6c7412e9FF5845743)
2021/05/31 17:10:41.568493 Coinbase balance:  3.141592653589793e+25
2021/05/31 17:10:41.568508 0xE7e24376Ed2C3E265032E939c4597A0f1c4756D9 balance:  4864
2021/05/31 17:10:41.568511 0x05E04EFad772eeA6AC5b09e2a31ed7bD8CCCE1F4 balance:  125
2021/05/31 17:10:41.568514 0xa3b2b0067e4B4b5d1f31afA6c7412e9FF5845743  (miner) balance:  61

```

## Reference:

- [Nonce and difficulty](https://medium.com/verifyas/what-you-should-know-about-nonces-and-difficulty-8c4ce499a766)
