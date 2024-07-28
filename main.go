package main

import (
	"golangcrud/handlers"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// define HTTP routes using the router
	r.HandleFunc("/user", handlers.CreateUserHandler).Methods("POST")
	r.HandleFunc("/user/{id}", handlers.GetUserHandler).Methods("GET")
	r.HandleFunc("/user/{id}", handlers.UpdateUserHandler).Methods("PUT")
	r.HandleFunc("/user/{id}", handlers.DeleteUserHandler).Methods("DELETE")

	log.Println("Server listening on :8090")
	log.Fatal(http.ListenAndServe(":8090", r))
}
