package config

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestLoadAppConfig(t *testing.T) {
	LoadAppConfig()
	assert.Equal(t, "8080", AppConfig.Port)
	assert.Equal(t, "root@tcp(127.0.0.1:3306)/crud_demo?parseTime=true", AppConfig.ConnectionString)
}
