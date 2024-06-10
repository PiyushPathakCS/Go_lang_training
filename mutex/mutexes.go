package main

import (
	"fmt"
	"sync"
)

var wg = &sync.WaitGroup{}

var mu = &sync.Mutex{}

func read(st []int) {
	for a := range st {
		fmt.Print(a)
	}

	fmt.Println()
	defer wg.Done()
}

func write(st *[]int) {

	*st = append(*st, 10)

	defer wg.Done()

}
func main() {
	st := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	wg.Add(3)

	fmt.Println("Before write operation")
	mu.Lock()
	go read(st)
	mu.Unlock()

	mu.Lock()
	go write(&st)
	mu.Unlock()

	fmt.Println("After write operation")
	mu.Lock()
	go read(st)
	mu.Unlock()

	defer wg.Wait()

}
