package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	Token string
}

func Load() (Env, error) {
	if err := godotenv.Load(); err != nil {
		return Env{}, fmt.Errorf("load .env: %w", err)
	}

	token := os.Getenv("TOKEN")
	if token == "" {
		return Env{}, fmt.Errorf("TOKEN is empty")
	}

	return Env{Token: token}, nil
}
