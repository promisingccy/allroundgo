package main

import "fmt"

type List[T any] struct {
	Next   *List[T]
	CurVal T
}

func main() {
	l1 := &List[int]{
		CurVal: 1,
	}
	l2 := &List[int]{
		CurVal: 2,
	}
	l2.Next = l1
	l3 := &List[int]{
		CurVal: 3,
	}
	l3.Next = l2
	l4 := &List[int]{
		CurVal: 4,
	}
	l4.Next = l3
	cur := l4
	for cur.Next != nil {
		fmt.Println(cur.CurVal)
		cur = cur.Next
	}
}
