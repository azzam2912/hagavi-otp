package util

import (
	"database/sql"
	"fmt"
	"hagavi-otp/model"
	"hagavi-otp/schema"
)

func FindUserByPhoneNumber(phoneNumber string, db *sql.DB) (*model.User, error) {
	var result model.User
	row := db.QueryRow(`SELECT * FROM user_table WHERE phone = $1`, phoneNumber)
	err := row.Scan(&result.ID, &result.CreatedAt, &result.UpdatedAt, &result.Phone, &result.Otp, &result.Password, &result.IsOTPVerified)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	} else {
		return &result, nil
	}
}

func AddUser(user *schema.RegisterBody, db *sql.DB) error {
	_, err := db.Exec(`INSERT INTO user_table (phone, password) VALUES ($1, $2)`,user.Phone, user.Password)
	if err != nil {
		fmt.Println(err)
	} else {
		err = nil
	}
	return err
}
