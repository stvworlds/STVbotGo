package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
)

type Env struct {
	Token               string   `yaml:"token"`
	Environment         string   `yaml:"environment"`
	StartupChannelID    string   `yaml:"startupChannelID"`
	NewsChannels        []string `yaml:"newsChannels"`
	UserNewsChannel     string   `yaml:"userNewsChannel"`
	NewsGuildID         string   `yaml:"newsGuildID"`
	NewsRoleID          string   `yaml:"newsRoleID"`
	BlogAPIKey          string   `yaml:"blogAPIKey"`
	ScamChannelID       string   `yaml:"scamChannelID"`
	AppLogChannelID     string   `yaml:"appLogChannelID"`
	ActionsLogChannelID string   `yaml:"actionsLogChannelID"`
}

func Load() (Env, error) {
	if err := godotenv.Load(); err != nil && !os.IsNotExist(err) {
		return Env{}, fmt.Errorf("load .env: %w", err)
	}

	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		configName := os.Getenv("APP_CONFIG")
		if configName == "" {
			configName = "dev"
		}

		// Try /configs first (Docker), then local ./configs
		configPath = filepath.Join("/configs", configName+".yml")
		if _, err := os.Stat(configPath); os.IsNotExist(err) {
			configPath = filepath.Join("configs", configName+".yml")
		}
	}

	fmt.Printf("Loading config from: %s\n", configPath)

	env, err := loadYAML(configPath)
	if err != nil {
		return Env{}, err
	}

	applyEnvOverrides(&env)

	if env.Token == "" {
		return Env{}, fmt.Errorf("TOKEN is empty")
	}

	return env, nil
}

func loadYAML(configPath string) (Env, error) {
	data, err := os.ReadFile(configPath)
	if err != nil {
		return Env{}, fmt.Errorf("read config file %s: %w", configPath, err)
	}

	var env Env
	if err = yaml.Unmarshal(data, &env); err != nil {
		return Env{}, fmt.Errorf("parse config file %s: %w", configPath, err)
	}

	return env, nil
}

func applyEnvOverrides(env *Env) {
	if token := os.Getenv("TOKEN"); token != "" {
		env.Token = token
	}

	if environment := os.Getenv("APP_ENV"); environment != "" {
		env.Environment = environment
	}

	if blogAPIKey := os.Getenv("BLOG_API_KEY"); blogAPIKey != "" {
		env.BlogAPIKey = blogAPIKey
	}
}
