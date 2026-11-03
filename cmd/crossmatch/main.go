// Command crossmatch is the CrossMatch Telegram bot.
package main

import (
	"fmt"
	"log/slog"
	"os"
)

func main() {
	if err := run(); err != nil {
		slog.Error("crossmatch stopped", "error", err)
		os.Exit(1)
	}
}

func run() error {
	if os.Getenv("BOT_TOKEN") == "" {
		return fmt.Errorf("BOT_TOKEN is not set")
	}
	slog.Info("crossmatch started")
	return nil
}
