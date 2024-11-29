package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var ops uint64

	//WaitGroup 帮助我们等待所有协程完成它们的工作
	var wg sync.WaitGroup

	//启动 50 个协程，并且每个协程会将计数器递增 1000 次。
	for i := 0; i < 50; i++ {
		wg.Add(1)

		go func() {
			for c := 0; c < 1000; c++ {
				atomic.AddUint64(&ops, 1)
			}
			wg.Done()
		}()
	}

	//等待，直到所有协程完成
	wg.Wait()

	fmt.Println("ops: ", ops)
	fmt.Println("ops safe: ", atomic.LoadUint64(&ops))

	//ops:  50000
}
