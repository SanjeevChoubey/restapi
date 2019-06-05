package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/SanjeevChoubey/RESTapi/BookList_Db/controllers"
	"github.com/SanjeevChoubey/RESTapi/BookList_Db/driver"
	"github.com/SanjeevChoubey/RESTapi/BookList_Db/models"
	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
)

var books []models.Book
var db *sql.DB

func init() {
	gotenv.Load()
}

func logFatal(err error) {
	if err != nil {
		log.Fatalln(err)
	}

}
func main() {
	db = driver.ConnectDB()
	router := mux.NewRouter()
	controller := controllers.Controller{}
	router.HandleFunc("/books", controller.GetBooks(db)).Methods("GET")
	router.HandleFunc("/books/{id}", controller.GetBook(db)).Methods("GET")
	router.HandleFunc("/books", controller.AddBook(db)).Methods("POST")
	router.HandleFunc("/books", controller.UpdateBook(db)).Methods("PUT")
	router.HandleFunc("/books/{id}", controller.DeleteBook(db)).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))
}
