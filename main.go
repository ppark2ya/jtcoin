package main

import (
	"crypto/sha256"
	"fmt"
)

type block struct {
	data string
	hash string
	prevHash string
}

// B1
// 	b1Hash = fn_hash(data + "")
// B2
// 	b2Hash = fn_hash(data + b1Hash)
// B3
// 	b3Hash = fn_hash(data + b2Hash)

func main() {
	genesisBlock := block{data: "Genesis Block", hash: "", prevHash: ""}
	hash := sha256.Sum256([]byte(genesisBlock.data + genesisBlock.prevHash))
	hexHash := fmt.Sprintf("%x", hash) // 16진수. base16
	genesisBlock.hash = hexHash

	fmt.Println(genesisBlock)
}
