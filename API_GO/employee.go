package main

import "gorm.io/gorm"

type Employee struct {
	gorm.Model
	EmpName   string  `json:"emp_name"`
	EmpSalary float64 `json:"salary"`
	Email     string  `json:"email"`
}
