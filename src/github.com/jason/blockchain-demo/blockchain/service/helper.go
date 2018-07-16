package service

import (
	"github.com/jason/blockchain-demo/blockchain/model"
	"time"
	"crypto/sha256"
	"encoding/hex"
)

func GenerateBlock(oldBlock model.Block,BPM int)(model.Block,error){
	newBlock := model.Block{}
	t :=time.Now()
	newBlock.Index = oldBlock.Index +1
	newBlock.Timestamp = t.String()
	newBlock.BPM = BPM
	newBlock.PrevHash = oldBlock.Hash
	newBlock.Hash = CalculateHash(newBlock)
	return newBlock,nil
}

func CalculateHash(block model.Block) string{
	record := string(block.Index) + block.Timestamp + string(block.BPM) + block.PrevHash
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

func isBlockValid(newBlock,oldBlock model.Block) bool{
	if oldBlock.Index+1 != newBlock.Index{
		return false
	}
	if oldBlock.Hash!=newBlock.PrevHash{
		return false
	}
	if CalculateHash(newBlock)!= newBlock.Hash{
		return false
	}
	return true
}