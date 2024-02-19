package database

import (
   "database/sql" // add this
   "hagavi-otp/config"
    _ "github.com/lib/pq"
)

func Connect() (*sql.DB) {
	connectString := config.Config("SQL_CONNECT")
	db, err := sql.Open("postgres", connectString)
	if err != nil {
		panic(err)
	}
	return db
}
