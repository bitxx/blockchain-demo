package service

import (
	"github.com/jason/blockchain-demo/pow/model"
	"time"
	"fmt"
	"log"
	"strings"
	"strconv"
	"crypto/sha256"
	"encoding/hex"
)

const difficulty = 1

func GenerateBlock(oldBlock model.Block, BPM int) model.Block {
	var newBlock model.Block
	t := time.Now()
	newBlock.Index = oldBlock.Index + 1
	newBlock.Timestamp = t.String()
	newBlock.Diffculty = difficulty
	newBlock.BPM = BPM
	newBlock.PrevHash = oldBlock.Hash

	for i := 0; ; i++ {
		hex := fmt.Sprintf("%x", i)
		newBlock.Nonce = hex
		newHash := calculateHash(newBlock)
		if !isHashValid(newHash,difficulty){
			log.Println("hash err:"+newHash+" continue do work")
			time.Sleep(time.Second)
		}else {
			log.Println("success hash")
			newBlock.Hash = newHash
			break
		}
	}
	return newBlock

}

/**
 * 验证hash是否有效，通过判断前面0的个数
 */
func isHashValid(hash string, difficulty int) bool {
	prefix := strings.Repeat("0",difficulty)
	hasPrefix := strings.HasPrefix(hash,prefix)
	return hasPrefix
}


func calculateHash(block model.Block) string {
	record := strconv.Itoa(block.Index)+block.Timestamp+strconv.Itoa(block.BPM)+block.PrevHash+block.Nonce
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}


func isBlockValid(newBlock,oldBlock model.Block) bool{
	if oldBlock.Index+1!=newBlock.Index{
		return false
	}
	if oldBlock.Hash!=newBlock.PrevHash{
		return false
	}
	if calculateHash(newBlock)!=newBlock.Hash{
		return false
	}
	return true
}