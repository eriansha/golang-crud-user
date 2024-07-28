package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	models "golangcrud/models/user"
	"golangcrud/repositories"
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

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	// Parse JSON data from the request body
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	repositories.CreateUser(db, user.Name, user.Email)
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
	user, err := repositories.GetUser(db, userId)
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
	_, errGetUser := repositories.GetUser(db, userId)
	if errGetUser != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	var user models.User
	err = json.NewDecoder(r.Body).Decode(&user)

	// Call the getUser function to fetch the user data from the database
	errUpdate := repositories.UpdateUser(db, userId, user.Name, user.Email)
	if errUpdate != nil {
		http.Error(w, "Failed to update models.User", http.StatusBadRequest)
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
	userId, errParse := strconv.Atoi(idStr)
	if errParse != nil {
		http.Error(w, "Invalid 'id' parameter", http.StatusBadRequest)
		return
	}

	// Call the getUser function to fetch the user data from the database
	_, errGetUser := repositories.GetUser(db, userId)
	if errGetUser != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	errDelete := repositories.DeleteUser(db, userId)
	if errDelete != nil {
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
