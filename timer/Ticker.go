package main

import (
	"fmt"
	"time"
)

// 定时器 是当你想要在未来某一刻执行一次时使用的
// 打点器 则是为你想要以固定的时间间隔重复执行而准备的。
// 这里是一个打点器的例子，它将定时的执行，直到我们将它停止。
func main() {

	//打点器和定时器的机制有点相似：使用一个通道来发送数据。
	ticker := time.NewTicker(time.Second)

	done := make(chan bool)

	go func() {
		//这里我们使用通道内建的 select，等待每 1s 到达一次的值。
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	time.Sleep(3100 * time.Millisecond)
	//打点器可以和定时器一样被停止。
	//打点器一旦停止，将不能再从它的通道中接收到值。
	//我们将在运行 3100ms 后停止这个打点器。
	ticker.Stop()
	done <- true
	fmt.Println("ticker stopped")

	//当我们运行这个程序时，打点器会在我们停止它前打点 3 次。
	//Tick at 2024-08-27 16:55:40.0549565 +0800 CST m=+1.008553301
	//Tick at 2024-08-27 16:55:41.0656984 +0800 CST m=+2.019295201
	//Tick at 2024-08-27 16:55:42.0584419 +0800 CST m=+3.012038701
	//ticker stopped
}
