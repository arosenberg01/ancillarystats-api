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

type Game struct {
	PlayerId string `json:"player_id" db:"player_id"`
	Date string `json:"date" db:"date"`
	Opponent string `json:"opponent" db:"opp"`
	Away int `json:"away" db:"away"`
	Score string `json:"score" db:"score"`
	SecondsPlayed int `json:"seconds_played" db:"sec_played"`
	FieldGoalsMade int `json:"field_goals_made" db:"fgm"`
	FieldGoalsAttempted int `json:"field_goals_attempted" db:"fga"`
	FieldGoalPercentage float32 `json:"field_goal_percentage" db:"fg_pct"`
	ThreePointersMade int `json:"three_pointers_made" db:"three_pm"`
	ThreePointersAttempted int `json:"three_pointers_attempted" db:"three_pa"`
	ThreePointPercentage float32 `json:"three_pointer_percentage" db:"three_pct"`
	FreeThrowsMade int `json:"free_throws_made" db:"ftm"`
	FreeThrowsAttempted int `json:"free_throws_attempted" db:"fta"`
	FreeThrowPercentage float32 `json:"free_throw_percentage" db:"ft_pct"`
	OffensiveRebounds int `json:"offensive_rebounds" db:"off_reb"`
	DefensiveRebounds int `json:"defensive_rebounds" db:"def_reb"`
	TotalRebounds int `json:"total_rebounds" db:"total_reb"`
	Assists int `json:"assists" db:"ast"`
	Turnovers int `json:"turnovers" db:"to"`
	Steals int `json:"steals" db:"stl"`
	Blocks int `json:"blocks" db:"blk"`
	PersonalFouls int `json:"personal_fouls" db:"pf"`
	Points int `json:"points" db:"pts"`
}

type CategoryLeaders struct {
	Category string `json:"category"`
	Leaders []CategoryLeader `json:"leaders"`
}

type CategoryLeader struct {
	Id string `json:"id" db:"id"`
	CatAvg string `json:"value" db:"cat_avg"`
}

type Team struct {
	Id string `json:"id" db:"id"`
}

func (db *DB) NbaPlayer(player_id string) (Player, error) {
	player := Player{}
	err := db.Get(&player, "SELECT id, name, number, team, pos, height, weight FROM nba_player WHERE id=?;", player_id)

	if err != nil {
		return player, err
	}

	return player, nil
}

func (db *DB) NbaCategoryLeaders(category string) (CategoryLeaders, error) {
	categoryLeaders := CategoryLeaders{category, []CategoryLeader{}}
	err := db.Select(&categoryLeaders.Leaders, "SELECT player_id AS id, AVG(" + category + ") AS cat_avg FROM nba_game GROUP BY player_id ORDER BY cat_avg DESC LIMIT 10")

	if err != nil {
		return categoryLeaders, err
	}

	return categoryLeaders, nil
}

func (db *DB) NbaTeams() ([]Team, error) {
	teams := []Team{}
	err := db.Select(&teams, "SELECT * FROM nba_team")

	if err != nil {
		return nil, err
	}

	return teams, nil
}

func (db *DB) NbaRoster(team_id string) ([]Player, error) {
	roster := []Player{}
	err := db.Select(&roster, "SELECT id, name, number, team, pos, height, weight FROM nba_player WHERE team=?;", team_id)

	if err != nil {
		return nil, err
	}

	return roster, nil
}

func (db *DB) NbaGames(player_id string) ([]Game, error) {
	games := []Game{}
	err := db.Select(&games, "SELECT player_id, date, opp, away, COALESCE(score, '') as score, sec_played, fgm, fga, fg_pct, three_pm, three_pa, three_pct, ftm, fta, ft_pct, off_reb, def_reb, total_reb, ast, `to` FROM nba_game WHERE player_id=?;", player_id)

	if err != nil {
		return nil, err
	}

	return games, nil
}
