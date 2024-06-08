package main

import (
	"fmt"
)

func mtpy(ch chan int) {
	fmt.Println(100 * <-ch)
}

func main() {

	ch := make(chan int)
	go mtpy(ch)

	ch <- 9

	select {

	default:
		fmt.Println("done")
	}

}
