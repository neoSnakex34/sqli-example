package main

import (
	"database/sql"
	"errors"
	"log"
	"net/http"
	"os"

	"github.com/neoSnakex34/sqli-demo/api"
	"github.com/neoSnakex34/sqli-demo/database"
)

// utility function to check if the db file exists
func checkExist(filepath string) bool {
	_, err := os.Stat(filepath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return false
		}
		// if any other error is returned in retrieving file
		// execution will stop with panic
		// one could decide to log error and rebuild db anyway
		// but since it's kinda crucial that db exists or not
		// i decided to make the program panic
		panic(err)
	}
	return true
}

func main() {

	// check if db file exist
	// if not create it
	var db *sql.DB
	var err error
	if !checkExist("vulnerable.db") {
		log.Println("db file does not exists, creating...")
		// create
		// open connection
		db, err = database.OpenDB("vulnerable.db")
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		// initializing with demo data
		err = database.InitializeDB(db)
		if err != nil {
			log.Fatal(err)
		}

	} else {
		db, err = database.OpenDB("vulnerable.db")
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()
	}

	// // server operations
	dbconnection := api.NewDBConnection(db)
	http.HandleFunc("/login", dbconnection.LoginHandler)

	log.Println("server listening on port 3000")
	log.Fatal(http.ListenAndServe(":3000", nil))

}
