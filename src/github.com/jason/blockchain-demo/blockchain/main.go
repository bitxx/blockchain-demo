package blockchain

import (
	"github.com/joho/godotenv"
	"time"
	"github.com/jason/blockchain-demo/blockchain/model"
	"github.com/davecgh/go-spew/spew"
	"log"
	"github.com/jason/blockchain-demo/blockchain/service"
)

func main() {
	err := godotenv.Load()
	if err!=nil{
		log.Fatal(err)
	}
	go func() {
		t := time.Now()
		genesisBlock := model.Block{0, t.String(), 0, "", ""}
		genesisBlock.Hash = service.CalculateHash(genesisBlock)
		spew.Dump(genesisBlock)
		model.Blockchain = append(model.Blockchain, genesisBlock)
	}()
	log.Fatal(service.Run())
}
