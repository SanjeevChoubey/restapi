package main

import (
	"log"
	"net/http"
      "os"
	_ "database/sql"
	_ "github.com/lib/pq"
	_ "github.com/subosito/gotenv"
	"github.com/gorilla/mux"
	
)

type Book struct {
	ID     int    `json:id`
	Title  string `json:title`
	Author string `json:author`
	Year   string `json:year`
}

var books []Book
var db *db.sql
func init(){
	gotenv.Load()
}

func logFatal(err error){
	if err != nil{
		log.Fatalln(err)
	}
	
}
func main() {
	// Connect to Database
	pgUrl ,err := pq.ParseURL(os.Getenv("ELEPHANTSQL_URL"))
	logFatal(err)
	/*
	-db name
	-host
	-password
	-port
	-user
	*/

	log.Println(pgUrl)



	router := mux.NewRouter()

	router.HandleFunc("/books", getBooks).Methods("GET")
	router.HandleFunc("/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/books", addBook).Methods("POST")
	router.HandleFunc("/books", updateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", deleteBook).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func getBooks(w http.ResponseWriter, r *http.Request) {

}

func getBook(w http.ResponseWriter, r *http.Request) {

}

func addBook(w http.ResponseWriter, r *http.Request) {

}

func updateBook(w http.ResponseWriter, r *http.Request) {

}

func deleteBook(w http.ResponseWriter, r *http.Request) {

}
