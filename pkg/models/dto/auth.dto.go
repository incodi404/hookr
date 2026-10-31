package dto

type GenerateOTPReq struct {
	Email string `json:"email" binding:"required"`
}

type VerifyOtpReq struct {
	Email string `json:"email" binding:"required"`
	Otp   string `json:"otp" binding:"required"`
}
