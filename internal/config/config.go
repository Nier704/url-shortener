package config

import (
	"os"

	"github.com/joho/godotenv"
)

type PostgresConfig struct {
	Username string
	Password string
	Host     string
	Port     string
	Name     string
}

func NewPostgresConfig() *PostgresConfig {
	godotenv.Load()

	user := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")

	return &PostgresConfig{
		Username: user,
		Password: password,
		Host:     host,
		Port:     port,
		Name:     name,
	}
}
