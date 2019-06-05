package driver

import (
	"database/sql"
	"log"
)

func logFatal(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

var db *sql.DB
var err error

func GetConnection() *sql.DB {

	db, err = sql.Open("mysql", "root:password@tcp(localhost:3306)/go_schema")
	logFatal(err)
	//defer db.Close()

	err = db.Ping() // after succesful operation we can use db instabnce for quering data
	logFatal(err)
	return db
}
