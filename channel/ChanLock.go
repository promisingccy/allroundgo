package main

import (
	"fmt"
	"time"
)

func main() {
	pipeline := make(chan bool, 1)

	var x int
	for i := 0; i < 1000; i++ {
		go func(ch chan bool, x *int) {
			//当信道里的数据量已经达到设定的容量时，此时再往里发送数据会阻塞整个程序。
			ch <- true
			*x = *x + 1
			<-ch
		}(pipeline, &x)
	}

	time.Sleep(time.Second)
	fmt.Println(x)
	//1000
}
