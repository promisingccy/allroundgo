package main

import "fmt"

func main() {

	fmt.Printf("%v\n", delete())
	base()
}

func delete() []int {
	a := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	//delete index at i
	toDeleteIndex := 2
	// a = append(a[:i], a[i+1:]...)
	a = append(a[:toDeleteIndex], a[toDeleteIndex+1:]...)
	return a
}

func base() {
	a := [5]int{1, 2, 3, 4, 5}
	//a[low : high : max]
	//索引的范围定义：左闭右开 [low,high)
	s := a[1:3] // s := a[low:high] --- 2 3
	fmt.Printf("%v\n", s)
}
