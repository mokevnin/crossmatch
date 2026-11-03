package config_test

import (
	"strings"
	"testing"

	"github.com/mokevnin/crossmatch/internal/config"
)

func env(values map[string]string) func(string) string {
	return func(key string) string { return values[key] }
}

func TestLoad(t *testing.T) {
	cfg, err := config.Load(env(map[string]string{"BOT_TOKEN": "t", "DATABASE_URL": "postgres://db"}))
	if err != nil {
		t.Fatalf("Load returned an error: %v", err)
	}
	if cfg.BotToken != "t" || cfg.DatabaseURL != "postgres://db" {
		t.Errorf("Load = %+v", cfg)
	}
}

func TestLoadReportsEverythingMissing(t *testing.T) {
	_, err := config.Load(env(nil))
	if err == nil {
		t.Fatal("Load accepted an empty environment")
	}
	for _, name := range []string{"BOT_TOKEN", "DATABASE_URL"} {
		if !strings.Contains(err.Error(), name) {
			t.Errorf("error %q does not name %s", err, name)
		}
	}
}
