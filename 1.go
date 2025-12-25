package main

import (
	"fmt"
	"sync"
	"time"
)

type Coffee struct {
	CoffeeType   string
	CustomerName string
}

// 1. 修正：咖啡师只需要通道，不需要提前知道具体的咖啡信息
// 注意：通道类型必须是 chan Coffee，不能是 chan string
func makeCoffee(ordersChan chan Coffee, doneChan chan bool) {
	// 2. 修正：从通道里取出来的数据，我们要给它起个名字叫 order
	for order := range ordersChan {
		time.Sleep(time.Second)
		// 3. 修正：直接打印取出来的 order 里的信息
		fmt.Printf("☕ 咖啡师: 给[%s]做好了一杯[%s]\n", order.CustomerName, order.CoffeeType)
	}
	// 打卡下班
	doneChan <- true
}

// 4. 修正：顾客函数需要传入具体的 名字(name) 和 咖啡名(coffee)
func orderCoffee(name string, coffee string, ordersChan chan Coffee, wg *sync.WaitGroup) {
	defer wg.Done()
	
	fmt.Printf("📝 顾客: [%s] 下单了 [%s]\n", name, coffee)

	// 5. 修正：关键！要创建一个“结构体实例”发送出去
	// 不能写 ordersChan <- Coffee
	order := Coffee{
		CustomerName: name,
		CoffeeType:   coffee,
	}
	ordersChan <- order
}

func main() {
	var wg sync.WaitGroup
	
	// 6. 修正：创建通道，必须指明是运送 Coffee 结构体的
	ordersChan := make(chan Coffee, 5)
	doneChan := make(chan bool)

	// 启动咖啡师
	go makeCoffee(ordersChan, doneChan)

	// 模拟 3 个顾客下单
	// 顾客 1
	wg.Add(1)
	go orderCoffee("张三", "拿铁", ordersChan, &wg)

	// 顾客 2
	wg.Add(1)
	go orderCoffee("李四", "美式", ordersChan, &wg)

	// 顾客 3
	wg.Add(1)
	go orderCoffee("王五", "卡布奇诺", ordersChan, &wg)

	// 等待所有顾客下完单
	wg.Wait()
	// 关闭点单通道
	close(ordersChan)

	// 等待咖啡师下班
	<-doneChan
	fmt.Println("🛑 咖啡店打烊了")
}