package services

import (
	"context"

	"github.com/incodi404/hookr.git/internal/api/src/repositories"
)

type AuthService struct {
	AuthRepo repositories.AuthRepository
	OTPRepo  repositories.OTPRepository
}

func NewAuthService(
	authRepo repositories.AuthRepository,
	otpRepo repositories.OTPRepository,
) *AuthService {
	return &AuthService{
		AuthRepo: authRepo,
		OTPRepo:  otpRepo,
	}
}

func (a *AuthService) GenerateOtp(
	ctx context.Context,
	email_id string,
) {
}
