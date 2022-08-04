package main

import (
	"crypto/sha256"
	"fmt"
	"time"
)

type Block struct {
	Index     int
	Timestamp time.Time
	BPM       int
	Hash      string
	PrevHash  string
}

var BlockChain []Block

func createHash(block Block) [32]byte {
	record := fmt.Sprintf("%d%s%d%s", block.Index, block.Timestamp.String(), block.BPM, block.PrevHash)

	hashed := sha256.Sum256([]byte(record))

	return hashed
}
