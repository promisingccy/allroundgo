学会 Go select 语句，轻松实现高效并发原创


##前言

在 Go 语言中，Goroutine 和 Channel 是非常重要的并发编程概念，它们可以帮助我们解决并发编程中的各种问题。
关于它们的基本概念和用法，前面的文章 一文初探 Goroutine 与 channel 中已经进行了介绍。
而本文将重点介绍 select，它是协调多个 channel 的桥梁。

##select 介绍

##什么是 select

select 是 Go 语言中的一种控制结构，用于在多个通信操作中选择一个可执行的操作。
它可以协调多个 channel 的读写操作，使得我们能够在多个 channel 中进行非阻塞的数据传输、同步和控制。

为什么需要 select

Go 语言中的 select 语句是一种用于多路复用通道的机制，它允许在多个通道上等待并处理消息。
相比于简单地使用 for 循环遍历通道，使用 select 语句能够更加高效地管理多个通道。

以下是一些 select 语句的使用场景：

    等待多个通道的消息（多路复用）
      当我们需要等待多个通道的消息时，使用 select 语句可以非常方便地等待这些通道中的任意一个通道有消息到达，
        从而避免了使用多个goroutine进行同步和等待。
    超时等待通道消息
      当我们需要在一段时间内等待某个通道有消息到达时，使用 select 语句可以与 time 包结合使用实现定时等待。
    在通道上进行非阻塞读写
      在使用通道进行读写时，如果通道没有数据，读操作或写操作将会阻塞。
        但是使用 select 语句结合 default 分支可以实现非阻塞读写，从而避免了死锁或死循环等问题。

因此，select 的主要作用是在处理多个通道时提供了一种高效且易于使用的机制，简化了多个 goroutine 的同步和等待，使程序更加可读、高效和可靠。

##select 基础

select {
    case <- channel1:
    // channel1准备好了
    case data := <- channel2:
    // channel2准备好了，并且可以读取到数据data
    case channel3 <- data:
    // channel3准备好了，并且可以往其中写入数据data
    default:
    // 没有任何channel准备好了
}

其中， <- channel1 表示读取 channel1 的数据，
data <- channel2 表示用 data 去接收数据；
channel3 <- data 表示往 channel3 中写入数据。

select 的语法形式类似于 switch，但是它只能用于 channel 操作。
在 select 语句中，我们可以定义多个 case，每个 case 都是一个 channel 操作，用于读取或写入数据。
如果有多个 case 同时可执行，则会随机选择其中一个。
如果没有任何可执行的 case，则会执行 default 分支（如果存在），或者阻塞等待直到至少有一个 case 可执行为止。


func main() {
    ch1 := make(chan int)
    ch2 := make(chan int)

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- 1
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- 2
	}()
	for i := 0; i < 2; i++ {
		select {
		case data, ok := <-ch1:
			if ok {
				fmt.Println("从 ch1 接收到数据:", data)
			} else {
				fmt.Println("通道已被关闭")
			}
		case data, ok := <-ch2:
			if ok {
				fmt.Println("从 ch2接收到数据: ", data)
			} else {
				fmt.Println("通道已被关闭")
			}
		}
	}

	select {
	case data, ok := <-ch1:
		if ok {
			fmt.Println("从 ch1 接收到数据:", data)
		} else {
			fmt.Println("通道已被关闭")
		}
	case data, ok := <-ch2:
		if ok {
			fmt.Println("从 ch2接收到数据: ", data)
		} else {
			fmt.Println("通道已被关闭")
		}
	default:
		fmt.Println("没有接收到数据，走 default 分支")
	}
}



执行结果

从 ch1 接收到数据: 1
从 ch2接收到数据:  2
没有接收到数据，走 default 分支


上述示例中，首先创建了两个 channel，ch1 和 ch2，分别在不同的 goroutine 中向两个 channel 中写入数据。
然后，在主 goroutine 中使用 select 语句监听两个channel，一旦某个 channel 上有数据流动，就打印出相应的数据。
由于 ch1 中的数据比 ch2 中的数据先到达，因此首先会打印出 "从 ch1 接收到数据: 1"，然后才打印出 "从 ch2接收到数据:  2"。

为了方便测试 default 分支，我写了两个 select 代码块，执行到第二个 select 代码块的时候，
由于 ch1 和 ch2 都没有数据了，因此执行 default 分支，打印 "没有接收到数据，走 default 分支"。


###一些使用 select 与 channel 结合的场景###


###实现超时控制

func main() {
    ch := make(chan int)

    go func() {
        time.Sleep(3 * time.Second)
        ch <- 1
    }()

	select {
	case data, ok := <-ch:
		if ok {
			fmt.Println("接收到数据: ", data)
		} else {
			fmt.Println("通道已被关闭")
		}
	case <-time.After(2 * time.Second):
		fmt.Println("超时了！")
	}
}

执行结果为：超时了！。

在这个例子中，程序将在 3 秒后向 ch 通道里写入数据，而我在 select 代码块里设置的超时时间为 2 秒，
如果在 2 秒内没有接收到数据，则会触发超时处理。


