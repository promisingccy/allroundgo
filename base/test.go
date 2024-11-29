package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	//a := 'a'

	//fmt.Println("Hello, World!", a)
	//traversalString()
	//
	//changeString()
	//array()
	//fmt.Println()
	//testArray()
	//
	//testSliceInit()
	//testSliceCreate()
	//sliceForEach()
	mapSort()
}

// [{"k":"v"},{"k":"v"}]
func slice4Map() {
	var mapSlice = make([]map[string]string, 3)
	for index, value := range mapSlice {
		fmt.Printf("index:%d value:%v\n", index, value)
	}
	fmt.Println("after init")
	// 对切片中的map元素进行初始化
	mapSlice[0] = make(map[string]string, 10)
	mapSlice[0]["name"] = "王五"
	mapSlice[0]["password"] = "123456"
	mapSlice[0]["address"] = "红旗大街"
	for index, value := range mapSlice {
		fmt.Printf("index:%d value:%v\n", index, value)
	}
}

// {"k":["v1","v2"]}
func map4key2Slice() {
	var sliceMap = make(map[string][]string, 3)
	fmt.Println(sliceMap)
	fmt.Println("after init")
	key := "中国"
	value, ok := sliceMap[key]
	if !ok {
		value = make([]string, 0, 2)
	}
	value = append(value, "北京", "上海")
	sliceMap[key] = value
	fmt.Println(sliceMap)
}

func mapSort() {
	rand.Seed(time.Now().UnixNano()) //初始化随机数种子

	var scoreMap = make(map[string]int, 200)

	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("stu%02d", i) //生成stu开头的字符串
		value := rand.Intn(100)          //生成0~99的随机整数
		scoreMap[key] = value
	}
	//取出map中的所有key存入切片keys
	var keys = make([]string, 0, 200)
	for key := range scoreMap {
		keys = append(keys, key)
	}
	//对切片进行排序
	sort.Strings(keys)
	//按照排序后的key遍历map
	for _, key := range keys {
		fmt.Println(key, scoreMap[key])
	}
}

func sliceForEach() {
	data := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	slice := data[:]
	for index, value := range slice {
		fmt.Printf("inde : %v , value : %v\n", index, value)
	}
}

var arr = [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
var slice0 []int = arr[2:8]
var slice1 []int = arr[0:6]        //可以简写为 var slice []int = arr[:end]
var slice2 []int = arr[5:10]       //可以简写为 var slice[]int = arr[start:]
var slice3 []int = arr[0:len(arr)] //var slice []int = arr[:]
var slice4 = arr[:len(arr)-1]      //去掉切片的最后一个元素
func testSliceInit() {
	fmt.Printf("全局变量：arr %v\n", arr)
	fmt.Printf("全局变量：slice0 arr[2:8] %v\n", slice0)
	fmt.Printf("全局变量：slice1 arr[0:6] %v\n", slice1)
	fmt.Printf("全局变量：slice2 arr[5:10] %v\n", slice2)
	fmt.Printf("全局变量：slice3 arr[0:len(arr)] %v\n", slice3)
	fmt.Printf("全局变量：slice4 arr[:len(arr)-1] %v\n", slice4)
	fmt.Printf("-----------------------------------\n")
	arr2 := [...]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	slice5 := arr[2:8]
	slice6 := arr[0:6]         //可以简写为 slice := arr[:end]
	slice7 := arr[5:10]        //可以简写为 slice := arr[start:]
	slice8 := arr[0:len(arr)]  //slice := arr[:]
	slice9 := arr[:len(arr)-1] //去掉切片的最后一个元素
	fmt.Printf("局部变量： arr2 %v\n", arr2)
	fmt.Printf("局部变量： slice5 arr[2:8] %v\n", slice5)
	fmt.Printf("局部变量： slice6 arr[0:6] %v\n", slice6)
	fmt.Printf("局部变量： slice7 arr[5:10]  %v\n", slice7)
	fmt.Printf("局部变量： slice8 arr[0:len(arr)] %v\n", slice8)
	fmt.Printf("局部变量： slice9 arr[:len(arr)-1] %v\n", slice9)
}

func testSliceCreate() {
	//1.声明切片
	var s1 []int
	if s1 == nil {
		fmt.Println("是空")
	} else {
		fmt.Println("不是空")
	}
	// 2.:=
	s2 := []int{}
	// 3.make()
	var s3 []int = make([]int, 0)
	fmt.Println(s1, s2, s3)
	// 4.初始化赋值
	var s4 []int = make([]int, 0, 0)
	fmt.Println(s4)
	s5 := []int{1, 2, 3}
	fmt.Println(s5)
	// 5.从数组切片
	arr := [5]int{1, 2, 3, 4, 5}
	var s6 []int
	// 前包后不包  左开右闭
	s6 = arr[1:4]
	fmt.Println(s6)
}

func printArr(arr *[5]int) {
	arr[0] = 10
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func testArray() {
	var arr1 [5]int
	printArr(&arr1)
	//[10 0 0 0 0]
	fmt.Println(arr1)
	arr2 := [...]int{2, 4, 6, 8, 10}
	printArr(&arr2)
	//[10 4 6 8 10]
	fmt.Println(arr2)
}

func traversalString() {
	s := "pprof.cn博客"
	for i := 0; i < len(s); i++ { //byte
		u := s[i]
		fmt.Printf("%v(%c) ", u, u)
	}
	fmt.Println()
	for _, r := range s { //rune
		i := r
		fmt.Printf("%v(%c) ", i, i)
	}
	fmt.Println()
}

func changeString() {
	s1 := "hello"
	// 强制类型转换
	byteS1 := []byte(s1)
	byteS1[0] = 'H'
	fmt.Println(string(byteS1)) //Hello

	s2 := "博客"
	runeS2 := []rune(s2)
	runeS2[0] = '狗'
	fmt.Println(string(runeS2)) //狗客
}

// -------------------------- array
var arr0 [5]int = [5]int{1, 2, 3}
var arr1 = [5]int{1, 2, 3, 4, 5}
var arr2 = [...]int{1, 2, 3, 4, 5, 6}
var str = [5]string{3: "hello world", 4: "tom"}

func array() {
	a := [3]int{1, 2}           // 未初始化元素值为 0。
	b := [...]int{1, 2, 3, 4}   // 通过初始化值确定数组长度。
	c := [5]int{2: 100, 4: 200} // 使用引号初始化元素。
	d := [...]struct {
		name string
		age  uint8
	}{
		{"user1", 10}, // 可省略元素类型。
		{"user2", 20}, // 别忘了最后一行的逗号。
	}
	fmt.Println(arr0, arr1, arr2, str)
	fmt.Println(a, b, c, d)
}
