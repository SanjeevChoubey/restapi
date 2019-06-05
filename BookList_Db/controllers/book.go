package controllers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	bookRepository "github.com/SanjeevChoubey/RESTapi/BookList_Db/bookRepository/books"
	"github.com/SanjeevChoubey/RESTapi/BookList_Db/models"
	"github.com/gorilla/mux"
)

func logFatal(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

type Controller struct{}

var books []models.Book

func (c *Controller) GetBooks(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var book models.Book
		books = []models.Book{}
		bookRepository := bookRepository.BookRepository{}
		books = bookRepository.GetBooks(db, book, books)

		json.NewEncoder(w).Encode(books)
	}

}
func (c *Controller) GetBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.Book
		parms := mux.Vars(r)
		bookRepository := bookRepository.BookRepository{}

		id, err := strconv.Atoi(parms["id"])
		logFatal(err)

		book = bookRepository.GetBook(db, book, id)
		json.NewEncoder(w).Encode(book)
	}
}

func (c *Controller) AddBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.Book
		var bookID int

		json.NewDecoder(r.Body).Decode(&book)

		bookRepository := bookRepository.BookRepository{}
		bookID = bookRepository.AddBook(db, book)
		json.NewEncoder(w).Encode(bookID)

	}
}

func (c *Controller) UpdateBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.Book
		// Get from Json format into struct
		json.NewDecoder(r.Body).Decode(&book)
		bookRepository := bookRepository.BookRepository{}
		rowsCount := bookRepository.UpdateBook(db, book)
		json.NewEncoder(w).Encode(rowsCount)
	}
}

func (c *Controller) DeleteBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		parms := mux.Vars(r)
		bookRepository := bookRepository.BookRepository{}

		id, err := strconv.Atoi(parms["id"])
		logFatal(err)

		status := bookRepository.DeleteBook(db, id)
		// Write into json format
		json.NewEncoder(w).Encode(status)

	}
}
