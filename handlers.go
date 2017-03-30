package main

import (
	"os"
	"fmt"
	"net/http"
	_ "database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/go-sql-driver/mysql"
)

var db *sqlx.DB
var err error

type Player struct {
	Id string `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
	Number string `json:"number" db:"number"`
	Team string `json:"team" db:"team"`
	Position string `json:"position" db:"pos"`
	Height int `json:"height" db:"height"`
	Weight int `json:"weight" db:"weight"`
}

type CategoryLeaders struct {
	Category string `json:"category"`
	Leaders []Leader
}

type Leader struct {
	Id string `json:"id" db:"id"`
	CatAvg string `json:"value" db:"cat_avg"`
}

var leaderCategories = map[string]string {
	"points": "pts",
	"rebounds": "total_reb",
	"assists": "ast",
}

func PlayerHandler(w http.ResponseWriter, r *http.Request) (int, error) {
	vars := mux.Vars(r)
	datasource := os.Getenv("USER") + ":@/" + os.Getenv("DB_NAME")
	db = sqlx.MustConnect("mysql", datasource)

	player := Player{}

	err = db.Get(&player, "SELECT id, name, number, team, pos, height, weight FROM nba_player WHERE id=?;", vars["player_id"])

	if err != nil {
		return http.StatusInternalServerError, err
	}

	jsonResponse, err := json.Marshal(player)

	if err != nil {
		return http.StatusInternalServerError, err
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)

	return http.StatusOK, nil
}

func LeadersHandler(w http.ResponseWriter, r *http.Request) (int, error) {
	vars := mux.Vars(r)
	datasource := os.Getenv("USER") + ":@/" + os.Getenv("DB_NAME")
	db = sqlx.MustConnect("mysql", datasource)

	category, ok := leaderCategories[vars["category"]]

	if ok {

		if err != nil {
			return http.StatusInternalServerError, err
		}

		leaders := []Leader{}
		err = db.Select(&leaders, "SELECT player_id AS id, AVG(" + category + ") AS cat_avg FROM nba_game GROUP BY player_id ORDER BY cat_avg DESC LIMIT 10")

		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(leaders)

		jsonResponse, err := json.Marshal(leaders)

		if err != nil {
			return http.StatusInternalServerError, err
		}

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonResponse)
		return http.StatusOK, nil
	} else {
		return http.StatusInternalServerError, err
	}
}