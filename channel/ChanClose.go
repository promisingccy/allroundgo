package main

import "fmt"

func main() {
	//接收任务
	jobs := make(chan int, 5)
	//标记任务是否已经结束
	done := make(chan bool)

	go func() {
		for {
			//使用 j, more := <- jobs 循环的从 jobs 接收数据。
			//如果接收到数据 hasV=true
			jv, hasV := <-jobs
			if hasV {
				fmt.Println("received job", jv)
			} else {
				//如果 jobs 已关闭，并且通道中所有的值都已经接收完毕
				//那么 hasV=false
				fmt.Println("received all jobs")
				//当我们完成所有的任务时，会使用这个特性通过 done 通道通知 main 协程。
				done <- true
				return
			}
		}
	}()

	//使用 jobs 发送 3 个任务到工作协程中，然后关闭 jobs
	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}

	// 关闭 一个通道意味着不能再向这个通道发送值了。
	// 该特性可以向通道的接收方传达工作已经完成的信息。
	close(jobs)
	fmt.Println("sent all jobs")

	// 通道同步的方式，作用：等待任务结束
	<-done
}

//sent job 1
//sent job 2
//sent job 3
//sent all jobs
//received job 1
//received job 2
//received job 3
//received all jobs
