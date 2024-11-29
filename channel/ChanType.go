package main

import (
	"fmt"
	"time"
)

// Sender 定义只写信道类型
type Sender = chan<- int

// Receiver 定义只读信道类型
type Receiver = <-chan int

func main() {
	var pipline = make(chan int)

	go func() {
		var receiver Receiver = pipline
		num := <-receiver
		fmt.Printf("接收到的数据是: %d", num)
	}()

	go func() {
		var sender Sender = pipline
		fmt.Println("准备发送数据: 100")
		sender <- 100
	}()

	// 主函数sleep，使得上面两个goroutine有机会执行
	time.Sleep(time.Second)
}
