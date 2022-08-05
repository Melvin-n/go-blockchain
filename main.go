package main

import (
	"crypto/sha256"
	"encoding/hex"
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

func createHash(block Block) string {
	record := fmt.Sprintf("%d%s%d%s", block.Index, block.Timestamp.String(), block.BPM, block.PrevHash)
	hashed := sha256.Sum256([]byte(record))

	//[:] turns array into slice, and slices have capabilities that arrays don't e.g converting to string
	return hex.EncodeToString(hashed[:])
}

func generateNewBlock(timestamp time.Time, bpm int, prevHash string, prevBlock Block) {
	var newBlock Block

	newBlock.Index = prevBlock.Index + 1
	newBlock.Timestamp = time.Now()
	newBlock.BPM = bpm
	newBlock.PrevHash = prevBlock.Hash
	newBlock.Hash = createHash(newBlock)
}
