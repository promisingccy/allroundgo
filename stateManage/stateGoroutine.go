package main

type ReadOp struct {
	key  int
	resp chan int
}

type WriteOp struct {
	key  int
	val  int
	resp chan bool
}

func main() {
	var readOps uint64
	var writeOps uint64

	reads := make(chan ReadOp)
	writes := make(chan WriteOp)

	go func() {
		var state = make(map[int]int)

		for {
			select {
			case read := <-reads:
				read.resp <- state[read.key]
			}
		}
	}()

}
