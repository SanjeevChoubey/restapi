package bookRepository

import (
	"database/sql"
	"log"

	"github.com/SanjeevChoubey/RESTapi/BookList_Db/models"
)

type BookRepository struct{}

func logFatal(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func (b *BookRepository) GetBooks(db *sql.DB, book models.Book, books []models.Book) []models.Book {
	// Fetch rows from data base
	rows, err := db.Query("Select * from books")
	logFatal(err)
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Title)
		logFatal(err)
		books = append(books, book)
	}
	return books

}

func (b *BookRepository) GetBook(db *sql.DB, book models.Book, bookID int) models.Book {
	rows := db.QueryRow("Select * from books  where id = $1", bookID)
	err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
	logFatal(err)
	return book
}

func (b *BookRepository) AddBook(db *sql.DB, book models.Book) int {
	err := db.QueryRow("Insert into books (title,author,year) values($1,$2,$3) returning id",
		&book.Title, &book.Author, &book.Year).Scan(&book.ID)
	logFatal(err)

	return book.ID
}

func (b *BookRepository) UpdateBook(db *sql.DB, book models.Book) int64 {
	result, err := db.Exec("Update books set title = $1, author = $2, year = $3 where id = $4 returning id",
		&book.Title, &book.Author, &book.Year, &book.ID)
	logFatal(err)
	rowsAffected, err := result.RowsAffected()
	logFatal(err)
	return rowsAffected
}

func (b *BookRepository) DeleteBook(db *sql.DB, bookID int) bool {
	_, err := db.Exec("Delete from books where id = $1", bookID)
	logFatal(err)
	return true
}
