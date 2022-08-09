package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetHttpConfig(t *testing.T) {
	assert.Equal(t, getHttpConfig(), &HttpConfig{
		Port: os.Getenv("HTTP_PORT"),
	})
}
