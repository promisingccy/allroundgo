package main

import (
	"errors"
	"fmt"
	"strconv"
	"sync"
	"time"
)

var lock = &sync.Mutex{}

// db start =========9
// db end =========9
// db start =========0
// db end =========0
// db start =========1
// db end =========1
// db start =========2
// db end =========2
// db start =========3
// db end =========3
// db start =========4
// db end =========4
// db start =========5
// db end =========5
// db start =========6
// db end =========6
// db start =========7
// db end =========7
// db start =========8
// db end =========8
// db start ----------0
// db end ----------0
// db start ----------1
// db end ----------1
// db start ----------2
// db end ----------2
// db start ----------3
// db end ----------3
// db start ----------4
// db end ----------4
// db start ----------5
// db end ----------5
// db start ----------9
// db end ----------9
// db start ----------6
// db end ----------6
// db start ----------7
// db end ----------7
// db start ----------8
// db end ----------8
func main() {
	fmt.Println(time.Now().Unix())
	fmt.Println("main----------")
	go ff1()
	go ff2()
	go ff1()
	go ff2()

	time.Sleep(60 * time.Second)
}

func ff1() {
	for i := 0; i < 10; i++ {
		go db("----------" + strconv.Itoa(i))
	}
}
func ff2() {
	for i := 0; i < 10; i++ {
		go db("=========" + strconv.Itoa(i))
	}
}

func db(name string) {
	lock.Lock()
	defer lock.Unlock()
	fmt.Println("db start", name)
	time.Sleep(1 * time.Second)
	errors.New("vvv")
	fmt.Println("db end", name)
	//使用 defer
	//lock.Unlock()
}

//db start ----------0
//db start =========6
//db start =========4
//db start =========5
//db start ----------9
//db start =========2
//db start =========8
//db start =========1
//db start =========0
//db start ----------4
//db start ----------1
//db start ----------2
//db start ----------3
//db start =========9
//db start ----------5
//db start ----------7
//db start ----------8
//db start =========7
//db start ----------6
//db start =========3
//db end =========3
//db end ----------6
//db end =========4
//db end =========9
//db end ----------3
//db end ----------2
//db end =========2
//db end ----------5
//db end ----------0
//db end ----------4
//db end =========6
//db end ----------7
//db end =========8
//db end =========0
//db end =========1
//db end ----------1
//db end =========5
//db end ----------8
//db end =========7
//db end ----------9
