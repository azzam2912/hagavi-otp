package schema

type RegisterBody struct {
    Phone       string `json:"phone"`
    Password    string `json:"password"`
}

type VerifyOTP struct {
    Phone string `json:"phone"`
    Otp   string `json:"otp"`
}

type LoginSchema struct {
	Phone       string `json:"phone"`
    Password    string `json:"password"`
}