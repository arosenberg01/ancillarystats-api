package main

import (
	"os"
	"fmt"
	"log"
	"net/http"
	"github.com/jmoiron/sqlx"
)

type Env struct {
	db *sqlx.DB
}

func main() {
	var datasource string

	appEnv := os.Getenv("APP_ENV")

	if appEnv == "loc" {
		datasource = os.Getenv("NBA_DSN")
	} else {
		datasource = fmt.Sprintf("%s:%s@tcp(%s)/%s", os.Getenv("DB_USER"), os.Getenv("DB_PW"), os.Getenv("DB_INSTANCE"), os.Getenv("DB_NAME"))
	}
	
	db := NewDB(datasource)
	env := &Env{db}
	router := NewRouter(env)

	log.Fatal(http.ListenAndServe(":8080", router))
}