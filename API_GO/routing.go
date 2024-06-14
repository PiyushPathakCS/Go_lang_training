package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandlerRouting() {

	r := mux.NewRouter()
	r.HandleFunc("/employee", GetEmployees).Methods("GET")
	r.HandleFunc("/employee/{empid}", GetEmployeeByID).Methods("GET")
	r.HandleFunc("/employee", CreateEmployee).Methods("POST")
	r.HandleFunc("/employeeupdate/{emp_update}", UpdateEmployee).Methods("PUT")
	r.HandleFunc("/employeedelete/{emp_delete}", DeleteEmployee).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", r))

}
