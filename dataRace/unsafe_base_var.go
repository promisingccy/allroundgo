package main

import (
	"fmt"
	"os"
	"sync/atomic"
	"time"
)

func main() {

}

// 数据竞争同样会发生在基原类型的变量上（如 bool、int、 int64 等），就像下面这样：
// 甚至“无辜”的数据竞争也会导致难以调试的问题：
// (1) 非原子性的内存访问
// (2) 编译器优化的干扰以及
// (3) 进程内存访问的重排序问题。
// 对此，典型的解决方案就是使用信道或互斥锁。要保护无锁的行为，一种方法就是使用 sync/atomic 包。
type unsafeWatchdog struct{ last int64 }

func (w *unsafeWatchdog) KeepAlive() {
	w.last = time.Now().UnixNano() // 第一个冲突的访问。
}

func (w *unsafeWatchdog) Start() {
	go func() {
		for {
			time.Sleep(time.Second)
			// 第二个冲突的访问。
			if w.last < time.Now().Add(-10*time.Second).UnixNano() {
				fmt.Println("No keepalives for 10 seconds. Dying.")
				os.Exit(1)
			}
		}
	}()
}

// ----------------
// 对此，典型的解决方案就是使用信道或互斥锁。要保护无锁的行为，一种方法就是使用 sync/atomic 包。
// - 读取 Load
// - 写入 Store
// - 交换 Swap
// - 比较并交换 CompareAndSwap
// - 增减 Add
type safeWatchdog struct{ last int64 }

func (w *safeWatchdog) KeepAlive() {
	atomic.StoreInt64(&w.last, time.Now().UnixNano())
}

func (w *safeWatchdog) Start() {
	go func() {
		for {
			time.Sleep(time.Second)
			if atomic.LoadInt64(&w.last) < time.Now().Add(-10*time.Second).UnixNano() {
				fmt.Println("No keepalives for 10 seconds. Dying.")
				os.Exit(1)
			}
		}
	}()
}
