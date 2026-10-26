package ginlogger

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/natefinch/lumberjack.v2"
)

func ConfigGinLogger() {
	gin.DisableConsoleColor()

	gin.DefaultWriter = &lumberjack.Logger{
		Filename:   "../../gin.log",
		MaxSize:    100,
		MaxBackups: 3,
		MaxAge:     30,
	}
}
