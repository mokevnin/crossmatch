// Package config reads the program's settings from the environment.
package config

import (
	"errors"
	"fmt"
)

// Config is everything the program needs from its environment.
type Config struct {
	// BotToken is the Telegram credential of the bot.
	BotToken string
	// DatabaseURL is the PostgreSQL connection string.
	DatabaseURL string
}

// Load reads the settings through getenv, so a test does not touch the process environment.
// It reports every missing setting at once.
func Load(getenv func(string) string) (Config, error) {
	cfg := Config{
		BotToken:    getenv("BOT_TOKEN"),
		DatabaseURL: getenv("DATABASE_URL"),
	}

	var missing []error
	if cfg.BotToken == "" {
		missing = append(missing, fmt.Errorf("BOT_TOKEN is not set"))
	}
	if cfg.DatabaseURL == "" {
		missing = append(missing, fmt.Errorf("DATABASE_URL is not set"))
	}
	if err := errors.Join(missing...); err != nil {
		return Config{}, err
	}
	return cfg, nil
}
