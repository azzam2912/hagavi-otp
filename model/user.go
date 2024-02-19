package model

import (
    "time"
)

type DateTime struct {
    time.Time
}

type User struct {
    ID              int			`json:"id"`
    CreatedAt       time.Time   `json:"createdAt"`
    UpdatedAt       time.Time   `json:"updatedAt"`
    Phone           string      `json:"phone"`
    Password        string      `json:"password"`
    Otp             string      `json:"otp"`
    IsOTPVerified   bool        `json:"isOTPVerified"`
}