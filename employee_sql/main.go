package main

import (
	"database/sql"
	"net/http"

	"log"

	"github.com/gorilla/mux"

	"github.com/SanjeevChoubey/RESTapi/employee_sql/controllers"
	"github.com/SanjeevChoubey/RESTapi/employee_sql/driver"
)

var db *sql.DB

func logFatal(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	router := mux.NewRouter()
	db = driver.GetConnection()
	controller := controllers.Controller{}
	router.HandleFunc("/employees", controller.GetEmployees(db)).Methods("GET")
	router.HandleFunc("/employees/{id}", controller.GetEmployee(db)).Methods("GET")
	router.HandleFunc("/employees", controller.AddEmployee(db)).Methods("POST")
	router.HandleFunc("/employees", controller.UpdateEmployee(db)).Methods("PUT")
	router.HandleFunc("/employees/{id}", controller.DeleteEmployee(db)).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}
