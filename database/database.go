package database

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

/* Connect to db file*/
func OpenDB(dbfile string) (*sql.DB, error) { // TODO maybe return err
	db, err := sql.Open("sqlite3", dbfile)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func InitializeDB(db *sql.DB) error {

	userTableStatement := `CREATE TABLE IF NOT EXISTS users (
	id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	username TEXT NOT NULL, 
	password TEXT NOT NULL
	);`

	_, err := db.Exec(userTableStatement)
	if err != nil {
		log.Println("an error occurred while creating the users table: ", err)
		return err
	}

	err = AddUser(db, "piero", "s3cret")
	if err != nil {
		log.Println("an error occurred while adding user piero: ", err)
		return err
	}

	err = AddUser(db, "giorgio", "prova!")
	if err != nil {
		log.Println("an error occurred while adding user giorgio: ", err)
		return err
	}

	err = AddUser(db, "susan", "l33tPasswd!")
	if err != nil {
		log.Println("an error occurred while adding user susan: ", err)
		return err
	}

	err = AddUser(db, "brian", "cann0tfind!")
	if err != nil {
		log.Println("an error occurred while adding user brian: ", err)
		return err
	}

	return nil

}

func AddUser(db *sql.DB, username string, password string) error { // passwords will be plain for now
	// this safely adds user with right written queries
	userInsertStatement := `INSERT INTO users (username, password) VALUES (?, ?);`

	_, err := db.Exec(userInsertStatement, username, password)
	if err != nil {
		log.Println("an error occurred while inserting users in db: ", err)
		return err
	}

	return nil

}

func LogIn(db *sql.DB, username string, password string) (bool, string, error) {
	// SECURITY RISK this code is vulnerable to sqlinjection
	query := fmt.Sprintf(`SELECT id, username FROM users WHERE username = '%s' AND password = '%s';`, username, password)

	log.Println("QUERY: ", query)
	rows, err := db.Query(query)
	if err != nil {
		log.Println("an error occurred while querying the db: ", err)
		return false, "", err
	}

	// SECURITY RISK this will give too much info in frontend
	var user string // this wont be a single user
	for rows.Next() {
		var id int
		var username string
		err = rows.Scan(&id, &username)
		if err != nil {
			log.Println("an error occurred while scanning the rows: ", err)
			return false, "", err
		}
		user += fmt.Sprintf("id: %d, username: %s\n", id, username)
	}

	// SECURITY RISK absolutely poor design, one may decide to use it for debugging reasons
	// but leaving this in the code is really bad (and stupid) since it gives attackers lots of details
	if user == "" {
		return false, "", errors.New("user " + username + " not found or wrong password")
	}
	// if !rows.Next() {
	// 	return false, "", errors.New("user " + username + " not found")
	// }

	defer rows.Close()
	return true, user, nil

}
