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
	"field_goals_made": "fgm",
	"field_goals_attempted": "fga",
	"field_goal_percentage": "fg_pct",
	"three_pointers_made": "three_pm",
	"three_pointers_attempted": "three_pa",
	"three_pointer_percentage": "three_pct",
	"free_throws_made": "ftm",
	"free_throws_attempted": "fta",
	"free_throw_percentage": "ft_pct",
	"offensive_rebounds": "off_reb",
	"defensive_rebounds": "def_reb",
	"total_rebounds": "total_reb",
	"assists": "ast",
	"turnovers": "to",
	"steals": "stl",
	"blocks": "blk",
	"personal_fouls": "pf",
	"points": "pts",
}

func SendJsonResponse(w http.ResponseWriter, jsonResponse []byte, err error) (int, error) {
	if err != nil {
		return http.StatusInternalServerError, err
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)

	return http.StatusOK, nil
}

func (env *Env) NbaPlayerHandler(w http.ResponseWriter, r *http.Request) (int, error) {
	vars := mux.Vars(r)
	player, err := env.db.NbaPlayer(vars["player_id"])

	if err != nil {
		return http.StatusInternalServerError, err
	}

	jsonResponse, err := json.Marshal(player)

	return SendJsonResponse(w, jsonResponse, err)
}

func (env *Env) NbaCategoriesHandler(w http.ResponseWriter, r *http.Request) (int, error) {
	availableCategories := StrMapKeys(leaderCategories)
	jsonResponse, err := json.Marshal(availableCategories)

	return SendJsonResponse(w, jsonResponse, err)
}

func (env *Env) NbaLeadersHandler(w http.ResponseWriter, r *http.Request) (int, error) {
	vars := mux.Vars(r)
	category, ok := leaderCategories[vars["category"]]

	if ok {
		leaders, err := env.db.NbaCategoryLeaders(category)

		if err != nil {
			return http.StatusInternalServerError, err
		}

		jsonResponse, err := json.Marshal(leaders)

		return SendJsonResponse(w, jsonResponse, err)
	} else {

		return http.StatusNotFound, errors.New("leaders category not available")
	}
}

func (env *Env) NbaTeamsHandler(w http.ResponseWriter, r *http.Request) (int, error) {
	teams, err := env.db.NbaTeams()

	if err != nil {
		return http.StatusInternalServerError, err
	}

	jsonResponse, err := json.Marshal(teams)

	return SendJsonResponse(w, jsonResponse, err)
}

func (env *Env) NbaRosterHandler(w http.ResponseWriter, r *http.Request) (int, error) {
	vars := mux.Vars(r)
	roster, err := env.db.NbaRoster(vars["team_id"])

	if err != nil {
		return http.StatusNotFound, err
	}

	jsonResponse, err := json.Marshal(roster)

	return SendJsonResponse(w, jsonResponse, err)
}

func (env *Env) NbaGamesHandler(w http.ResponseWriter, r *http.Request) (int, error) {
	vars := mux.Vars(r)
	games, err := env.db.NbaGames(vars["player_id"])

	if err != nil {
		return http.StatusNotFound, err
	}

	jsonResponse, err := json.Marshal(games)

	return SendJsonResponse(w, jsonResponse, err)
}
