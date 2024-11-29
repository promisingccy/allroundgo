package main

import (
	"fmt"
	"sync"
	"time"
)

// waitForLittleChan 使用信道的方法，在单个协程或者协程数少的时候，并不会有什么问题，但在协程数多的时候，代码就会显得非常复杂
func waitForLittleChan() {
	done := make(chan bool)
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println(i)
		}
		done <- true
	}()
	<-done
}

func main() {
	// 为了保证 main goroutine 在所有的 goroutine 都执行完毕后再退出
	// 在实际开发中，开发人员是无法预知，所有的 goroutine 需要多长的时间才能执行完毕，sleep 多了吧主程序就阻塞了， sleep 少了吧有的子协程的任务就没法完成。
	// “不要通过共享内存来通信，要通过通信来共享内存”
	waitForLittleChan()

	// 实现一主多子的协程协作方式
	waitForGroup()
}

func waitForGroup() {
	//这个 WaitGroup 用于等待这里启动的所有协程完成。
	//注意：如果 WaitGroup 显式传递到函数中，则应使用 指针 。
	//
	//实例化完成后，就可以使用它的几个方法：
	//    Add：初始值为0，你传入的值会往计数器上加，这里直接传入你子协程的数量
	//    Done：当某个子协程完成后，可调用此方法，会从计数器上减一，通常可以使用 defer 来调用。
	//    Wait：阻塞当前协程，直到实例里的计数器归零。
	var wg sync.WaitGroup

	//启动几个协程，并为其递增 WaitGroup 的计数器。
	for i := 1; i <= 5; i++ {
		wg.Add(1)

		//避免在每个协程闭包中重复利用相同的 i 值
		i := i

		//将 worker 调用包装在一个闭包中，
		//可以确保通知 WaitGroup 此工作线程已完成。
		//这样，worker 线程本身就不必知道其执行中涉及的并发原语。
		go func() {
			defer wg.Done()
			worker1(i)
		}()
	}

	//阻塞，直到 WaitGroup 计数器恢复为 0； 即所有协程的工作都已经完成。
	wg.Wait()

	//请注意，WaitGroup 这种使用方式没有直观的办法传递来自 worker 的错误。
	//更高级的用例，请参见 errgroup package
}

func worker1(id int) {
	fmt.Println("worker start", id)
	time.Sleep(time.Second)
	fmt.Println("worker end", id)
}

//worker start 5
//worker start 2
//worker start 3
//worker start 1
//worker start 4
//worker end 2
//worker end 4
//worker end 5
//worker end 1
//worker end 3
