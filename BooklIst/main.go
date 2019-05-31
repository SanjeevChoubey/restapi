package main 

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Book struct {
	ID     int    `json:id`
	Title  string `json:title`
	Author string `json:author`
	Year   string `json:year`
}

var books []Book

func main() {
	router := mux.NewRouter()
	books = append(books,
		Book{ID: 1, Title: "Goroutines", Author: "Mr goRoutine", Year: "2011"},
		Book{ID: 2, Title: "Channels", Author: "Mr Channels", Year: "2012"},
		Book{ID: 3, Title: "Concurrency", Author: "Mr Concurrency", Year: "2012"},
		Book{ID: 4, Title: "Parallelism", Author: "Mr parallelism", Year: "2013"},
		Book{ID: 5, Title: "Rest Api", Author: "Mr Rest Api", Year: "2019"},
		Book{ID: 6, Title: "AWS", Author: "Mr AWS", Year: "1996"},
	)
	router.HandleFunc("/books", getBooks).Methods("GET")
	router.HandleFunc("/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/books", addBook).Methods("POST")
	router.HandleFunc("/books", updateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", deleteBook).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	parms := mux.Vars(r)go get gopkg.in/mgo.v2
	i, _ := strconv.Atoi(parms["id"])

	for _, book := range books {
		if book.ID == i {
			json.NewEncoder(w).Encode(book)
		}
	}
}

func addBook(w http.ResponseWriter, r *http.Request) {
	var book Book
	// since it is post method hence data will be passed from body
	_ = json.NewDecoder(r.Body).Decode(&book)
	books = append(books, book)
	json.NewEncoder(w).Encode(&books)

}

func updateBook(w http.ResponseWriter, r *http.Request) {
	log.Println("update a book")
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	log.Println("Delete a keyed book")
}
