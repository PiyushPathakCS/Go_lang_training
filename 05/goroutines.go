package main

import "fmt"

func pr(n int) {

	for i := 1; i <= n; i++ {
		fmt.Println(i)
	}

}

func main() {

	fmt.Println("welcome to goroutine")
	go pr(15)

	func(n int) {

		for i := 1; i <= n; i++ {
			fmt.Println(i)
		}

	}(10)

	fmt.Println("Goroutines working")

	select {}
}
