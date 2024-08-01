package api

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/neoSnakex34/sqli-demo/database"
)

type DBConnection struct {
	db *sql.DB
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func NewDBConnection(db *sql.DB) *DBConnection {
	return &DBConnection{db: db}
}

// barebone cors allowing
func allowCORS(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS") // delete is allowed, we will use it s
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
}

func (d *DBConnection) LoginHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("login handler called")
	allowCORS(&w)

	log.Println("CORS allowed")

	// PLEASE NOTICE
	// i will check method without middlewares for simplicity
	// remember that this is not a real use server but a barebone demo
	// for demonstrating bad code vulnerabilities for sql injection threats

	// barebone preflight check
	if r.Method == http.MethodOptions {
		log.Println("preflight check entered")
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodPost {
		log.Println("method not allowed")
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// in real use password would be encrypted and decripted
	// and authorization would be handled differently
	var credentials Credentials
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		log.Println("error decoding credentials")
		// or internal server err?
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	logged, user, err := database.LogIn(d.db, credentials.Username, credentials.Password)
	if err != nil {
		log.Println("error logging in")
		// SECURITY RISK sql error status will be passed to frontend
		errmsg := "error logging in: " + err.Error()
		http.Error(w, errmsg, http.StatusInternalServerError)
		return
	}
	if !logged {
		log.Println("login failed")
		http.Error(w, "login failed", http.StatusForbidden)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte(user))
	if err != nil {
		log.Println("error writing response")
	}
	log.Println("login successful")

}
