package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url string = "https://www.youtube.com/"

func main() {
	fmt.Println("Welcome to web handling")

	response, err := http.Get(url)

	if err != nil {
		fmt.Println("err occured")
	} else {

		re, _ := ioutil.ReadAll(response.Body)

		fmt.Printf("Response is of type %T\n", response)

		defer response.Body.Close()

		fmt.Println(string(re))

	}

}
