package main

import "fmt"

func data_in1(ch chan int, arr3 *[]int) {
	*arr3 = append(*arr3, <-ch)
}

func main() {

	arr1 := []int{1, 5, 6, 9, 10}
	arr2 := []int{3, 7, 8}

	arr3 := []int{}

	pt := &arr3

	ch := make(chan int)

	i := 0
	j := 0

	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {

			go data_in1(ch, pt)
			ch <- arr1[i]
			i++
		} else {

			go data_in1(ch, pt)
			ch <- arr2[j]
			j++
		}
	}

	fmt.Println(i, j)
	fmt.Println(len(arr1), len(arr2))

	if i == len(arr1) {
		for j < len(arr2) {
			go data_in1(ch, pt)
			ch <- arr2[j]
			j++
		}
	} else {
		for i < len(arr1) {
			go data_in1(ch, pt)
			ch <- arr1[i]
			i++
		}

	}

	fmt.Println("helloooo", arr3)

	select {

	default:
		fmt.Println("done")
	}

}
