package main

import "fmt"

// 实现多任务并发控制
func main() {
	// 在这个例子中，启动了 10 个 goroutine 并发执行任务，
	// 并使用一个 channel 来接收任务的完成情况。
	// 在主函数中，使用 select 语句监听这个 channel，
	// 每当接收到一个完成的任务时，就进行处理。
	ch := make(chan int)

	for i := 0; i < 10; i++ {
		go func(id int) {
			ch <- id
		}(i)
	}

	// 使用 select 语句监听这个 channel，
	// 每当接收到一个完成的任务时，就进行处理
	for i := 0; i < 10; i++ {
		select {
		case data, ok := <-ch:
			if ok {
				fmt.Println(data)
			} else {
				fmt.Println("channel closed")
			}
		}
	}
	// 执行结果（每次执行的顺序都会不一致）：
	//0
	//6
	//5
	//7
	//8
	//2
	//1
	//3
	//4
	//9

	//0
	//3
	//1
	//2
	//4
	//6
	//8
	//7
	//9
	//5
}
