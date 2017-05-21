package main

import (
	_ "database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Datastore interface {
	NbaPlayer(string) (Player, error)
	NbaCategoryLeaders(string) (CategoryLeaders, error)
	NbaTeams() ([]Team, error)
	NbaRoster(string) ([]Player, error)
	NbaGames(string) ([]Game, error)
}

type DB struct {
	*sqlx.DB
}

func NewDB(dataSource string) (*DB) {
	db := sqlx.MustConnect("mysql", dataSource)

	return &DB{db}
}

