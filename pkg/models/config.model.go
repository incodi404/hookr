package models

import "time"

type ServerConfig struct {
	Addr           string
	ReadTimeout    time.Duration
	WriteTimeout   time.Duration
	MaxHeaderBytes int
}

type DBConfig struct {
	Host     string
	Port     int
	DbName   string
	User     string
	Password string
}
