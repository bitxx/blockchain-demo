package service

import (
	"net"
	"github.com/jason/blockchain-demo/pos/model"
	"io"
	"bufio"
	"strconv"
	"log"
	"time"
	"fmt"
	"encoding/json"
	"math/rand"
)

/**
	当一个验证者连接到我们的TCP服务，我们需要提供一些函数达到以下目标：

	输入令牌的余额（之前提到过，我们不做钱包等逻辑)
	接收区块链的最新广播
	接收验证者赢得区块的广播信息
	将自身节点添加到全局的验证者列表中（validators)
	输入Block的BPM数据- BPM是每个验证者的人体脉搏值
	提议创建一个新的区块
 */

func HandleConn(conn net.Conn) {
	defer conn.Close()

	go func() {
		for {
			msg := <-model.Announcements
			io.WriteString(conn, msg)
		}
	}()

	//模拟一个用户钱包
	var address string

	io.WriteString(conn, "Enter token balance:")

	scanBalance := bufio.NewScanner(conn)
	for scanBalance.Scan() {
		balance, err := strconv.Atoi(scanBalance.Text())
		if err != nil {
			log.Printf("%v not a number:%v", scanBalance.Text(), err)
			return
		}
		t := time.Now()
		address = calculateHash(t.String()) //相当于随机生成一个地址
		model.Validators[address] = balance
		fmt.Println(model.Validators)
		break
	}

	io.WriteString(conn, "\nEnter a new BPM:")
	scanBPM := bufio.NewScanner(conn)
	go func() {
		for {
			for scanBPM.Scan() {
				bpm, err := strconv.Atoi(scanBPM.Text())
				if err != nil {
					log.Printf("%v not a number:%v", scanBPM.Text(), err)
					delete(model.Validators, address) //切片中删除该地址
					conn.Close()
				}
				model.Mutex.Lock()
				oldLastIndex := model.Blockchain[len(model.Blockchain)-1]
				model.Mutex.Unlock()
				newBlock, err := generateBlock(oldLastIndex, bpm, address)
				if err != nil {
					log.Println(err)
					continue
				}
				if isBlockValid(newBlock, oldLastIndex) {
					model.CandidateBlocks <- newBlock
				}
				io.WriteString(conn, "\nEnter a new BPM:")
			}
		}
	}()

	for {
		time.Sleep(time.Minute)
		model.Mutex.Lock()
		output, err := json.Marshal(model.Blockchain)
		model.Mutex.Unlock()
		if err != nil {
			log.Fatal(err)
		}
		io.WriteString(conn, string(output)+"\n")
	}
}

//选择可以生产块的生产者，他们所持有的令牌数量越高，他们就越有可能被选为胜利者。
func PickWinner() {
	time.Sleep(10 * time.Second)
	model.Mutex.Lock()
	temp := model.TempBlocks
	model.Mutex.Unlock()

	lotteryPool := []string{}
	if len(temp) > 0 {
	OUTER:
		for _, block := range temp {
			for _, node := range lotteryPool {
				if block.Validator == node {
					continue OUTER
				}
			}

			model.Mutex.Lock()
			setValidators := model.Validators
			model.Mutex.Unlock()

			k, ok := setValidators[block.Validator]
			if ok {
				for i := 0; i < k; i++ {
					lotteryPool = append(lotteryPool, block.Validator)
				}
			}

		}

		s := rand.NewSource(time.Now().Unix())
		r := rand.New(s)
		lotteryWinner := lotteryPool[r.Intn(len(lotteryPool))]

		for _, block := range temp {
			if block.Validator == lotteryWinner {
				model.Mutex.Lock()
				model.Blockchain = append(model.Blockchain, block)
				model.Mutex.Unlock()
				for _ = range model.Validators {
					model.Announcements <- "\nwinning validator: " + lotteryWinner + "\n"
				}
				break
			}
		}
	}

	model.Mutex.Lock()
	model.TempBlocks = []model.Block{}
	model.Mutex.Unlock()
}
