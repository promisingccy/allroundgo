package main

import "fmt"

// 非阻塞通道操作
func main() {
	messages := make(chan string)
	signals := make(chan bool)

	//【非阻塞/接收】【读取】从信道接收到消息
	select {
	case msg := <-messages:
		fmt.Println("received msg:", msg)
	//default 实现非阻塞 代码会走到这个分支
	default:
		fmt.Println("no message received")
	}

	msg := "hi"
	//【非阻塞/发送】【写入】从信道接收到消息
	select {
	//因为messages是无缓冲区通道，并且也没有接收者
	//因此， default 会执行。
	case messages <- msg:
		fmt.Println("sent msg:", msg)
	//default 实现非阻塞 代码会走到这个分支
	default:
		fmt.Println("no message sent")
	}

	//【非阻塞】【接收】从信道接收到消息
	//多个 case 来实现 多路选择器
	select {
	//多路1
	case msg := <-messages:
		fmt.Println("received msg:", msg)
	//多路2
	case sig := <-signals:
		fmt.Println("received signal:", sig)
	//default 实现非阻塞 代码会走到这个分支
	default:
		fmt.Println("no activity")
	}
}

//no message received
//no message sent
//no activity
