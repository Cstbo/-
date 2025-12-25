package main

import (
	"fmt"
	"sync"
	"time"
)

type Parcel struct {
	ID      string
	Owner   string
	Content string
}

func worker(chanName string, Channel <-chan Parcel, wg *sync.WaitGroup) {
	defer wg.Done()
	All := 0
	Past := 0
	Block := 0
	for p := range Channel {
		time.Sleep(time.Second)
		if p.Content == "炸弹" || p.Content == "刀" {
			fmt.Printf("❌[%s],ID[%s],收件人名字[%s],物品[%s],禁止放行\n", chanName, p.ID, p.Owner, p.Content)
			Block++
			All++
		} else {
			fmt.Printf("✔[%s],ID[%s],收件人名字[%s],物品[%s],放行\n", chanName, p.ID, p.Owner, p.Content)
			Past++
			All++
		}
	}
	fmt.Printf("[%s]关闭,累计处理[%d]人,通行[%d]个物品\n", chanName, All, Past)
}

func main() {

	parcels := []Parcel{
		{"PKG-001", "1", "衣服"},
		{"PKG-002", "2", "刀"},
		{"PKG-003", "3", "炸弹"},
		{"PKG-004", "4", "衣服"},
		{"PKG-005", "5", "刀"},
		{"PKG-006", "6", "衣服"},
		{"PKG-007", "7", "衣服"},
		{"PKG-008", "8", "衣服"},
		{"PKG-009", "9", "炸弹"},
	}
	var workerwg sync.WaitGroup
	var senderwg sync.WaitGroup

	Channel := make(chan Parcel, 5)

	chans := []string{"GateA", "GateB", "GateC"}
	workerwg.Add(len(chans)) //?

	for _, g := range chans {
		go worker(g, Channel, &workerwg)
	}

	senderwg.Add(len(parcels))
	for _, p := range parcels {
		go func(ps Parcel) {
			defer senderwg.Done()
			fmt.Printf("[%s]正在前往发送包裹给[%s]...\n", p.ID, p.Owner)
			Channel <- ps
		}(p)
	}

	senderwg.Wait()
	close(Channel)
	workerwg.Wait()

}
