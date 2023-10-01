package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func loadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
}

type jwtEnv struct {
	AccessSecret  string
	RefreshSecret string
}

type configMap struct {
	Jwt *jwtEnv
}

func Env() *configMap {
	if os.Getenv("GO_ENV") != "production" {
		loadEnv()
	}
	return &configMap{
		&jwtEnv{
			AccessSecret:  os.Getenv("ACCESS_SECRET"),
			RefreshSecret: os.Getenv("REFRESH_SECRET"),
		},
	}
}
