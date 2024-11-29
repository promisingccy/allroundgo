package main

import (
	"fmt"
	"time"
)

func main() {
	const numJobs = 5
	//为了使用 worker 工作池并且收集其的结果，我们需要 2 个通道。
	jobs := make(chan int, numJobs)

	results := make(chan int, numJobs)

	//这里启动了 3 个 worker， 初始是阻塞的，因为还没有传递任务。
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	//这里我们发送 5 个 jobs，
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	//然后 close 这些通道， 表示这些就是所有的任务了。
	close(jobs)

	//最后，我们收集所有这些任务的返回值。
	//这也确保了所有的 worker 协程都已完成。
	for a := 1; a <= numJobs; a++ {
		i := <-results
		fmt.Println("get res :", i)
	}

	//运行程序，显示 5 个任务被多个 worker 执行。
	//尽管所有的工作总共要花费 5 秒钟，
	//但该程序只花了 2 秒钟， 因为 3 个 worker 是并行的。

}

// 这是 worker 程序，我们会并发的运行多个 worker。
// worker 将在 jobs 频道上接收工作，并在 results 上发送相应的结果。
// 每个 worker 我们都会 sleep 一秒钟，以模拟一项昂贵的（耗时一秒钟的）任务。
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println(id, j, "worker", id, "started job", j)
		time.Sleep(time.Second)
		fmt.Println(id, j, "worker", id, "finished job", j)
		results <- j * 2
		fmt.Println("------")
	}
}

//3 1 worker 3 started job 1
//1 2 worker 1 started job 2
//2 3 worker 2 started job 3
//2 3 worker 2 finished job 3
//------
//2 4 worker 2 started job 4
//get res : 6
//1 2 worker 1 finished job 2
//------
//1 5 worker 1 started job 5
//get res : 4
//3 1 worker 3 finished job 1
//------
//get res : 2
//2 4 worker 2 finished job 4
//------
//get res : 8
//1 5 worker 1 finished job 5
//------
//get res : 10
