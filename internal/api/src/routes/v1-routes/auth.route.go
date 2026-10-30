package v1routes

import (
	"github.com/gin-gonic/gin"
	"github.com/incodi404/hookr.git/internal/api/src/controllers"
)

func AuthRoutes(rg *gin.RouterGroup) {
	rg.GET("/signup/otp/generate", controllers.UserGenerateSignupOtp)
	rg.GET("/signup/otp/verify", controllers.UserVerifyOtp)
	rg.GET("/signin", controllers.UserSignin)
}
