package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var (
	DB *sql.DB
)

func Initialize(dbDriver string, connString string) {
	var err error

	DB, err = sql.Open(dbDriver, connString)

	if err != nil {
		panic(fmt.Sprintf("could not connect with the database! Error: %s", err.Error()))
	}
}

func Close() {
	DB.Close()
}
