package main

import (
	"fmt"
	"time"
)

// make(chan string,10) 缓冲通道只要通道容量给的足够，可以不使用协程
// make(chan string) 非缓冲通道则必须使用协程了
// 如果通道写满，向通道中写入的协程将被阻塞，直到通道中数据被接收时该协程被唤醒。
func main() {

	fmt.Println("WaitChan")
	WaitChan()
	//fmt.Println("SingleDirSingleChan")
	//SingleDirSingleChan()
	//fmt.Println("Select4Chan")
	//Select4Chan()
	//fmt.Println("SingleDirChanWith2Chan")
	//SingleDirChanWith2Chan()
	//fmt.Println("SingleDirChan")
	//SingleDirChan()
	//fmt.Println("Close4Chan")
	//Close4Chan()
	//fmt.Println("Foreach4Chan")
	//Foreach4Chan()
}

// ----------------------- 单向的非缓冲通道
func WaitChan() {
	nobufChan := make(chan string)
	go func() {
		time.Sleep(time.Second * 3)
		msg := <-nobufChan
		fmt.Println(msg)
	}()

	msg := "this msg"
	nobufChan <- msg
}

func SingleDirSingleChan() {
	ch := make(chan string)
	go inputOnly(ch)
	go outputOnly(ch)
	time.Sleep(time.Second)
}

// ----------------------- select 通道的选择
func Select4Chan() {
	// 缓冲通道 可以不使用协程执行inputWithGroup
	chGA := make(chan string, 8)
	chGB := make(chan string, 8)
	inputWithGroup(chGA, "A")
	inputWithGroup(chGB, "B")

	// 非缓冲通道 必须使用协程执行 inputWithGroup 否则报错
	//chGA := make(chan string)
	//chGB := make(chan string)
	//go inputWithGroup(chGA, "A")
	//go inputWithGroup(chGB, "B")

	go func() {
		for {
			time.Sleep(500 * time.Millisecond)
			select {
			case msg, ok := <-chGA:
				if ok {
					fmt.Println(msg)
				}
			case msg, ok := <-chGB:
				if ok {
					fmt.Println(msg)
				}
			default:
				fmt.Println("hello world")
			}
		}
	}()
	// 由于两个通道一共8个元素，所以需要至少8秒才能打印结束
	time.Sleep(10 * time.Second)
}

func inputWithGroup(c chan string, g string) {
	names := []string{"a", "b", "c", "d"}
	for _, name := range names {
		c <- g + "-" + name
	}
	close(c)
}

// -----------------------
// 【写入】 chan <- v    chan<-
// 【读出】v := <-chan   <-chan
func SingleDirChanWith2Chan() {
	singleChanIn := make(chan string)
	singleChanOut := make(chan string)
	// 非协程执行会报错 非缓冲通道则必须使用协程了
	// 如果通道写满，向通道中写入的协程将被阻塞，直到通道中数据被接收时该协程被唤醒。
	go inputOnly(singleChanIn)
	go in2out(singleChanOut, singleChanIn)
	go outputOnly(singleChanOut)
	time.Sleep(time.Second)
}

func in2out(in chan<- string, out <-chan string) {
	for s := range out {
		in <- s
	}
	close(in)
}

// -----------------------
func SingleDirChan() {
	// 发送通道和接收通道 单向通道可分为接收通道和发送通道 一个仅发送，一个仅接收（否则报错）
	// 在函数中将通道定义为单项，避免函数对通道造成污染。

	ch := make(chan string, 10)
	// [发送通道] 通道名 chan<- 传入数据类型
	// var in chan<- string
	go inputOnly(ch)
	//[传出通道] 通道名 <-chan 传出数据类型
	// var out <-chan string
	go outputOnly(ch)
	time.Sleep(time.Second)
}

func inputOnly(in chan<- string) {
	names := []string{"a", "b", "c", "d"}
	for _, name := range names {
		in <- name
	}
	close(in)
}

func outputOnly(out <-chan string) {
	for v := range out {
		fmt.Println(v)
	}
}

// -----------------------
func Close4Chan() {
	ch := make(chan string, 10)
	go input(ch)
	fmt.Println(len(ch)) //0
	go output(ch)
	fmt.Println(cap(ch)) //10
	fmt.Println(len(ch)) //0-4 随机获取根据是否能消费到
	time.Sleep(time.Second * 3)
	fmt.Println(len(ch)) //0
}

// -----------------------
func Foreach4Chan() {
	// 如果通道写满，向通道中写入的协程将被阻塞，直到通道中数据被接收时该协程被唤醒。
	// 变量将值传给通道之后，变量值变化与通道中数据无关。
	names := make(chan string, 13)
	input(names)
	for {
		// 接收元素值
		// ok值：如果读出数据为true ，未读出为false（比如通道中没有数据了）
		// ok值的作用：作为一个布尔值，可用作之后的if判断
		// 从未初始化的通道接收数据 此时通道会被永远阻塞
		name, ok := <-names
		if ok {
			// 没有顺序，协程谁先打印出结果是随机的
			go fmt.Println(name)
		} else {
			break
		}
	}
	time.Sleep(time.Second)
}

// -----------------------
func input(c chan string) {
	names := []string{"a", "b", "c", "d"}
	for _, v := range names {
		// 变量将值传给通道之后，变量值变化与通道中数据无关。
		c <- v
	}
	//关闭通道
	//关闭时机 建议在输入端关闭 通道关闭之后仍可以读出
	close(c)
}

func output(c chan string) {
	for s := range c {
		fmt.Println(s)
	}
}
