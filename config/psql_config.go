package config

import "os"

type PsqlConfig struct {
	Username string
	Password string
	Database string
	Host     string
	Port     string
}

func getPsqlConfig() *PsqlConfig {
	return &PsqlConfig{
		Username: os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		Host:     os.Getenv("POSTGRES_HOST"),
		Database: os.Getenv("POSTGRES_DATABASE"),
		Port:     os.Getenv("POSTGRES_PORT"),
	}
}
