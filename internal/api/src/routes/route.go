package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/incodi404/hookr.git/internal/api/src/middlewares"
	v1routes "github.com/incodi404/hookr.git/internal/api/src/routes/v1-routes"
)

func ConfigRoutes(router *gin.Engine) {
	// v1
	{
		v1 := router.Group("v1")
		v1.Use(middlewares.ErrorHandler()) // error handler middleware

		// auth
		{
			auth := v1.Group("/auth")
			v1routes.AuthRoutes(auth)
		}
	}
}
