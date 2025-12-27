package main

import (
	"context"
	"fmt"
	"sync"
)

type Parcel struct {
	ID    string
	Owner string
	Item  string
}

type Result struct {
	Gate   string
	Passed bool
}

func worker(ChanName string, Channel <-chan Parcel, ResultChannel chan<- Result, ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	All := 0
	Block := 0
	Pass := 0
	for {

		select {
		case <-ctx.Done():
			fmt.Printf("[%s]收到信号，退出\n", ChanName)
			return
		case p, ok := <-Channel:
			if !ok {
				fmt.Printf("[%s]输入通道关闭,退出\n", ChanName)
				return
			}

			Passed := true
			if p.Item == "炸弹" || p.Item == "刀" {
				Passed = false
			}
			All++
			if Passed {
				Pass++
				fmt.Printf("✔[%s],ID[%s],收件人名字[%s],物品[%s],放行\n", ChanName, p.ID, p.Owner, p.Item)
			} else {
				Block++
				fmt.Printf("❌[%s],ID[%s],收件人名字[%s],物品[%s],禁止放行\n", ChanName, p.ID, p.Owner, p.Item)
			}
			select {
			case ResultChannel <- Result{Gate: ChanName, Passed: Passed}:
			case <-ctx.Done():
				fmt.Printf("[%s]收到信号，退出\n", ChanName)
				fmt.Printf("名字:%s,总数:%d,放行:%d,阻止:%d\n", ChanName, All, Pass, Block)
				return
			}
		}
	}
}

func main() {
	Channel := make(chan Parcel, 5)
	ResultChannel := make(chan Result, 100)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var workerwg sync.WaitGroup
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
	chans := []string{"GateA", "GateB", "GateC"}
	type GateStat struct {
		Total   int
		Passed  int
		Blocked int
	}
	byGate := make(map[string]*GateStat)
	workerwg.Add(len(chans))
	for _, g := range chans {
		go worker(g, Channel, ResultChannel, ctx, &workerwg)
	}
	for _, p := range parcels {
		Channel <- p
	}
	close(Channel)
	go func() {
		workerwg.Wait()
		close(ResultChannel)
	}()

	total, passed, blocked := 0, 0, 0
	for r := range ResultChannel {
		if !r.Passed {
			if blocked >= 2 {
				cancel()
			}
		}
		s := byGate[r.Gate]
		if s == nil {
			s = &GateStat{}
			byGate[r.Gate] = s
		}
		fmt.Printf("r.Gate=%s, s=%p\n", r.Gate, s)
		s.Total++
		if r.Passed {
			s.Passed++
		} else {
			s.Blocked++
		}
		total++
		if r.Passed {
			passed++
		} else {
			blocked++
		}
	}
	fmt.Println("各 Gate 统计：")
	for gate, s := range byGate {
		fmt.Printf("%s：总数 %d，放行 %d，拦截 %d\n", gate, s.Total, s.Passed, s.Blocked)
	}
	fmt.Printf("总数:%d,放行:%d,阻止:%d\n", total, passed, blocked)
}
