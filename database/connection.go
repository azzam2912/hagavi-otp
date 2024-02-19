package database

import (
   "database/sql" // add this
   "fmt"
   "hagavi-otp/config"
)

func Connect() *sql.DB {
	connectString := config.Config("SQL_CONNECT")
	db, err := sql.Open("postgres", connectString)
	if err != nil {
		fmt.Println(err)
	}
	return db
}
