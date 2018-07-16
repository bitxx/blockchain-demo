package model

import "sync"

type Block struct {
	Index int
	Timestamp string
	BPM int
	Hash string
	PrevHash string
	Validator string  //存储的是获取记账权的节点地址
}

var Blockchain []Block
var TempBlocks []Block

var CandidateBlocks = make(chan Block)

var Announcements = make(chan string)

var Mutex = &sync.Mutex{}

var Validators = make(map[string]int)
