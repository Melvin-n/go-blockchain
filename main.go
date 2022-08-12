package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/joho/godotenv"
	"github.com/melvin-n/go-blockchain/models"
)

var Blockchain []models.Block

func generateHash(block models.Block) string {
	record := fmt.Sprintf("%d%s%d%s", block.Index, block.Timestamp.String(), block.BPM, block.PrevHash)
	hashed := sha256.Sum256([]byte(record))

	//[:] turns array into slice, and slices have capabilities that arrays don't e.g converting to string
	return hex.EncodeToString(hashed[:])
}

func generateNewBlock(bpm int, prevBlock models.Block) models.Block {
	var newBlock models.Block

	newBlock.Index = prevBlock.Index + 1
	newBlock.Timestamp = time.Now()
	newBlock.BPM = bpm
	newBlock.PrevHash = prevBlock.Hash
	newBlock.Hash = generateHash(newBlock)

	return newBlock
}

func validateBlock(oldBlock, newBlock models.Block) (bool, error) {
	if oldBlock.Index+1 != newBlock.Index {
		err := fmt.Errorf("Block index error")
		return false, err
	}
	if oldBlock.Hash != newBlock.PrevHash {
		err := fmt.Errorf("Hash history error")
		return false, err
	}
	if generateHash(newBlock) != newBlock.Hash {
		err := fmt.Errorf("Hashing error")
		return false, err
	}
	return true, nil
}

//checks for latest version of blockchain, replaces if newer version is available
func refreshChain(newBlocks []models.Block) {
	if len(newBlocks) > len(Blockchain) {
		Blockchain = newBlocks
	}
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		genesisBlock := models.Block{Index: 0, Timestamp: time.Now(), BPM: 0, Hash: "", PrevHash: ""}
		spew.Dump(genesisBlock)
		Blockchain = append(Blockchain, genesisBlock)
	}()
	Run()
}
