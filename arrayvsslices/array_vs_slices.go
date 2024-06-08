package main

import "fmt"

func main() {
	arr1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	arr2 := arr1

	a1 := [9]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	a2 := a1

	arr2[5] = 10

	a2[5] = 10

	fmt.Println(arr1)

	fmt.Println(a1)
}
