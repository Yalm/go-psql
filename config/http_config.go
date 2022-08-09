package config

import "os"

type HttpConfig struct {
	Port string
}

func getHttpConfig() *HttpConfig {
	return &HttpConfig{
		Port: os.Getenv("HTTP_PORT"),
	}
}
