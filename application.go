package main

import (
	"os"
	"fmt"
	"log"
	"net/http"
	"github.com/urfave/negroni"
)

type Env struct {
	db Datastore
}

func main() {
	datasource := fmt.Sprintf("%s:%s@tcp(%s)/%s", os.Getenv("DB_USER"), os.Getenv("DB_PW"),
		os.Getenv("DB_INSTANCE"), os.Getenv("DB_NAME"))
	db := NewDB(datasource)
	env := &Env{db}

	router := NewRouter(env)

	n := negroni.New()
	n.UseHandler(router)


	port := ":5000"
	if envPort := os.Getenv("PORT"); envPort != "" {
		port = ":" + envPort
	}

	//log.Fatal(http.ListenAndServe(port, router))
	log.Fatal(http.ListenAndServe(port, n))

}