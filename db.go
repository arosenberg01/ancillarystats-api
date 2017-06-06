package main

import (
	_ "database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Datastore interface {
	NbaPlayer(string) (NbaPlayer, error)
	NbaCategoryLeaders(string) (NbaCategoryLeaders, error)
	NbaTeams() (NbaTeams, error)
	NbaRoster(string) (NbaRoster, error)
	NbaGames(string) (NbaGames, error)
}

type DB struct {
	*sqlx.DB
}

func NewDB(dataSource string) (*DB) {
	db := sqlx.MustConnect("mysql", dataSource)

	return &DB{db}
}

