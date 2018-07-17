package model

type Block struct {
	Index int  //是区块链中数据记录的位置
	Timestamp string  //是自动确定的，并且是写入数据的时间
	BPM int  //或每分钟跳动，是你的脉率,
	Hash string  //是代表这个数据记录的SHA256标识符
	PrevHash string  //是链中上一条记录的SHA256标识符
	Delegate string  //这个就是dpos中所说的委托人，也就是超级节点人之一
}

var Blockchain []Block //主链

var Delegates = []string{"001","002","003","004","005"} //预先定义的委托人，以及顺序

var IndexDelegate int //当前的 delegates 的索引