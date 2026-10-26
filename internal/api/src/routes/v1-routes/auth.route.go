package v1routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserSignin(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "API is working"})
}

func AuthRoutes(rg *gin.RouterGroup) {
	rg.GET("/signin", UserSignin)
}
