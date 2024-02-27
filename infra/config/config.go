package config

import (
	"os"
	"strconv"
)

type Config struct {
	Port    string
	Timeout int64
}

func GetEnvOr(key string, or string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return or
}

func NewConfig() *Config {
	num, _ := strconv.ParseInt(GetEnvOr("TIMEOUT", "500000"), 10, 64)
	return &Config{
		Port:    GetEnvOr("PORT", "8080"),
		Timeout: num,
	}
}
