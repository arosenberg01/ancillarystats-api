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
	datasource := fmt.Sprintf("%s:%s@tcp(%s)/%s", os.Getenv("DB_USER"), os.Getenv("DB_PW"), os.Getenv("DB_INSTANCE"), os.Getenv("DB_NAME"))

	fmt.Printf(datasource)

	db := NewDB(datasource)
	env := &Env{db}
	router := NewRouter(env)

	log.Fatal(http.ListenAndServe(":3000", router))
}