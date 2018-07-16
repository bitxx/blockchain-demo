package main

import (
	"github.com/joho/godotenv"
	"github.com/jason/blockchain-demo/pow/service"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	go func() {

	}()
	log.Fatal(service.Run())
}
