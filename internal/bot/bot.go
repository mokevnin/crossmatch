// Package bot is the Telegram side of CrossMatch: it receives updates and answers them.
package bot

import (
	"context"
	"fmt"
	"log/slog"

	tgbot "github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"

	"github.com/mokevnin/crossmatch/internal/i18n"
)

// Bot answers Telegram updates.
type Bot struct {
	api   *tgbot.Bot
	texts *i18n.Catalog
}

// New creates the bot. Options go to the Telegram client, which lets a test point it at a fake server.
func New(token string, texts *i18n.Catalog, opts ...tgbot.Option) (*Bot, error) {
	api, err := tgbot.New(token, opts...)
	if err != nil {
		return nil, fmt.Errorf("create the Telegram client: %w", err)
	}
	b := &Bot{api: api, texts: texts}
	api.RegisterHandler(tgbot.HandlerTypeMessageText, "/start", tgbot.MatchTypePrefix, b.start)
	return b, nil
}

// Start receives updates until ctx is cancelled.
func (b *Bot) Start(ctx context.Context) {
	b.api.Start(ctx)
}

// localizer speaks the language Telegram reports for the sender of a message.
func (b *Bot) localizer(m *models.Message) *i18n.Localizer {
	code := ""
	if m.From != nil {
		code = m.From.LanguageCode
	}
	return b.texts.For(code)
}

// Process handles one update as if Telegram had delivered it; tests use it instead of polling.
func (b *Bot) Process(ctx context.Context, update *models.Update) {
	b.api.ProcessUpdate(ctx, update)
}

func (b *Bot) start(ctx context.Context, api *tgbot.Bot, update *models.Update) {
	if update.Message == nil {
		return
	}
	_, err := api.SendMessage(ctx, &tgbot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   b.localizer(update.Message).T(i18n.StartIntro, nil),
	})
	if err != nil {
		slog.ErrorContext(ctx, "answer /start", "error", err)
	}
}
