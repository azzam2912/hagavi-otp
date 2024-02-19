package util

import (
	"database/sql"
	"fmt"
	"hagavi-otp/config"
	"hagavi-otp/model"
	"hagavi-otp/schema"
)


func FindUserByPhoneNumber(phoneNumber string, db *sql.DB) (*model.User, error) {
	var result model.User
	row := db.QueryRow(`SELECT * FROM $1 WHERE phone = $2`, config.Config("SQL_TABLE_NAME"), phoneNumber)
	err := row.Scan(&result)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	} else {
		return &result, nil
	}
}

func AddUser(user *schema.RegisterBody, db *sql.DB) error {
	_, err := db.Exec(`INSERT INTO $1 (phone, password) VALUES ($2, $3)`, config.Config("SQL_TABLE_NAME"), user.Phone, user.Password)
	if err != nil {
		fmt.Println(err)
	} else {
		err = nil
	}
	return err
}

