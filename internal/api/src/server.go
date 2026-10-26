package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/incodi404/hookr.git/internal/api/src/routes"
	ginlogger "github.com/incodi404/hookr.git/pkg/gin-logger"
	logruslogger "github.com/incodi404/hookr.git/pkg/logrus-logger"
	"github.com/incodi404/hookr.git/pkg/models"
)

func StartServer(config models.ServerConfig) (*gin.Engine, error) {
	// gin logger
	ginlogger.ConfigGinLogger()
	logruslogger.ConfigLogrus()

	router := gin.Default()

	// include routes
	routes.ConfigRoutes(router)

	s := &http.Server{
		Addr:           config.Addr,
		Handler:        router,
		ReadTimeout:    time.Duration(config.ReadTimeout),
		WriteTimeout:   time.Duration(config.WriteTimeout),
		MaxHeaderBytes: config.MaxHeaderBytes,
	}

	fmt.Println("[API] Server is listening on: 7000")

	// start listening
	err := s.ListenAndServe()
	if err != nil {
		logruslogger.Writer.Error("[API] Error starting Gin server: " + err.Error())
		return nil, fmt.Errorf("[API] Error starting Gin server: %w", err)
	}

	return router, nil
}
