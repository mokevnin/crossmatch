package bot_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	tgbot "github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"

	"github.com/mokevnin/crossmatch/internal/bot"
	"github.com/mokevnin/crossmatch/internal/i18n"
)

type sentMessage struct {
	ChatID string
	Text   string
}

// fakeTelegram answers every Bot API call with success and reports each sendMessage it gets.
func fakeTelegram(t *testing.T) (url string, sent <-chan sentMessage) {
	t.Helper()
	messages := make(chan sentMessage, 4)
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "/sendMessage") {
			// The client sends parameters as a multipart form.
			if err := r.ParseMultipartForm(1 << 20); err == nil {
				messages <- sentMessage{ChatID: r.FormValue("chat_id"), Text: r.FormValue("text")}
			}
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"ok":true,"result":{"message_id":1,"date":0,"chat":{"id":1,"type":"private"}}}`))
	}))
	t.Cleanup(srv.Close)
	return srv.URL, messages
}

func newBot(t *testing.T, serverURL string) *bot.Bot {
	t.Helper()
	texts, err := i18n.New()
	if err != nil {
		t.Fatalf("i18n.New: %v", err)
	}
	b, err := bot.New("test-token", texts, tgbot.WithServerURL(serverURL), tgbot.WithSkipGetMe())
	if err != nil {
		t.Fatalf("New: %v", err)
	}
	return b
}

func TestStartAnswersInTheSameChat(t *testing.T) {
	url, sent := fakeTelegram(t)
	b := newBot(t, url)

	b.Process(context.Background(), &models.Update{
		Message: &models.Message{Text: "/start", Chat: models.Chat{ID: 42}},
	})

	select {
	case m := <-sent:
		if m.ChatID != "42" {
			t.Errorf("answered chat %v, want 42", m.ChatID)
		}
		if !strings.Contains(m.Text, "Я CrossMatch") {
			t.Errorf("answer %q is not the Russian introduction", m.Text)
		}
	case <-time.After(5 * time.Second):
		t.Fatal("no answer to /start")
	}
}

func TestStartAnswersInTheSendersLanguage(t *testing.T) {
	url, sent := fakeTelegram(t)
	b := newBot(t, url)

	b.Process(context.Background(), &models.Update{
		Message: &models.Message{
			Text: "/start",
			Chat: models.Chat{ID: 42},
			From: &models.User{ID: 7, LanguageCode: "en"},
		},
	})

	select {
	case m := <-sent:
		if !strings.Contains(m.Text, "I'm CrossMatch") {
			t.Errorf("answer %q is not the English introduction", m.Text)
		}
	case <-time.After(5 * time.Second):
		t.Fatal("no answer to /start")
	}
}

func TestOtherTextIsNotAnswered(t *testing.T) {
	url, sent := fakeTelegram(t)
	b := newBot(t, url)

	b.Process(context.Background(), &models.Update{
		Message: &models.Message{Text: "hello", Chat: models.Chat{ID: 42}},
	})

	select {
	case m := <-sent:
		t.Errorf("answered an unknown message with %q", m.Text)
	case <-time.After(300 * time.Millisecond):
	}
}
