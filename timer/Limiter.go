package main

import (
	"fmt"
	"time"
)

func main() {
	//////////// demo1 平稳消费

	//模拟5个请求
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	//limiter 通道每 200ms 接收一个值。 这是我们任务速率限制的调度器。
	limiter := time.Tick(500 * time.Millisecond)

	for req := range requests {
		//通过在每次请求前阻塞 limiter 通道的一个接收，
		//可以将频率限制为，每 200ms 执行一次请求。
		<-limiter
		fmt.Println("req:", req, time.Now())
	}

	fmt.Println("------------")
	//////////// demo2 允许短暂的并发请求，并同时保留总体速率限制
	//限定并发任务的数量=3
	burstyLimiter := make(chan time.Time, 3)

	//填充通道，表示允许的爆发（bursts）
	//填满chan 在消费时 会被瞬间全部消费完  从而现象是三个并发
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	//必须放到协程中执行 不然会卡住
	//本质是根据 chan 中是否有消息 来控制的！！！
	go func() {
		//在chan中的消息被消费后(非满信道)，定时向chan中写入新的消息
		//通过控制chan中新消息的写入频率 来控制消费频率(写入则立即消费)
		for t := range time.Tick(500 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	//模拟5个请求
	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)

	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("req:", req, time.Now())
	}

	//req: 1 2024-09-09 16:01:12.8606441 +0800 CST m=+0.516587101
	//-----500ms
	//req: 2 2024-09-09 16:01:13.3542816 +0800 CST m=+1.010224601
	//-----500ms
	//req: 3 2024-09-09 16:01:13.8564872 +0800 CST m=+1.512430201
	//-----500ms
	//req: 4 2024-09-09 16:01:14.3539376 +0800 CST m=+2.009880601
	//-----500ms
	//req: 5 2024-09-09 16:01:14.8503123 +0800 CST m=+2.506255301
	//------------
	//req: 1 2024-09-09 16:01:14.8505381 +0800 CST m=+2.506481101
	//req: 2 2024-09-09 16:01:14.8505381 +0800 CST m=+2.506481101
	//req: 3 2024-09-09 16:01:14.8505381 +0800 CST m=+2.506481101
	//-----500ms
	//req: 4 2024-09-09 16:01:15.3645753 +0800 CST m=+3.020518301
	//-----500ms
	//req: 5 2024-09-09 16:01:15.8567247 +0800 CST m=+3.512667701
}
