package model

import "sync"

type Block struct {
	Index int
	Timestamp string
	BPM int
	Hash string
	PrevHash string
	Diffculty int
	Nonce string
}

var Blockchain []Block

type Message struct {
	BPM int
}

var Mutex = &sync.Mutex{}