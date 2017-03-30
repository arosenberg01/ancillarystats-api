package main

type CategoryLeaders struct {
	Category string `json:"category"`
	Leaders []CategoryLeader
}

type CategoryLeader struct {
	Id string `json:"id" db:"id"`
	CatAvg string `json:"value" db:"cat_avg"`
}

func (db *DB) Leaders(category string) (CategoryLeaders, error) {
	categoryLeaders := CategoryLeaders{category, []CategoryLeader{}}
	err := db.Select(&categoryLeaders.Leaders, "SELECT player_id AS id, AVG(" + category + ") AS cat_avg FROM nba_game GROUP BY player_id ORDER BY cat_avg DESC LIMIT 10")

	if err != nil {
		return nil, err
	}

	return categoryLeaders, nil
}