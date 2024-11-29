package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	//一个过期时间为 1s 的上下文
	//handle 函数没有进入超时的 select 分支，
	//但是 main 函数的 select 却会等待 context.Context 超时
	//并打印出 main context deadline exceeded。
	execCxt(1*time.Second, 500*time.Millisecond)
	//process time 500ms
	//main context deadline exceeded

	//处理时间已经超出上下文的存活时间
	//整个程序都会因为上下文的过期而被中止
	//多个 Goroutine 同时订阅 ctx.Done() 管道中的消息，
	//一旦接收到取消信号就立刻停止当前正在执行的工作。
	execCxt(1*time.Second, 1500*time.Millisecond)
	//main context deadline exceeded
	//handle context deadline exceeded
}

func execCxt(total time.Duration, process time.Duration) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), total)
	defer cancelFunc()

	go handle(ctx, process)
	select {
	case <-ctx.Done():
		fmt.Println("main", ctx.Err())
	}
}

func handle(ctx context.Context, duration time.Duration) {
	select {
	case <-ctx.Done():
		fmt.Println("handle", ctx.Err())
	case <-time.After(duration):
		fmt.Println("process time", duration)
	}
}
