package service

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/bitxx/blockchain-demo/src/github.com/jason/blockchain-demo/pos/model"
	"time"
)

func generateBlock(oldBlock model.Block, BPM int, address string) (model.Block, error) {
	var newBlock model.Block
	t := time.Now()
	newBlock.Index = oldBlock.Index + 1
	newBlock.Timestamp = t.String()
	newBlock.BPM = BPM
	newBlock.PrevHash = oldBlock.Hash
	newBlock.Hash = CalculateBlockHash(newBlock)
	newBlock.Validator = address //存储的是获取记账权的节点地址
	return newBlock, nil
}

// 将原字符串生成一个hash值
func calculateHash(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

// 将一个块生成一个hash值
func CalculateBlockHash(block model.Block) string {
	record := string(block.Index) + block.Timestamp + string(block.BPM) + block.PrevHash
	return calculateHash(record)
}

// 验证区块
func isBlockValid(newBlock, oldBlock model.Block) bool {
	if oldBlock.Index+1 != newBlock.Index {
		return false
	}

	if oldBlock.Hash != newBlock.PrevHash {
		return false
	}

	if CalculateBlockHash(newBlock) != newBlock.Hash {
		return false
	}

	return true

}
