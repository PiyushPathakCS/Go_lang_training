package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type response struct {
	code    int
	body    string
	message string
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			data := response{
				code:    http.StatusOK,
				body:    "Hello",
				message: "Welcome",
				
			}
			json.NewEncoder(w).Encode(data)
		}
	})
	fmt.Println("Welcome to api development ")
	http.ListenAndServe(":3000", mux)
}
