// Command crossmatch is the CrossMatch Telegram bot.
package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/mokevnin/crossmatch/internal/bot"
	"github.com/mokevnin/crossmatch/internal/config"
	"github.com/mokevnin/crossmatch/migrations"
)

func main() {
	os.Exit(realMain())
}

// realMain returns the exit code, so that its deferred calls run before the program exits.
func realMain() int {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	if err := run(ctx); err != nil {
		slog.Error("crossmatch stopped", "error", err)
		return 1
	}
	return 0
}

func run(ctx context.Context) error {
	cfg, err := config.Load(os.Getenv)
	if err != nil {
		return err
	}

	pool, err := pgxpool.New(ctx, cfg.DatabaseURL)
	if err != nil {
		return fmt.Errorf("connect to the database: %w", err)
	}
	defer pool.Close()

	if err := migrations.Up(ctx, pool); err != nil {
		return err
	}

	b, err := bot.New(cfg.BotToken)
	if err != nil {
		return err
	}

	slog.Info("crossmatch started")
	b.Start(ctx)
	slog.Info("crossmatch stopped")
	return nil
}
