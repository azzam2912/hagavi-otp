package schema

type RegisterBody struct {
    Name  string `json:"name"`
    Phone string `json:"phone"`
}

type VerifyOTP struct {
    Phone string `json:"phone"`
    Otp   string `json:"otp"`
}