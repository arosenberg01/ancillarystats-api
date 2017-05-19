package main

import (
	"github.com/jmoiron/sqlx"
)

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
	Leaders []CategoryLeader
}

type CategoryLeader struct {
	Id string `json:"id" db:"id"`
	CatAvg string `json:"value" db:"cat_avg"`
}

type Team struct {
	Id string `json:"id" db:"id"`
}

func Leaders(db *sqlx.DB, category string) (CategoryLeaders, error) {
	categoryLeaders := CategoryLeaders{category, []CategoryLeader{}}
	err := db.Select(&categoryLeaders.Leaders, "SELECT player_id AS id, AVG(" + category + ") AS cat_avg FROM nba_game GROUP BY player_id ORDER BY cat_avg DESC LIMIT 10")

	if err != nil {
		return categoryLeaders, err
	}

	return categoryLeaders, nil
}

func PlayerById(db *sqlx.DB, player_id string) (Player, error) {
	player := Player{}
	err := db.Get(&player, "SELECT id, name, number, team, pos, height, weight FROM nba_player WHERE id=?;", player_id)

	if err != nil {
		return player, err
	}

	return player, nil
}

func Teams(db *sqlx.DB) ([]Team, error) {
	teams := []Team{}
	err := db.Select(&teams, "SELECT * FROM nba_team")

	if err != nil {
		return nil, err
	}

	return teams, nil
}

func Roster(db *sqlx.DB, team_id string) ([]Player, error) {
	roster := []Player{}
	err := db.Select(&roster, "SELECT id, name, number, team, pos, height, weight FROM nba_player WHERE team=?;", team_id)

	if err != nil {
		return nil, err
	}

	return roster, nil
}