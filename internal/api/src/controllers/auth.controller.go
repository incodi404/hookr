package controllers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/incodi404/hookr.git/internal/api/src/middlewares"
	"github.com/incodi404/hookr.git/pkg/models/dto"
	"github.com/incodi404/hookr.git/pkg/utils"
)

func UserGenerateSignupOtp(c *gin.Context) {
	// get data
	var reqBody dto.GenerateOTPReq

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.Error(middlewares.ErrorBadRequest)
		return
	}

	// validating
	email := strings.ToLower(strings.TrimSpace(reqBody.Email))
	if email == "" {
		c.Error(&middlewares.ApiError{
			Status:  http.StatusBadRequest,
			Code:    "BAD_REQUEST",
			Message: "Email is required",
		})
	}

	fmt.Println("Email: ", reqBody.Email)

	utils.OK(c, nil, "OTP has been generated", nil)
}

func UserVerifyOtp(c *gin.Context) {
	// get data
	var reqBody dto.VerifyOtp

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.Error(middlewares.ErrorBadRequest)
		return
	}

	fmt.Println("OTP: ", reqBody.Otp)

	utils.OK(c, nil, "OTP has been verified", nil)
}

func UserSignin(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "API is working inb controller"})
}
