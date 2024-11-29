package main

import (
	"fmt"
	"sync"
	"time"
)

// ===== 空城计 第 1 次 =====
// 【司马懿】吓跑了 ==》
// +++胜利+++
// ===== 空城计 第 2 次 =====
// ---失败---
// ===== 空城计 第 3 次 =====
// ---失败---
func main() {
	var result = false
	var once sync.Once
	f := func() {
		fmt.Println("【司马懿】吓跑了 ==》")
		result = true
	}

	for i := 1; i < 4; i++ {
		fmt.Printf("===== 空城计 第 %d 次 =====\n", i)
		once.Do(f)
		if result {
			fmt.Println("+++胜利+++")
		} else {
			fmt.Println("---失败---")
		}
		time.Sleep(time.Second)
		result = false

	}
}
