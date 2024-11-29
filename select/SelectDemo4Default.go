package main

import (
	"fmt"
	"time"
)

// 使用 default 实现非阻塞读写
// 这个代码中，使用了 default 分支来实现非阻塞的通道读取和写入操作。
// 在 select 语句中，如果有通道已经准备好进行读写操作，那么就会执行相应的分支。
// 但是如果没有任何通道准备好读写，那么就会执行 default 分支中的代码。
func main() {

	ch := make(chan int, 1)
	go func() {
		for i := 0; i <= 5; i++ {
			ch <- i
			time.Sleep(1 * time.Second)
		}
		close(ch)
	}()

	// 不带条件的 for 循环将一直重复执行，
	// 直到在循环体内使用了 break 或者 return 跳出循环。
	for {
		select {
		case v, ok := <-ch:
			if ok {
				fmt.Println(v)
			} else {
				ch = nil
				fmt.Println("channel closed")
			}
		default:
			fmt.Println("no value ready")
			time.Sleep(500 * time.Millisecond)
		}
		if ch == nil {
			//跳出循环
			fmt.Println("break for...")
			break
		}
	}
	//执行结果（每次执行程序打印的顺序都不一致）：
	//no value ready
	//0
	//no value ready
	//1
	//no value ready
	//no value ready
	//2
	//no value ready
	//no value ready
	//3
	//no value ready
	//no value ready
	//4
	//no value ready
	//no value ready
	//5
	//no value ready
	//no value ready
	//channel closed
	//break for...
}
