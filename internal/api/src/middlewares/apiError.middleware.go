package middlewares

import (
	"errors"

	"github.com/gin-gonic/gin"
)

type ApiError struct {
	Status  int    `json:"status"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (e *ApiError) Error() string {
	return e.Message
}

var (
	ErrorNotFound       = &ApiError{Status: 404, Code: "NOT_FOUND", Message: "Data not found"}
	ErrorUnauthorized   = &ApiError{Status: 401, Code: "UNAUTHORIZED", Message: "Authentication required"}
	ErrorBadRequest     = &ApiError{Status: 400, Code: "BAD_REQUEST", Message: "Invalid request"}
	ErrorInternalServer = &ApiError{Status: 500, Code: "INTERNAL_SERVER_ERROR", Message: "Internal server error"}
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) == 0 {
			return
		}

		err := c.Errors.Last().Err
		var apiErr *ApiError
		if errors.As(err, &apiErr) {
			{
				c.JSON(apiErr.Status, gin.H{
					"code":    apiErr.Code,
					"message": apiErr.Message,
				})
			}
		} else {
			{
				c.JSON(ErrorInternalServer.Status, gin.H{
					"code":    ErrorInternalServer.Code,
					"message": ErrorInternalServer.Message,
				})
			}
		}
	}
}
