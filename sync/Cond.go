package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	fmt.Println("ForeachSignalDemo")
	fmt.Println(time.Now().String())
	ForeachSignalDemo()
	//fmt.Println("BroadCastDemo")
	//fmt.Println(time.Now().String())
	//BroadCastDemo()
	//fmt.Println("baseUsage")
	//baseUsage()

}

var done = false

// ForeachSignalDemo
// 2024-08-02 10:34:41.7253011 +0800 CST m=+0.004333401
// ========= 玄德公 升帐了 =============
// main-10sbe2024-08-02 10:34:41.7414163 +0800 CST m=+0.020448601
// 诸葛亮: 入帐，开始汇报进度~~~~~
// write4ForeachSignal-2sbe2024-08-02 10:34:41.7414163 +0800 CST m=+0.020448601
// c: 来了，等在门外
// d: 来了，等在门外
// a: 来了，等在门外
// b: 来了，等在门外
// write4ForeachSignal-2saf2024-08-02 10:34:43.7522596 +0800 CST m=+2.031291901
// 诸葛亮: 出帐： 下一个进
// c: 入账 ====》
// c: 出帐： 下一个进
// d: 入账 ====》
// d: 出帐： 下一个进
// a: 入账 ====》
// a: 出帐： 下一个进
// b: 入账 ====》
// b: 出帐： 下一个进
// main-10saf2024-08-02 10:34:51.7460036 +0800 CST m=+10.025035901
func ForeachSignalDemo() {
	cond := sync.NewCond(&sync.Mutex{})
	fmt.Println("========= 玄德公 升帐了 =============")
	names := []string{"a", "b", "c", "d"}
	for _, v := range names {
		go read4ForeachSignal(v, cond)
	}
	go write4ForeachSignal("诸葛亮", cond)
	fmt.Println("main-10sbe" + time.Now().String())
	time.Sleep(time.Second * 10)
	fmt.Println("main-10saf" + time.Now().String())
}

func read4ForeachSignal(name string, c *sync.Cond) {
	// 获取锁
	c.L.Lock()
	for !done {
		fmt.Println(name + ": 来了，等在门外")
		c.Wait()
	}
	fmt.Println(name + ": 入账 ====》")
	time.Sleep(time.Second)
	// 释放锁
	c.L.Unlock()
	fmt.Println(name + ": 出帐： 下一个进")
	// 通知下一个
	c.Signal()
}

func write4ForeachSignal(name string, c *sync.Cond) {
	fmt.Println(name + ": 入帐，开始汇报进度~~~~~")
	fmt.Println("write4ForeachSignal-2sbe" + time.Now().String())
	time.Sleep(time.Second * 2)
	fmt.Println("write4ForeachSignal-2saf" + time.Now().String())
	// 获取锁
	c.L.Lock()
	// 将标记设置为true
	done = true
	// 释放锁
	c.L.Unlock()
	fmt.Println(name + ": 出帐： 下一个进")
	// 通知下一个
	c.Signal()
}

// /---------------------------------

// BroadCastDemo
// 2024-08-02 10:15:32.8204026 +0800 CST m=+0.003784001
// ========= 玄德公 升帐了 =============
// main-3sbe2024-08-02 10:15:32.8375637 +0800 CST m=+0.020945101
// 诸葛亮: 入帐，作战计划制定中~~~~~
// write4Broadcast-2sbe2024-08-02 10:15:32.8375637 +0800 CST m=+0.020945101
// a: 来了，等在门外
// d: 来了，等在门外
// b: 来了，等在门外
// c: 来了，等在门外
// write4Broadcast-2saf2024-08-02 10:15:34.8399071 +0800 CST m=+2.023288501
// ======== 作战计划制定完毕 ========
// 诸葛亮: 大家进来开会吧！
// c: 入账 ====》
// a: 入账 ====》
// d: 入账 ====》
// b: 入账 ====》
// main-3saf2024-08-02 10:15:35.8446958 +0800 CST m=+3.028077201
func BroadCastDemo() {
	cond := sync.NewCond(&sync.Mutex{})
	fmt.Println("========= 玄德公 升帐了 =============")
	names := []string{"a", "b", "c", "d"}
	for _, v := range names {
		go read4Broadcast(v, cond)
	}
	go write4Broadcast("诸葛亮", cond)
	fmt.Println("main-3sbe" + time.Now().String())
	time.Sleep(time.Second * 3)
	fmt.Println("main-3saf" + time.Now().String())
}

func read4Broadcast(name string, c *sync.Cond) {
	// 获取锁
	c.L.Lock()
	for !done {
		fmt.Println(name + ": 来了，等在门外")
		c.Wait()
	}
	fmt.Println(name + ": 入账 ====》")
	// 释放锁
	c.L.Unlock()
}

func write4Broadcast(name string, c *sync.Cond) {
	fmt.Println(name + ": 入帐，作战计划制定中~~~~~")
	fmt.Println("write4Broadcast-2sbe" + time.Now().String())
	time.Sleep(time.Second * 2)
	fmt.Println("write4Broadcast-2saf" + time.Now().String())
	// 获取锁
	c.L.Lock()
	// 将标记设置为true
	done = true
	// 释放锁
	c.L.Unlock()
	fmt.Println("======== 作战计划制定完毕 ========")
	fmt.Println(name + ": 大家进来开会吧！")
	// 广播唤醒所有
	// Cond 的作用：用于等待一个或一组协程满足条件后唤醒这些协程。
	//c.Broadcast()
	// 单独提醒 随机有个一等待的人入账，之后就没有然后了。没有执行唤醒其他人。
	c.Signal()
}

// ---------------------------
func baseUsage() {

	// 创建实例
	cond := sync.NewCond(&sync.Mutex{})

	// 广播唤醒所有
	go func() {
		for range time.Tick(time.Millisecond) {
			cond.Broadcast()
		}
	}()

	// 唤醒一个协程
	go func() {
		for range time.Tick(time.Millisecond) {
			cond.Signal()
		}
	}()

	// 等待
	cond.L.Lock()
	condition := false

	for !condition {
		cond.Wait()
	}

	cond.L.Unlock()
}
