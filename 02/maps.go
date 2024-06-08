package main

import "fmt"

/*func main() {
	multiValueMap := make(map[string]map[string]string)
	multiValueMap["person"] = make(map[string]string)
	multiValueMap["person"]["firstName"] = "Albert"
	multiValueMap["person"]["lastName"] = "Don"
	multiValueMap["person"]["note"] = "Welcome"
	fmt.Println("Values for key 'person':", multiValueMap["person"])
}*/

func main() {

	dicenumber := 5

	switch dicenumber {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	default:
		fmt.Println(dicenumber)
	}
}
