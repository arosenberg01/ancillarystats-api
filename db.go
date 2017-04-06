package main

import (
	_ "database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func NewDB(dataSource string) (*sqlx.DB) {
	db := sqlx.MustConnect("mysql", dataSource)

	return db
}
