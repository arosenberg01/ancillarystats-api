package main

import (
	"fmt"
	"os"
	"log"
	"net/http"
	"database/sql"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var db *sql.DB
var err error

type Player struct {
	Name string `json:"name"`
	Number string `json:"number"`
	Team string `json:"team"`
	Position string `json:"position"`
	Height int `json:"height"`
	Weight int `json:"weight"`
}

/*
type appError struct {
	Error error
	Message string
	Code int
}
*/

func PlayerHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var name, number, team, pos string
	var height, weight int

	row := db.QueryRow("SELECT name, number, team, pos, height, weight FROM nba_player WHERE id=?;", vars["player_id"])
	err = row.Scan(&name, &number, &team, &pos, &height, &weight)

	if err != nil {
		fmt.Println(err)
	}

	player := Player{name, number, team, pos, height, weight}
	jsonResponse, err := json.Marshal(player)

	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func main() {
	datasource := os.Getenv("USER") + ":@/" + os.Getenv("DB_NAME")
	db, err = sql.Open("mysql", datasource)

	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println(err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/player/{player_id}", PlayerHandler)

	log.Fatal(http.ListenAndServe(":8080", r))
}