###实现多任务并发控制

func main() {
    ch := make(chan int)

	for i := 0; i < 10; i++ {
		go func(id int) {
			ch <- id
		}(i)
	}

	for i := 0; i < 10; i++ {
		select {
		case data, ok := <-ch:
			if ok {
				fmt.Println("任务完成：", data)
			} else {
				fmt.Println("通道已被关闭")
			}
		}
	}
}

执行结果（每次执行的顺序都会不一致）：

任务完成： 1
任务完成： 5
任务完成： 2
任务完成： 3
任务完成： 4
任务完成： 0
任务完成： 9
任务完成： 6
任务完成： 7
任务完成： 8

在这个例子中，启动了 10 个 goroutine 并发执行任务，并使用一个 channel 来接收任务的完成情况。
在主函数中，使用 select 语句监听这个 channel，每当接收到一个完成的任务时，就进行处理。


###监听多个通道的消息

func main() {
    ch1 := make(chan int)
    ch2 := make(chan int)

	// 开启 goroutine 1 用于向通道 ch1 发送数据
	go func() {
		for i := 0; i < 5; i++ {
			ch1 <- i
			time.Sleep(time.Second)
		}
	}()

	// 开启 goroutine 2 用于向通道 ch2 发送数据
	go func() {
		for i := 5; i < 10; i++ {
			ch2 <- i
			time.Sleep(time.Second)
		}
	}()

	// 主 goroutine 从 ch1 和 ch2 中接收数据并打印
	for i := 0; i < 10; i++ {
		select {
		case data := <-ch1:
			fmt.Println("Received from ch1:", data)
		case data := <-ch2:
			fmt.Println("Received from ch2:", data)
		}
	}

	fmt.Println("Done.")
}


执行结果（每次执行程序打印的顺序都不一致）：

Received from ch2: 5
Received from ch1: 0
Received from ch1: 1
Received from ch2: 6
Received from ch1: 2
Received from ch2: 7
Received from ch1: 3
Received from ch2: 8
Received from ch1: 4
Received from ch2: 9
Done.

该示例代码中，通过使用 select 多路复用，可以同时监听多个通道的数据，并避免了使用多个 goroutine 进行同步和等待的问题。


###使用 default 实现非阻塞读写

func main() {
    ch := make(chan int, 1)

	go func() {
		for i := 1; i <= 5; i++ {
			ch <- i
			time.Sleep(1 * time.Second)
		}
		close(ch)
	}()

	for {
		select {
		case val, ok := <-ch:
			if ok {
				fmt.Println(val)
			} else {
				ch = nil
			}
		default:
			fmt.Println("No value ready")
			time.Sleep(500 * time.Millisecond)
		}
		if ch == nil {
			break
		}
	}
}

执行结果（每次执行程序打印的顺序都不一致）：

No value ready
1
No value ready
2
No value ready
No value ready
3
No value ready
No value ready
4
No value ready
No value ready
5
No value ready
No value ready

这个代码中，使用了 default 分支来实现非阻塞的通道读取和写入操作。
在 select 语句中，如果有通道已经准备好进行读写操作，那么就会执行相应的分支。
但是如果没有任何通道准备好读写，那么就会执行 default 分支中的代码。

select 的注意事项

以下是关于 select 语句的一些注意事项：

    select 语句只能用于通信操作，如 channel 的读写，不能用于普通的计算或函数调用。
    select 语句会阻塞，直到至少有一个 case 语句满足条件。
    如果有多个 case 语句满足条件，则会随机选择一个执行。
    如果没有 case 语句满足条件，并且有 default 语句，则会执行 default 语句。
    在 select 语句中使用 channel 时，必须保证 channel 是已经初始化的。
    如果一个通道被关闭，那么仍然可以从它中读取数据，直到它被清空，此时会返回通道元素类型的零值和一个布尔值，指示通道是否已关闭。

总之，在使用 select 语句时，要仔细考虑每个 case 语句的条件和执行顺序，避免死锁和其他问题。
总结

本文主要介绍了 Go 语言中的 select 语句。
在文章中，首先介绍了 select 的基本概念，包括它是一种用于在多个通道之间进行选择的语句，以及为什么需要使用 select。

接下来，文章详细介绍了 select 的基础知识，包括语法和基础用法。
在语法方面，讲解了 select 语句的基本结构以及如何使用 case 子句进行通道选择。
在基础用法方面，介绍了如何使用 select 语句进行通道的读取和写入操作，并讲解了一些注意事项。

在接下来的内容中，文章列举了一些使用 select 与 channel 结合的场景。
这些场景包括实现超时控制、实现多任务并发控制、监听多个通道的消息以及使用 default 实现非阻塞读写。
对于每个场景，文章都详细介绍了如何使用 select 语句实现。

最后，文章总结了 select 的注意事项，包括选择的通道必须是可读或可写的通道、
select 语句中的 case 子句必须是通道操作或者空的 default 子句，不能是其他类型的语句等等。


