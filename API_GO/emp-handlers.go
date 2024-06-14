package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateEmployee(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json") //json type ma convert karna kan liya from text
	var emp Employee
	json.NewDecoder(r.Body).Decode(&emp)
	Database.Create(&emp)
	json.NewEncoder(w).Encode(emp)

}

func GetEmployees(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var emps []Employee
	Database.Find(&emps)
	json.NewEncoder(w).Encode(emps)

}

func GetEmployeeByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var empid Employee
	Database.First(&empid, mux.Vars(r)["eid"])
	json.NewEncoder(w).Encode(empid)

}

func UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var emp_update Employee
	Database.First(&emp_update, mux.Vars(r)["emp_update"])
	json.NewDecoder(r.Body).Decode(&emp_update)
	Database.Save(&emp_update)
	json.NewEncoder(w).Encode(emp_update)
}

func DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var emp_delete Employee
	Database.Delete(&emp_delete, mux.Vars(r)["emp_delete"])
	json.NewDecoder(r.Body).Decode(&emp_delete)
	json.NewEncoder(w).Encode("Employee Deleted Successfully")
}
