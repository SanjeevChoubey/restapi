package controllers

import (
	"strconv"
	"github.com/gorilla/mux"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/SanjeevChoubey/RESTapi/employee_sql/models"
	repository "github.com/SanjeevChoubey/RESTapi/employee_sql/repository/employee"

	_ "github.com/go-sql-driver/mysql"
)

type Controller struct{}

func logFatal(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func (c *Controller) GetEmployees(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var employee models.Employee
		var employees []models.Employee
		repository := repository.Repository{}
		employees = repository.GetEmployees(db, employee, employees)
		json.NewEncoder(w).Encode(employees)
	}
}

func (c *Controller) GetEmployee(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var employee models.Employee
		parms := mux.Vars(r)
		id,err := strconv.Atoi(parms["id"])
		logFatal(err)
		repository := repository.Repository{}
		employee = repository.GetEmployee(db, employee, id)
		json.NewEncoder(w).Encode(employee)
	}
}

func (c *Controller) AddEmployee(db *sql.DB) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		var employee models.Employee
		json.NewDecoder(r.Body).Decode(&employee)
		repo := repository.Repository{}
		rowsaffected:= repo.AddEmployee(db,employee)
		json.NewEncoder(w).Encode(rowsaffected)
	}
}

func (c *Controller) UpdateEmployee(db *sql.DB) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		var employee models.Employee
		json.NewDecoder(r.Body).Decode(&employee)
		repo := repository.Repository{}
		rowsaffected:= repo.UpdateEmployee(db,employee)
		json.NewEncoder(w).Encode(rowsaffected)
	}
}


func (c *Controller) DeleteEmployee(db *sql.DB) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		parms := mux.Vars(r)
		id,err := strconv.Atoi(parms["id"])
		logFatal(err)
		repo := repository.Repository{}
		rowsaffected:= repo.DeleteEmployee(db,id)
		json.NewEncoder(w).Encode(rowsaffected)
	}
}