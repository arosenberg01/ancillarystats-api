package main

import (
	_ "database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Datastore interface {

}

type DB struct {
	*sqlx.DB
}

func NewDb(dataSource string) (*DB) {
	db := sqlx.MustConnect("mysql", dataSource)

	return &DB{db}
}
