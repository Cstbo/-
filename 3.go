package main

import (
	"fmt"
	"sync"
	"time"
)

type Vehicle struct {
	Plate string
	Type  string
}

func collectMoney(carChan chan Vehicle, doneChan chan bool) {
	for car := range carChan {
		time.Sleep(time.Second)
		if car.Type == "Car" {
			fmt.Printf("车牌是[%s]是小车，收费十元", car.Type)
		} else {
			fmt.Printf("车牌是[%s]是大车，收费三十元", car.Type)
		}

	}
	doneChan <- true
}

func driver(carPlate string, carType string, carChan chan Vehicle, wg *sync.WaitGroup) {
	defer wg.Done()
	car1 := Vehicle{
		Plate: carPlate,
		Type:  carType,
	}
	fmt.Printf("司机:车牌[%s]正在进入收费站", carPlate)
	carChan <- car1
}

func main() {
	var wg sync.WaitGroup
	carChan := make(chan Vehicle)
	doneChan := make(chan bool)

	go collectMoney(carChan, doneChan)

	wg.Add(1)
	go driver("A-11111", "Car", carChan, &wg)

	wg.Add(1)

	go driver("A-22222", "Trunk", carChan, &wg)

	wg.Add(1)

	go driver("A-33333", "Car", carChan, &wg)

	wg.Wait()
	close(carChan)
	<-doneChan
}
