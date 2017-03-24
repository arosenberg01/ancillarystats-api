package main

import (
	"fmt"
	"os"
	"net/http"
	"database/sql"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
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

func handler(w http.ResponseWriter, r *http.Request) {
	var name, number, team, pos string
	var height, weight int

	player_id := r.URL.Path[1:]
	row := db.QueryRow("SELECT name, number, team, pos, height, weight FROM nba_player WHERE id=?;", player_id)
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
		panic(err.Error())
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		panic(err.Error())
	}
}