package main

import (
	"fmt"
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"os"
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


func PlayerHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var name, number, team, pos string
	var height, weight int


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