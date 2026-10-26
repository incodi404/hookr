package logruslogger

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

var Writer *logrus.Logger

func ConfigLogrus() {
	logger := logrus.New()

	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.SetOutput(&lumberjack.Logger{
		Filename:   "../../log/app.log",
		MaxSize:    100, // MB
		MaxBackups: 5,   // Number of old files
		MaxAge:     30,  // Days
		Compress:   true,
	})

	Writer = logger
}
