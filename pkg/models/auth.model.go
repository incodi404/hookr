package models

type GenerateOtp struct {
	Email string `json:"email"`
}

type VerifyOtp struct {
	Email string `json:"email"`
	Otp   string `json:"otp"`
}
