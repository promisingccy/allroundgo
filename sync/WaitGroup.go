package main

import (
	"fmt"
	"sync"
	"time"
)

//定义等待组
// var wg sync.WaitGroup

// 等待组计数器加
//go.add(N)

// 计数器减
//wg.Done()

// 等待组等待计数结束
//wg.Wait()

func main() {
	var wg sync.WaitGroup
	times := 7
	wg.Add(times)

	//Async(times, wg)
	Sync(times, &wg)
	fmt.Println("===== ")
	wg.Wait()
	fmt.Println("====== END 孟获投降了 ======")
}

// Async 同步
// 孟获被抓了1次
// 孟获被抓了2次
// 孟获被抓了3次
// 孟获被抓了4次
// 孟获被抓了5次
// 孟获被抓了6次
// 孟获被抓了7次
// =====
// ====== END 孟获投降了 ======
func Async(times int, wg *sync.WaitGroup) {
	for i := 1; i <= times; i++ {
		fmt.Printf("孟获被抓了%d次\n", i)
		time.Sleep(time.Second)
		wg.Done()
	}
}

// Sync
// =====
// 孟获被抓了6次
// 孟获被抓了1次
// 孟获被抓了5次
// 孟获被抓了3次
// 孟获被抓了7次
// 孟获被抓了2次
// 孟获被抓了4次
// ====== END 孟获投降了 ======
func Sync(times int, wg *sync.WaitGroup) {
	for i := 1; i <= times; i++ {
		tmp := i
		go func(tmp int) {
			fmt.Printf("孟获被抓了%d次\n", tmp)
			wg.Done()
		}(tmp)
	}
}
