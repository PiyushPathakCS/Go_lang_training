package main

import (
	"fmt"
	"net/http"
)

const url string = ""

func main() {
	fmt.Println("Welcome to web")

	response, err := http.Get(url)

	if err != nil {
		fmt.Println("err occured")
	}

	fmt.Printf("Response is of type %T\n", response)

	defer response.Body.Close()
}
