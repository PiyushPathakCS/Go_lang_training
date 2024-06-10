package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func check(ch chan int) {
	for val := range ch {
		fmt.Println(val)
	}

	defer wg.Done()
}

func main() {
	ch := make(chan int, 2)

	wg.Add(1)

	// go check(ch)

	ch <- 1
	ch <- 2

	go check(ch)

	close(ch)

	defer wg.Wait()

	// select {
	// default:
	// 	fmt.Println("Code Completed")
	// }

}
