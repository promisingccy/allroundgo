package main

import (
	"fmt"
	"time"
)

// 监听多个通道的消息
func main() {

	ch1 := make(chan int)
	ch2 := make(chan int)

	// 开启 goroutine 1 用于向通道 ch1 发送数据
	go func() {
		for i := 0; i < 5; i++ {
			ch1 <- i
			time.Sleep(time.Second)
		}
	}()

	// 开启 goroutine 2 用于向通道 ch2 发送数据
	go func() {
		for i := 5; i < 10; i++ {
			ch2 <- i
			time.Sleep(time.Second)
		}
	}()

	// 主 goroutine 从 ch1 和 ch2 中接收数据并打印
	for i := 0; i < 10; i++ {
		select {
		case data, ok := <-ch1:
			if ok {
				fmt.Println("ch1:", data)
			} else {
				fmt.Println("channel closed")
			}
		case data, ok := <-ch2:
			if ok {
				fmt.Println("ch2:", data)
			} else {
				fmt.Println("channel closed")
			}
		}
	}
	fmt.Println("done.....")
}

// 执行结果（每次执行程序打印的顺序都不一致）：
//ch1: 0
//ch2: 5
//ch2: 6
//ch1: 1
//ch1: 2
//ch2: 7
//ch1: 3
//ch2: 8
//ch2: 9
//ch1: 4
//done.....
