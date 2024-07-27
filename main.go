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
	dbUser = "root"
	dbPass = "root"
	dbName = "gocrud_app"
)

type User struct {
	Id    int
	Name  string
	Email string
}

func main() {
	r := mux.NewRouter()

	// define HTTP routes using the router

	log.Println("Server listening on :8090")
	log.Fatal(http.ListenAndServe(":8090", r))
}