package main

import (
	"fmt"
	"github.com/bitxx/blockchain-demo/src/github.com/jason/blockchain-demo/dpos/model"
	"github.com/bitxx/blockchain-demo/src/github.com/jason/blockchain-demo/dpos/service"
	"log"
	"math/rand"
	"time"
)

func main() {
	model.IndexDelegate = 0 //初始化委托人索引

	t := time.Now()
	genesisBlock := model.Block{}
	genesisBlock = model.Block{0, t.String(), 0, service.CalculateBlockHash(genesisBlock), "", ""}
	model.Blockchain = append(model.Blockchain, genesisBlock)
	model.IndexDelegate++

	countDelegate := len(model.Delegates)

	for model.IndexDelegate < countDelegate {
		time.Sleep(time.Second * 3)
		fmt.Println(model.IndexDelegate)

		//创建新的区块
		rand.Seed(int64(time.Now().Unix())) //随机种子
		bpm := rand.Intn(100)
		oldLastIndex := model.Blockchain[len(model.Blockchain)-1]
		newBlock, err := service.GenerateBlock(oldLastIndex, bpm, model.Delegates[model.IndexDelegate])
		if err != nil {
			log.Println(err)
			continue
		}
		fmt.Printf("Blockchain...%v\n", newBlock)
		if service.IsBlockValid(newBlock, oldLastIndex) {
			model.Blockchain = append(model.Blockchain, newBlock)
		}

		model.IndexDelegate = (model.IndexDelegate + 1) % countDelegate
		if model.IndexDelegate == 0 {
			model.Delegates = service.RandDelegates(model.Delegates) //第一轮块生成结束，则洗牌，进行下一轮投票
		}

	}

}
