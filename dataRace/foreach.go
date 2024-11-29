package main

import (
	"fmt"
	"sync"
)

// 40321
// 40123
func main() {
	Demo1()
	fmt.Println()
	Demo2()
}

// 1
// 2
// 3
func Demo2() {
	arr := []int{1, 2, 3}
	newArr := []*int{}
	for _, v := range arr {
		newArr = append(newArr, &v)
	}
	for _, v := range newArr {
		fmt.Println(*v)
	}
}

func Demo1() {
	var wg sync.WaitGroup
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func(i int) {
			fmt.Print(i) // 你要找的不是“i”。
			wg.Done()
		}(i)
	}
	wg.Wait()
}
