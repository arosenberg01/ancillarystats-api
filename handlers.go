package main

import (
	"fmt"
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"os"
	"log"
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

type CategoryLeaders struct {
	Category string `json:"category"`
	Leaders []Leader
}

type Leader struct {
	Name string `json:"name"`
	Value string `json:"value"`
}

var leaderCategories = map[string]string {
	"points": "pts",
	"rebounds": "total_reb",
	"assists": "ast",
}


type appHandler func(http.ResponseWriter, *http.Request) (int, error)

func (fn appHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if status, err := fn(w, r); err != nil {
		switch status {
			case http.StatusNotFound:
				http.NotFound(w, r)
			case http.StatusInternalServerError:
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			default:
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
	}
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

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func LeadersHandler(w http.ResponseWriter, r *http.Request) {
	var (
		player_id string
		cat_avg string
	)
	vars := mux.Vars(r)
	datasource := os.Getenv("USER") + ":@/" + os.Getenv("DB_NAME")

	db, err = sql.Open("mysql", datasource)

	if err != nil {
		log.Fatal(err)
	}

	category, ok := leaderCategories[vars["category"]]
	leaders := []Leader{}

	if ok {

		query := "SELECT player_id, AVG(" + category + ") AS cat_avg FROM nba_game GROUP BY player_id ORDER BY cat_avg DESC LIMIT 10"
		rows, err := db.Query(query)

		defer rows.Close()

		for rows.Next() {
			err := rows.Scan(&player_id, &cat_avg)

			if err != nil {
				log.Fatal(err)
			}

			leader := Leader{player_id, cat_avg}
			leaders = append(leaders, leader)
		}

		if err = rows.Err(); err != nil {
			log.Fatal(err)
		}
	}

	jsonResponse, err := json.Marshal(leaders)

	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}