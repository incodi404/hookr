package dto

type GenerateOTPReq struct {
	Email string `json:"email" binding:"required"`
}

type VerifyOtp struct {
	Email string `json:"email" binding:"required"`
	Otp   string `json:"otp" binding:"required"`
}
