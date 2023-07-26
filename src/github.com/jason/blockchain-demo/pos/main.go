package main

import (
	"github.com/bitxx/blockchain-demo/src/github.com/jason/blockchain-demo/pos/model"
	"github.com/bitxx/blockchain-demo/src/github.com/jason/blockchain-demo/pos/service"
	"github.com/davecgh/go-spew/spew"
	"log"
	"net"
	"os"
	"time"
)

func main() {

	t := time.Now()
	genesisBlock := model.Block{}
	genesisBlock = model.Block{Timestamp: t.String(), Hash: service.CalculateBlockHash(genesisBlock)}
	spew.Dump(genesisBlock) //美观的将数据打印出，创始块
	model.Blockchain = append(model.Blockchain, genesisBlock)
	httpPort := os.Getenv("PORT")

	server, err := net.Listen("tcp", ":"+httpPort) //tcp形式监听，
	if err != nil {
		log.Fatal(err)
	}
	log.Println("HTTP Server Listening on port :", httpPort)
	defer server.Close()

	go func() {
		for candidate := range model.CandidateBlocks {
			model.Mutex.Lock()
			model.TempBlocks = append(model.TempBlocks, candidate)
			model.Mutex.Unlock()
		}
	}()

	go func() {
		for {
			service.PickWinner()
		}
	}()

	for {
		conn, err := server.Accept() //tcp接收监听
		if err != nil {
			log.Fatal(err)
		}
		go service.HandleConn(conn) //
	}

}
