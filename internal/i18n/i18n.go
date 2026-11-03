// Package i18n holds every text the bot shows a user, in each language it speaks.
//
// The catalogs are locales/*.json, embedded in the binary. Russian is the source of truth and the
// fallback: a message another catalog lacks comes through in Russian. A message is named by an ID
// that cmd/i18ngen generates from locales/ru.json (ids_gen.go), so a call can only name a message
// that exists.
package i18n

import (
	"embed"
	"encoding/json"
	"fmt"
	"log/slog"

	goi18n "github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

//go:embed locales/*.json
var localesFS embed.FS

// supported lists the languages the bot speaks. The first is the default and the fallback; it is
// also the catalog ids_gen.go is generated from.
var supported = []language.Tag{language.Russian, language.English}

// ID names a message of the catalog. Only the generated ids (ids_gen.go) make one.
type ID struct{ key string }

// Key is the id's key in the catalog (start.intro).
func (id ID) Key() string { return id.key }

// Catalog holds the loaded messages of every language.
type Catalog struct {
	bundle  *goi18n.Bundle
	matcher language.Matcher
}

// New loads the embedded catalogs. A catalog that fails to load is a build mistake, so it is an error.
func New() (*Catalog, error) {
	bundle := goi18n.NewBundle(supported[0])
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
	for _, tag := range supported {
		if _, err := bundle.LoadMessageFileFS(localesFS, "locales/"+tag.String()+".json"); err != nil {
			return nil, fmt.Errorf("load the %s catalog: %w", tag, err)
		}
	}
	return &Catalog{bundle: bundle, matcher: language.NewMatcher(supported)}, nil
}

// For returns the localizer for a Telegram language code such as "en", "ru-RU" or "" when
// Telegram sent none. A language the bot does not speak gets the default one.
func (c *Catalog) For(code string) *Localizer {
	_, index, confidence := c.matcher.Match(language.Make(code))
	if confidence == language.No {
		index = 0
	}
	tag := supported[index]
	return &Localizer{
		tag:       tag,
		localizer: goi18n.NewLocalizer(c.bundle, tag.String(), supported[0].String()),
	}
}

// Localizer renders messages in one language.
type Localizer struct {
	tag       language.Tag
	localizer *goi18n.Localizer
}

// Tag is the language the localizer speaks.
func (l *Localizer) Tag() language.Tag { return l.tag }

// T renders a message, filling {{.Name}} placeholders from data. It never fails: if a message
// cannot be rendered (a placeholder without data), the id's key is shown and the mistake logged,
// and the tests that every catalog agrees with the Russian one keep that from shipping.
func (l *Localizer) T(id ID, data map[string]any) string {
	text, err := l.localizer.Localize(&goi18n.LocalizeConfig{MessageID: id.key, TemplateData: data})
	if err != nil {
		slog.Error("render a message", "message", id.key, "language", l.tag, "error", err)
		return id.key
	}
	return text
}
