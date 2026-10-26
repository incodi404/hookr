package main

import (
	"fmt"
	"time"

	api "github.com/incodi404/hookr.git/internal/api/src"
	logruslogger "github.com/incodi404/hookr.git/pkg/logrus-logger"
	"github.com/incodi404/hookr.git/pkg/models"
)

func main() {
	// connecting db
	// starting server
	_, err := api.StartServer(models.ServerConfig{
		Addr:           ":7000",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	})

	if err != nil {
		logruslogger.Writer.Error(err)
		fmt.Println(err)
		return
	}
}
