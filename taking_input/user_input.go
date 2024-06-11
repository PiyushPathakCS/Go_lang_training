package main

import (
	"fmt"
)

type inter interface {
	isEven()
}

type number struct {
	n int
}

func (v *number) num(n int) {
	v.n = n
	fmt.Println("Struct value is :", v.n)
}

func print(n *int) {
	fmt.Println("value received in function")
	fmt.Println(*n)
	*n = 80

	//print(n)
}

func (v *number) isEven() bool {
	if v.n%2 != 0 {
		return false
	}

	return true
}
func mps(m map[int]string) map[int]string {
	var y int
	var k int
	var v string
	fmt.Println("Enter the number of key value Pair to be enterd in maps")
	fmt.Scan(&y)
	fmt.Println("Enter the key value Pair to be enterd in maps")

	for i := 1; i <= y; i++ {

		fmt.Scan(&k)
		fmt.Scan(&v)
		m[k] = v

	}
	return m

}
func main() {
	fmt.Println("Enter the value")
	var n int
	fmt.Scan(&n)
	fmt.Printf("Entered value is : %v\n", n)
	print(&n)
	fmt.Printf("Value changed by function is : %v\n", n)
	v := number{}
	fmt.Println("Enter the value to be entered in the struct:")
	var x int
	fmt.Scan(&x)
	v.num(x)
	fmt.Println("Is the number int he struct even or not :", v.isEven())

	fmt.Println("implementing maps")
	m := make(map[int]string)
	t := mps(m)
	var id_slc []int
	var name_slc []string
	fmt.Println("Printing maps key value pair")

	fmt.Println("Map values for variable recived from function")
	for k, v := range t {
		id_slc = append(id_slc, k)
		name_slc = append(name_slc, v)
	}

	fmt.Println("Map values in variable m")
	for k, v := range m {
		fmt.Println(k, v)
	}

}
