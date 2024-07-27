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

func main() {
	r := mux.NewRouter()

	// define HTTP routes using the router
}