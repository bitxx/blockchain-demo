package main

import (
	"github.com/bitxx/blockchain-demo/src/github.com/jason/blockchain-demo/blockchain/model"
	"github.com/bitxx/blockchain-demo/src/github.com/jason/blockchain-demo/blockchain/service"
	"github.com/davecgh/go-spew/spew"
	"log"
	"time"
)

func main() {
	go func() {
		t := time.Now()
		genesisBlock := model.Block{Timestamp: t.String()}
		genesisBlock.Hash = service.CalculateHash(genesisBlock)
		spew.Dump(genesisBlock)
		model.Blockchain = append(model.Blockchain, genesisBlock)
	}()
	log.Fatal(service.Run())
}
