package model


type User struct {
    ID    int				`json:"id" bson:"_id"`
    Name  string            `json:"name"`
    Phone string            `json:"phone"`
    Otp   string            `json:"otp,omitempty"`
}