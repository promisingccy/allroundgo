package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	lock := sync.RWMutex{}

	go Read(&lock, "1")
	go Read(&lock, "2")
	go Write(&lock, "3")
	go Read(&lock, "4")
	go Write(&lock, "5")
	go Read(&lock, "6")
	go Read(&lock, "7")

	time.Sleep(20 * time.Second)
}

func Read(lock *sync.RWMutex, t string) {
	lock.RLock()
	fmt.Println(t, "---ReadPre ", time.Now())
	time.Sleep(time.Second)
	fmt.Println(t, "---Reading ", time.Now())
	lock.RUnlock()
	fmt.Println(t, "---ReadPost ", time.Now())
	return
}

func Write(lock *sync.RWMutex, t string) {
	lock.Lock()
	fmt.Println(t, "===WritePre ", time.Now())
	time.Sleep(time.Second)
	fmt.Println(t, "===Writing ", time.Now())
	lock.Unlock()
	fmt.Println(t, "===WritePost ", time.Now())
	return
}
