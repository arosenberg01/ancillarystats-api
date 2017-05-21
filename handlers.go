package main

import (
	"net/http"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	_ "database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var leaderCategories = map[string]string {
	"points": "pts",
	"rebounds": "total_reb",
	"assists": "ast",
}

func (env *Env) PlayerHandler(w http.ResponseWriter, r *http.Request) (int, error) {
	vars := mux.Vars(r)
	player, err := env.db.NbaPlayer(vars["player_id"])

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

func (env *Env) LeadersHandler(w http.ResponseWriter, r *http.Request) (int, error) {
	vars := mux.Vars(r)
	category, ok := leaderCategories[vars["category"]]


	if ok {
		leaders, err := env.db.NbaCategoryLeaders(category)

		if err != nil {
			return http.StatusInternalServerError, err
		}

		jsonResponse, err := json.Marshal(leaders)

		if err != nil {
			return http.StatusInternalServerError, err
		}

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonResponse)
		return http.StatusOK, nil
	} else {

		return http.StatusNotFound, errors.New("leaders category not available")
	}
}

func (env *Env) TeamsHandler(w http.ResponseWriter, r *http.Request) (int, error) {
	teams, err := env.db.NbaTeams()

	if err != nil {
		return http.StatusInternalServerError, err
	}

	jsonResponse, err := json.Marshal(teams)

	if err != nil {
		return http.StatusInternalServerError, err
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
	return http.StatusOK, nil
}

func (env *Env) RosterHandler(w http.ResponseWriter, r *http.Request) (int, error) {
	vars := mux.Vars(r)
	roster, err := env.db.NbaRoster(vars["team_id"])

	if err != nil {
		return http.StatusNotFound, err
	}

	jsonResponse, err := json.Marshal(roster)

	if err != nil {
		return http.StatusInternalServerError, err
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
	return http.StatusOK, nil
}

func (env *Env) GamesHandler(w http.ResponseWriter, r *http.Request) (int, error) {
	vars := mux.Vars(r)
	games, err := env.db.NbaGames(vars["player_id"])

	if err != nil {
		return http.StatusNotFound, err
	}

	jsonResponse, err := json.Marshal(games)

	if err != nil {
		return http.StatusInternalServerError, err
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
	return 200, nil
}