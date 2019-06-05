package driver

import (
	"database/sql"
	"log"
	"os"

	"github.com/lib/pq"
)

var db *sql.DB

func logFatal(err error) {
	if err != nil {
		log.Fatalln(err)
	}

}

func ConnectDB() *sql.DB {
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
return db
}
