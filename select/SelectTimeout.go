package main

import (
	"fmt"
	"time"
)

func main() {

	//----------- 2s后写入信道 等待时间1s
	// 会超时
	ch1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- "result 1"
	}()

	select {
	case res := <-ch1:
		fmt.Println(res)
	case <-time.After(time.Second):
		fmt.Println("timeout 1")
	}

	// ---------- 2s后写入信道 等待3秒
	// 不会超时
	ch2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "result 2"
	}()

	select {
	case res := <-ch2:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2")
	}
}

//timeout 1  ch1会超时
//result 2  ch2不会超时
