package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	fmt.Println(time.Now().Format("01"))
	fmt.Println(time.Now().Format("02"))
	fmt.Println(time.Now().String(), "main start")

	lock := &sync.RWMutex{}
	lock.Lock()
	//-----

	for i := 0; i < 4; i++ {
		go func(i int) {
			fmt.Println(time.Now().String(), "goroutine", i, "--start")
			lock.RLock()
			//---
			fmt.Println(time.Now().String(), "goroutine", i, "--process")
			time.Sleep(1 * time.Second)
			//---
			lock.RUnlock()
		}(i)
	}

	time.Sleep(time.Second * 2)
	//-----
	lock.Unlock()

	//-----

	lock.Lock()
	fmt.Println(time.Now().String(), "main exit")
	lock.Unlock()

	//-----

}
