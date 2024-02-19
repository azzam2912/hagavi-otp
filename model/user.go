package model

import (
    "time"
)

type DateTime struct {
    time.Time
}

type User struct {
    ID              int			`json:"id"`
    CreatedAt       DateTime    `json:"createdAt"`
    UpdatedAt       DateTime    `json:"updatedAt"`
    Phone           string      `json:"phone"`
    Password        string      `json:"password"`
    Otp             string      `json:"otp"`
    IsOTPVerified   string      `json:"isOTPVerified"`
}