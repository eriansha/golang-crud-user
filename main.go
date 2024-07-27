package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

const (
	dbDriver = "mysql"
	dbUser   = "root"
	dbPass   = "root"
	dbName   = "gocrud_app"
)

type User struct {
	Id    int
	Name  string
	Email string
}

func createUser(db *sql.DB, name, email string) error {
	query := "INSERT INTO users(name, email) values (?, ?)"

	_, err := db.Exec(query, name, email)
	if err != nil {
		return err
	}

	return nil
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	var user User
	json.NewDecoder(r.Body).Decode(&user)

	CreateUser(db, user.Name, user.Email)
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, "User created successfully")
}

func main() {
	r := mux.NewRouter()

	// define HTTP routes using the router

	log.Println("Server listening on :8090")
	log.Fatal(http.ListenAndServe(":8090", r))
}
