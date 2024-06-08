/*package main

import "fmt"

func rev_int(x int) int {
	var a int = x
	var temp int
	var new int

	for a != 0 {
		temp = a % 10
		a = a / 10
		new = new * 10
		new = new + temp
	}

	return new
}

func main() {
	v := rev_int(456789)

	if v > 2147483647 || v < -2147483648 {
		fmt.Println("Out of Range")
	} else {
		fmt.Println(v)
	}
}*/

package main

import "fmt"

type vehicals interface {
	vehical_type()
}

type truck struct {
	wheels int
}

type car struct {
	wheels int
}

func (T truck) vehical_type() {
	fmt.Println(T.wheels)
	fmt.Println("Heavy Vehicals")
}

func (C car) vehical_type() {
	fmt.Println(C.wheels)
	fmt.Println("Light Vehicals")
}
func main() {
	fmt.Println()
	t := truck{wheels: 16}
	c := car{wheels: 4}

	t.vehical_type()
	c.vehical_type()
	var s = []int{}
	s = append(s, 100, 200, 300, 400)
	fmt.Print(s)
}
