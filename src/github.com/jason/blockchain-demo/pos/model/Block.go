package model

import "sync"

type Block struct {
	Index     int    //是区块链中数据记录的位置
	Timestamp string //是自动确定的，并且是写入数据的时间
	BPM       int    //或每分钟跳动，是你的脉率,
	Hash      string //是代表这个数据记录的SHA256标识符
	PrevHash  string //是链中上一条记录的SHA256标识符
	Validator string //存储的是获取记账权的节点地址
}

var Blockchain []Block
var TempBlocks []Block //这个是块提议队列，CandidateBlocks将块发送到它里面

var CandidateBlocks = make(chan Block) //任何一个节点在提出一个新块时都将它发送到这个通道

var Announcements = make(chan string)

var Mutex = &sync.Mutex{}

var Validators = make(map[string]int)
