package service

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/jason-wj/blockchain-demo/src/github.com/jason/blockchain-demo/dpos/model"
	"time"
)

func GenerateBlock(oldBlock model.Block, BPM int, address string) (model.Block, error) {
	var newBlock model.Block

	t := time.Now()
	newBlock.Index = oldBlock.Index + 1
	newBlock.Timestamp = t.String()
	newBlock.BPM = BPM
	newBlock.PrevHash = oldBlock.Hash
	newBlock.Hash = CalculateBlockHash(newBlock)
	newBlock.Delegate = address
	return newBlock, nil
}

func calculateHash(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

func CalculateBlockHash(block model.Block) string {
	record := string(block.Index) + block.Timestamp + string(block.BPM) + block.PrevHash
	return calculateHash(record)
}

func IsBlockValid(newBlock model.Block, oldBlock model.Block) bool {
	if oldBlock.Index+1 != newBlock.Index {
		return false
	}
	if oldBlock.Hash != newBlock.PrevHash {
		return false
	}
	if calculateBlockHash(newBlock) != newBlock.Hash {
		return false
	}
	return true
}
