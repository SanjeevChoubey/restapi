package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/lib/pq"
	"github.com/subosito/gotenv"
)

type Book struct {
	ID     int    `json:id`
	Title  string `json:title`
	Author string `json:author`
	Year   string `json:year`
}

var books []Book
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
	// Connect to Database
	pgURL, err := pq.ParseURL(os.Getenv("ELEPHANTSQL_URL"))
	logFatal(err)
	/*
		-db name
		-host
		-password
		-port
		-user
	*/
	db, err = sql.Open("postgres", pgURL)
	logFatal(err)
	//defer db.Close()

	err = db.Ping()
	logFatal(err)

	log.Println(pgURL)

	router := mux.NewRouter()

	router.HandleFunc("/books", getBooks).Methods("GET")
	router.HandleFunc("/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/books", addBook).Methods("POST")
	router.HandleFunc("/books", updateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", deleteBook).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func getBooks(w http.ResponseWriter, r *http.Request) {

	var book Book
	books = []Book{}

	rows, err := db.Query("SELECT * FROM books")
	logFatal(err)
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
		logFatal(err)
		books = append(books, book)
	}

	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	var book Book
	parms := mux.Vars(r)
	rows := db.QueryRow("Select * from books where id = $1", parms["id"])
	err := rows.Scan(&book.ID, &book.Author, &book.Title, &book.Year)
	logFatal(err)
	json.NewEncoder(w).Encode(book)
}

func addBook(w http.ResponseWriter, r *http.Request) {
	var book Book
	var bookID int

	json.NewDecoder(r.Body).Decode(&book)

	log.Println(book)
	err := db.QueryRow("Insert into books (title,author, year) values($1,$2,$3) Returning id;",
		book.Title, book.Author, book.Year).Scan(&bookID)
	logFatal(err)
	json.NewEncoder(w).Encode(bookID)

}

func updateBook(w http.ResponseWriter, r *http.Request) {
	var book Book
	// Get from Json format into struct
	json.NewDecoder(r.Body).Decode(&book)
	result, err := db.Exec("Update books set title=$1, author=$2 , year=$3 where id=$4 RETURNING id",
		&book.Title, &book.Author, &book.Year, &book.ID)
	logFatal(err)
	rowsCount, err := result.RowsAffected()

	json.NewEncoder(w).Encode(rowsCount)
}

func deleteBook(w http.ResponseWriter, r *http.Request) {

	parms := mux.Vars(r)
	result, err := db.Exec("Delete from books where id = $1;", parms["id"])
	logFatal(err)
	// Count the number of rows affected
	count, err := result.RowsAffected()
	logFatal(err)
	// Write into json format
	json.NewEncoder(w).Encode(count)

}
