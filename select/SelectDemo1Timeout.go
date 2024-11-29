package main

import (
	"fmt"
	"time"
)

// 实现超时控制
func main() {

	// 在这个例子中，程序将在 3 秒后向 ch 通道里写入数据，
	// 而我在 select 代码块里设置的超时时间为 2 秒，
	// 如果在 2 秒内没有接收到数据，则会触发超时处理。
	ch := make(chan int)

	go func() {
		time.Sleep(3 * time.Second)
		ch <- 1
	}()

	select {
	case data, ok := <-ch:
		if ok {
			fmt.Println(data)
		} else {
			fmt.Println("ch closed")
		}
	case <-time.After(2 * time.Second):
		fmt.Println("timeout")
	}
}
