package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	//过1秒写入数据1
	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- 1
	}()

	//过两秒写入数据2
	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- 2
	}()

	// ch1 和ch2 的数据会在这个select块里被接收到并打印输出
	for i := 0; i < 2; i++ {
		select {
		case data, ok := <-ch1:
			if ok {
				fmt.Println("ch1:", data)
			} else {
				fmt.Println("ch1 closed")
			}
		case data, ok := <-ch2:
			if ok {
				fmt.Println("ch2:", data)
			} else {
				fmt.Println("ch2 closed")
			}
		}
	}

	//到此select块时，ch1和ch2都没有数据了，此时会走到default分支
	select {
	case data, ok := <-ch1:
		if ok {
			fmt.Println("ch1:", data)
		} else {
			fmt.Println("ch1 closed")
		}
	case data, ok := <-ch2:
		if ok {
			fmt.Println("ch2:", data)
		} else {
			fmt.Println("ch2 closed")
		}
	default:
		fmt.Println("此处用于验证信道没有数据时，会走到此处")
	}
}

//ch1: 1
//ch2: 2
//此处用于验证信道没有数据时，会走到此处
