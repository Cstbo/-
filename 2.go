package main

import (
	"fmt"
	"sync"
	"time"
)

type Luggage struct {
	Owner   string
	Content string
}

func securityCheck(beltChan chan Luggage, doneChan chan bool) {
	for Check := range beltChan {
		time.Sleep(time.Microsecond * 500)
		if Check.Content == "炸弹" || Check.Content == "Knife" {
			fmt.Println("拦截!")
		} else {
			fmt.Println("通过!")
		}
	} 
	doneChan <- true
}

func Passenger(name string, item string, Beltchan chan Luggage, wg *sync.WaitGroup) {
	defer wg.Done()
	L1 := Luggage{
		Owner:   name,
		Content: item,
	}
	fmt.Printf("[%s]把[%s]放上了行李\n", name, item)
	Beltchan <- L1
}

func main() {
	var wg sync.WaitGroup
	beltChan := make(chan Luggage, 5)
	doneChan := make(chan bool)

	go securityCheck(beltChan, doneChan)

	wg.Add(1)
	go Passenger("张三", "衣服", beltChan, &wg)

	wg.Add(1)
	go Passenger("李四", "炸弹", beltChan, &wg)

	wg.Add(1)
	go Passenger("王五", "电脑", beltChan, &wg)

	wg.Wait()
	close(beltChan)
	<-doneChan
}
