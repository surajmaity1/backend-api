package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"os"
)

type Config struct {
	ENV            Environment
	MYSQL_HOST     string
	MYSQL_PORT     string
	MYSQL_USER     string
	MYSQL_PASS     string
	MYSQL_DATABASE string
}

var AppConfig Config

func init() {
	if err := godotenv.Load(); err != nil {
		logrus.Error(fmt.Sprintf("Error loading .env file %s", err))
	} else {
		logrus.Info(fmt.Sprintf("ENV loaded"))
	}

	AppConfig = Config{
		ENV:            Environment(loadEnv("ENV")).Validate(),
		MYSQL_HOST:     loadEnv("MYSQL_HOST"),
		MYSQL_PORT:     loadEnv("MYSQL_PORT"),
		MYSQL_USER:     loadEnv("MYSQL_USER"),
		MYSQL_PASS:     loadEnv("MYSQL_PASS"),
		MYSQL_DATABASE: loadEnv("MYSQL_DATABASE"),
	}
}

func loadEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		logrus.Panic(fmt.Sprintf("Environment variable %s is not found", key))
	}
	return value
}
