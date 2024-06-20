package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandlerRouting() {

	r := mux.NewRouter()
	r.HandleFunc("/signup", Register).Methods("POST")
	r.HandleFunc("/login", Authenticate).Methods("POST")
	r.HandleFunc("/logout", Logout).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", r))

}
