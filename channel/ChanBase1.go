package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	//demo0GoRoutine()
	//demo1ChanBase()
	//demo1ChanMsg()
	Sync()
	//demoxxx()

}

func demoxxx() {
	intChan0 := make(chan int, 10)

	for i := 0; i < 10; i++ {
		go func() {
			intChan0 <- i
		}()
	}
	defer close(intChan0)

	for i := range intChan0 {
		fmt.Printf("range i value is %d, pointer is %d\n", i, &i)
	}

}

// /-------------------------

// Sync 并发处理耗时过程 并把耗时结果收集后 同步向下进行
func Sync() {

	strChan := make(chan string, 10)
	for i := 0; i < 10; i++ {
		nu := i
		go func(i int) {
			repeat := strings.Repeat("i", i+1)
			strChan <- repeat
		}(nu)
	}
	defer close(strChan)

	strs := make([]string, 0)
	for i := 0; i < 10; i++ {
		select {
		case cur, ok := <-strChan:
			if ok {
				strs = append(strs, cur)
			}
		}
	}

	fmt.Println(strs)
	//[iiiiiiiiii iiii iiiii iiiiii iiiiiii iiiiiiii iiiiiiiii i ii iii]
}

func ppp(strChan chan string) {
	for i := range strChan {
		//strs = append(strs, i)
		fmt.Printf("range i value is %d, pointer is %d\n", i, &i)
	}
}

// /-------------------------
func demo0GoRoutine() {
	f("direct")
	go f("goroutine")
	go func(msg string) {
		fmt.Println(msg)
	}("no name function")
	time.Sleep(1 * time.Second)
	fmt.Println("done")
}

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

// /-------------------------
// 默认发送和接收操作是阻塞的，直到发送方和接收方都就绪
func demo1ChanBase() {
	//创建一个新的通道
	msg := make(chan string)
	go func() {
		//channel <- 语法 发送 一个新的值到通道中
		msg <- "hello world"
	}()
	//使用 <-channel 语法从通道中 接收 一个值
	s := <-msg
	fmt.Println(s)
}
func demo1ChanMsg() {
	done := make(chan bool, 1)
	//运行一个 worker 协程，并给予用于通知的通道。
	//程序将一直阻塞，直至收到 worker 使用通道发送的通知。
	go worker(done)
	//若注释掉 <-done 什么都没有输出 程序就退出了
	<-done
}
func worker(done chan bool) {
	fmt.Println("worker started")
	time.Sleep(1 * time.Second)
	fmt.Println("worker finished")
	done <- true
}
