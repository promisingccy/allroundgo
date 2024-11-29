package main

import "fmt"

func main() {
	// Index 可以在整数切片上使用
	si := []int{10, 20, 15, -10}
	fmt.Println(FindIndex(si, 15))

	// Index 也可以在字符串切片上使用
	ss := []string{"foo", "bar", "baz"}
	fmt.Println(FindIndex(ss, "hello"))
}

// FindIndex 适用于任何支持比较的类型
// func FindIndex(s []T, x T) int {
// func FindIndex[T comparable](s []T, x T) int {  //函数的类型参数出现在函数参数之前的方括号之间
// [T comparable] 类型约束  comparable 对任意满足该类型的值使用 == 和 != 运算符
func FindIndex[T comparable](s []T, x T) int {
	for i, v := range s {
		if v == x {
			return i
		}
	}
	return -1
}
