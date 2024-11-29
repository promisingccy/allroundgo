package main

import (
	"fmt"
	"sync"
)

// EntWithLock 希望从多个 goroutine 同时更新它，
// 因此我们添加了一个 互斥锁Mutex 来同步访问。
// 请注意不能复制互斥锁，如果需要传递这个 struct，应使用指针完成。
type EntWithLock struct {
	lock sync.Mutex
	data map[string]int
}

// 在循环中递增对 name 的计数
func (c *EntWithLock) inc(name string) {
	//在访问 counters 之前锁定互斥锁；
	c.lock.Lock()
	//使用 [defer]（defer） 在函数结束时解锁。
	defer c.lock.Unlock()
	//在循环中递增对 name 的计数
	c.data[name]++
}

func main() {
	c := EntWithLock{
		//请注意，互斥量的零值是可用的，因此这里不需要初始化 lock
		data: map[string]int{"a": 0, "b": 0},
	}

	var wg sync.WaitGroup

	doInc := func(name string, n int) {
		for i := 0; i < n; i++ {
			//在循环中递增对 name 的计数
			c.inc(name)
		}
		wg.Done()
	}

	//设置需要等待同步完成的 任务/协程数量为 3
	wg.Add(3)

	//同时运行多个 goroutines;
	//请注意，它们都访问相同的 Container，其中两个访问相同的计数器。
	go doInc("a", 10000)
	go doInc("a", 10000)
	go doInc("b", 10000)
	wg.Wait()

	fmt.Println(c.data)
	// map[a:20000 b:10000]
}
