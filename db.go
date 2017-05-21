package main

import (
	_ "database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"database/sql"
)

type Datastore interface {
	NbaPlayer(string) (Player, error)
	NbaCategoryLeaders (CategoryLeaders, error)
	NbaTeams ([]Team, error)
	NbaRoster([]Player, error)
	NbaGames([]Game, error)
}

type DB struct {
	*sqlx.DB
}

//func NewDB(dataSource string) (*sqlx.DB) {
func NewDB(dataSource string) (*DB) {

	db := sqlx.MustConnect("mysql", dataSource)

	//return db
	return &DB{db}
}

