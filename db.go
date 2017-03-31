package main

import (
	_ "database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	_ "database/sql"
)

//type Datastore interface {
//	Player() (Player, error)
//	Leaders() ([]CategoryLeaders, error)
//}
//
//type DB struct {
//	*sqlx.DB
//}

func NewDB(dataSource string) (*sqlx.DB) {
	db, _ := sqlx.Open("mysql", dataSource)

	return db
}
