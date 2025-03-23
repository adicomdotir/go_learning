package main

import (
	"fmt"
	"sync"
)

func main() {
	// var ch chan int
	fmt.Println("main start")
	var wg sync.WaitGroup
	var mut sync.Mutex
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go addNumber(&wg, &mut)
	}
	wg.Wait()
	fmt.Println("x =", x)
}

var x int = 0

func addNumber(wg *sync.WaitGroup, mut *sync.Mutex) {
	mut.Lock()
	x++
	mut.Unlock()
	wg.Done()
}
