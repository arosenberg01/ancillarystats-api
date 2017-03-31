package main

import (
	"os"
	"log"
	"net/http"
	"github.com/jmoiron/sqlx"
)

type Env struct {
	db *sqlx.DB
}

func main() {
	db := NewDB(os.Getenv("NBA_DSN"))
	env := &Env{db}
	router := NewRouter(env)

	log.Fatal(http.ListenAndServe(":8080", router))
}