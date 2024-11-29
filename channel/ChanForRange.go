package main

import "fmt"

func main() {
	pipeline := make(chan int, 10)

	go fibonacci(pipeline)

	for k := range pipeline {
		fmt.Println(k)
	}
	//1
	//2
	//3
	//5
	//8
	//13
	//21
	//34
	//55
	//89
}

func fibonacci(ch chan int) {
	n := cap(ch)
	x, y := 1, 2
	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x+y
	}
	// 记得 close 信道
	// 不然主函数中遍历完并不会结束，而是会阻塞。
	close(ch)
}
