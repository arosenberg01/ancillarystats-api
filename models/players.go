package main

type Player struct {
	Id string `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
	Number string `json:"number" db:"number"`
	Team string `json:"team" db:"team"`
	Position string `json:"position" db:"pos"`
	Height int `json:"height" db:"height"`
	Weight int `json:"weight" db:"weight"`
}

func (db *DB) Player(player_id string) (Player, error) {
	player := Player{}
	err := db.Get(&player, "SELECT id, name, number, team, pos, height, weight FROM nba_player WHERE id=?;", player_id)

	if err != nil {
		return nil, err
	}

	return player, nil
}