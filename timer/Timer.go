package main

import (
	"fmt"
	"time"
)

func main() {
	TriggeredWithoutStop()
	TriggeredWithStopFail()
	TriggeredWithStopSuccess()

}

// TriggeredWithoutStop main 2024-11-28 15:13:15.5336125 +0800 CST m=+0.004432901
// timer2触发 2024-11-28 15:13:16.558141 +0800 CST m=+1.028961401
// main2024-11-28 15:13:17.5532201 +0800 CST m=+2.024040501
func TriggeredWithoutStop() {
	fmt.Println("main " + time.Now().String())
	t2 := time.NewTimer(time.Second)

	go func() {
		<-t2.C
		fmt.Println("timer2触发 " + time.Now().String())
	}()

	time.Sleep(2 * time.Second)
	fmt.Println("main" + time.Now().String())
}

// TriggeredWithStopFail main 2024-11-28 13:48:58.8209598 +0800 CST m=+0.004634901
// timer1触发 2024-11-28 13:49:00.8485713 +0800 CST m=+2.032246401
func TriggeredWithStopFail() {
	fmt.Println("main " + time.Now().String())
	//创建一个两秒后的定时器
	t1 := time.NewTimer(2 * time.Second)

	//等待触发 此处会阻塞直到定时器指定时间后才会往下走
	<-t1.C
	fmt.Println("timer1触发 " + time.Now().String())
	// timer1 已经被触发了，无法被取消了
	stop := t1.Stop()
	if stop {
		fmt.Println("timer 1 被取消了")
	}
}

// TriggeredWithStopSuccess main 2024-11-28 13:50:40.9042468 +0800 CST m=+0.004747701
// timer2被取消2024-11-28 13:50:40.9220809 +0800 CST m=+0.022581801
func TriggeredWithStopSuccess() {
	//如果你需要的仅仅是单纯的等待，使用 time.Sleep 就够了。
	//使用定时器的原因之一就是，你可以在定时器触发之前将其取消。
	//例如这样 创建一个等待1秒的定时器
	fmt.Println("main " + time.Now().String())
	t2 := time.NewTimer(time.Second)

	go func() {
		<-t2.C
		fmt.Println("timer2触发 " + time.Now().String())
	}()
	// 在触发前停止t2
	success := t2.Stop()
	if success {
		fmt.Println("timer2被取消" + time.Now().String())
	}
}
