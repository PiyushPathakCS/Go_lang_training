package main

import (
	"fmt"
	"sync"
)

var msg string

var wg sync.WaitGroup

func updateMessage(s string) {
	msg = s
	wg.Done()
}

func printMessage() {
	fmt.Println(msg)

}

func main() {
	msg = "Hello, world,this is experiment "

	wg.Add(1)
	go updateMessage("Hello, universe!")
	wg.Wait()
	printMessage()

	wg.Add(1)
	go updateMessage("Hello, cosmos!")
	wg.Wait()
	printMessage()

	wg.Add(1)
	go updateMessage("Hello, world!")
	wg.Wait()
	printMessage()

}
