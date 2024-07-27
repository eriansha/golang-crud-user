package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

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

func DeleteUser(db *sql.DB, id int) error {
	query := "DELETE FROM users WHERE id = ?"
	_, err := db.Exec(query, id)

	if err != nil {
		return err
	}

	return nil
}

func updateUser(db *sql.DB, id int, name, email string) error {
	query := "UPDATE users SET name = ?, email = ? where id = ?"

	_, err := db.Exec(query, name, email, id)
	if err != nil {
		return err
	}

	return nil
}

func createUser(db *sql.DB, name, email string) error {
	query := "INSERT INTO users(name, email) values (?, ?)"

	_, err := db.Exec(query, name, email)
	if err != nil {
		return err
	}

	return nil
}

func getUser(db *sql.DB, id int) (*User, error) {
	query := "SELECT * FROM users WHERE id = ?"
	row := db.QueryRow(query, id)

	user := &User{}
	err := row.Scan(&user.Id, &user.Name, &user.Email)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	// Parse JSON data from the request body
	var user User
	json.NewDecoder(r.Body).Decode(&user)

	createUser(db, user.Name, user.Email)
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "User created successfully")
}

func getUserHandler(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	// Get the 'Id' parameter from the URL
	vars := mux.Vars(r)
	idStr := vars["id"]

	// Convert 'id' to an integer
	userId, err := strconv.Atoi(idStr)

	// Call the getUser function to fetch the user data from the database
	user, err := getUser(db, userId)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Convert the user object to JSON and send it in the response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func updateUserHandler(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	// Get the 'id' parameter from the URL
	vars := mux.Vars(r)
	idStr := vars["id"]

	// Convert 'id' to an integer
	userId, errParse := strconv.Atoi(idStr)
	if errParse != nil {
		http.Error(w, "Invalid 'id' parameter", http.StatusBadRequest)
		return
	}

	// Call the getUser function to fetch the user data from the database
	_, errGetUser := getUser(db, userId)
	if errGetUser != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	var user User
	err = json.NewDecoder(r.Body).Decode(&user)

	// Call the getUser function to fetch the user data from the database
	errUpdate := updateUser(db, userId, user.Name, user.Email)
	if errUpdate != nil {
		http.Error(w, "Failed to update User", http.StatusBadRequest)
		return
	}

	fmt.Fprintln(w, "User update successfully")
}

func deleteUserHandler(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	// Get the 'id' parameter from the URL
	vars := mux.Vars(r)
	idStr := vars["id"]

	// Convert 'id' to an integer
	userID, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid 'id' parameter", http.StatusBadRequest)
		return
	}

	user := DeleteUser(db, userID)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	fmt.Fprintln(w, "User deleted successfully")
}

func main() {
	r := mux.NewRouter()

	// define HTTP routes using the router
	r.HandleFunc("/user", createUserHandler).Methods("POST")
	r.HandleFunc("/user/{id}", getUserHandler).Methods("GET")
	r.HandleFunc("/user/{id}", updateUserHandler).Methods("PUT")
	r.HandleFunc("/user/{id}", deleteUserHandler).Methods("DELETE")

	log.Println("Server listening on :8090")
	log.Fatal(http.ListenAndServe(":8090", r))
}
