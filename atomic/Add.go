package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	SimpleAdd()
	LockAdd()
	AtomicAdd()
	// SimpleAdd time cost: 4.9267ms, count: 9800
	// LockAdd time cost: 4.4213ms, count: 10000
	// AtomicAdd time cost: 4.2684ms, count: 10000
}

var simpleCount = 0

func SimpleAdd() {
	wg := sync.WaitGroup{}
	start := time.Now()
	for _ = range 10000 {
		wg.Add(1)
		go func() {
			simpleCount++
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("")
	fmt.Printf("SimpleAdd time cost: %v, count: %d", time.Since(start), simpleCount)
}

var LockCount = 0

func LockAdd() {
	wg := sync.WaitGroup{}
	lock := sync.Mutex{}
	start := time.Now()
	for _ = range 10000 {
		wg.Add(1)
		go func() {
			lock.Lock()
			LockCount++
			lock.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("")
	fmt.Printf("LockAdd time cost: %v, count: %d", time.Since(start), LockCount)
}

var AtomicCount int64 = 0

func AtomicAdd() {
	wg := sync.WaitGroup{}
	start := time.Now()
	for _ = range 10000 {
		wg.Add(1)
		go func() {
			atomic.AddInt64(&AtomicCount, 1)
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("")
	fmt.Printf("AtomicAdd time cost: %v, count: %d", time.Since(start), AtomicCount)
}
