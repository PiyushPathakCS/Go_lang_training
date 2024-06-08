package main

import "fmt"

func main() {
	v := make(map[string][]string)
	v["fname"] = []string{"ramesh", "rajesh"}
	fmt.Print(v)
}
