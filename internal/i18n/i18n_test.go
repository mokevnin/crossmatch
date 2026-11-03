package i18n

import (
	"encoding/json"
	"regexp"
	"slices"
	"strings"
	"testing"

	"golang.org/x/text/language"
)

func newCatalog(t *testing.T) *Catalog {
	t.Helper()
	c, err := New()
	if err != nil {
		t.Fatalf("New: %v", err)
	}
	return c
}

func TestAnIDNamesItsKey(t *testing.T) {
	if got := StartIntro.Key(); got != "start.intro" {
		t.Errorf("StartIntro.Key() = %q, want start.intro", got)
	}
}

func TestForPicksTheLanguage(t *testing.T) {
	c := newCatalog(t)
	tests := []struct {
		code string
		want language.Tag
	}{
		{"ru", language.Russian},
		{"ru-RU", language.Russian},
		{"en", language.English},
		{"en-US", language.English},
		{"", language.Russian},
		{"de", language.Russian},
		{"garbage", language.Russian},
	}
	for _, tt := range tests {
		t.Run(tt.code, func(t *testing.T) {
			if got := c.For(tt.code).Tag(); got != tt.want {
				t.Errorf("For(%q) speaks %v, want %v", tt.code, got, tt.want)
			}
		})
	}
}

func TestTRendersInTheUsersLanguage(t *testing.T) {
	c := newCatalog(t)
	if got := c.For("en").T(StartIntro, nil); !strings.Contains(got, "I'm CrossMatch") {
		t.Errorf("English start text = %q", got)
	}
	if got := c.For("ru").T(StartIntro, nil); !strings.Contains(got, "Я CrossMatch") {
		t.Errorf("Russian start text = %q", got)
	}
}

func TestAMissingMessageShowsItsKey(t *testing.T) {
	c := newCatalog(t)
	if got := c.For("en").T(ID{key: "no.such.message"}, nil); got != "no.such.message" {
		t.Errorf("T of an unknown message = %q, want its key", got)
	}
}

// The tests below keep the catalogs in step with the Russian one, which is the source of truth.

var placeholder = regexp.MustCompile(`\{\{\s*\.(\w+)`)

// pluralForms are the keys of a plural message that carry text; anything else (a description) does not.
var pluralForms = []string{"zero", "one", "two", "few", "many", "other"}

// requiredForms are the plural forms CLDR gives each language (https://cldr.unicode.org).
var requiredForms = map[language.Tag][]string{
	language.Russian: {"one", "few", "many", "other"},
	language.English: {"one", "other"},
}

// message is a catalog entry: a plain string or an object with one text per plural form.
type message struct {
	texts map[string]string // form -> text; a plain string is stored under "other"
	plain bool
}

func loadCatalog(t *testing.T, tag language.Tag) map[string]message {
	t.Helper()
	raw, err := localesFS.ReadFile("locales/" + tag.String() + ".json")
	if err != nil {
		t.Fatalf("read the %v catalog: %v", tag, err)
	}
	var entries map[string]json.RawMessage
	if err := json.Unmarshal(raw, &entries); err != nil {
		t.Fatalf("parse the %v catalog: %v", tag, err)
	}
	out := make(map[string]message, len(entries))
	for key, value := range entries {
		var text string
		if json.Unmarshal(value, &text) == nil {
			out[key] = message{texts: map[string]string{"other": text}, plain: true}
			continue
		}
		var object map[string]any
		if err := json.Unmarshal(value, &object); err != nil {
			t.Fatalf("%v catalog, %q: neither a string nor an object", tag, key)
		}
		texts := map[string]string{}
		for _, form := range pluralForms {
			if s, ok := object[form].(string); ok {
				texts[form] = s
			}
		}
		out[key] = message{texts: texts}
	}
	return out
}

func placeholders(m message) []string {
	set := map[string]bool{}
	for _, text := range m.texts {
		for _, match := range placeholder.FindAllStringSubmatch(text, -1) {
			set[match[1]] = true
		}
	}
	names := make([]string, 0, len(set))
	for name := range set {
		names = append(names, name)
	}
	slices.Sort(names)
	return names
}

func TestCatalogsHaveTheSameKeysAsRussian(t *testing.T) {
	source := loadCatalog(t, supported[0])
	for _, tag := range supported[1:] {
		other := loadCatalog(t, tag)
		for key := range source {
			if _, ok := other[key]; !ok {
				t.Errorf("%v catalog lacks %q", tag, key)
			}
		}
		for key := range other {
			if _, ok := source[key]; !ok {
				t.Errorf("%v catalog has %q, which the Russian one does not", tag, key)
			}
		}
	}
}

func TestCatalogsUseTheSamePlaceholdersAsRussian(t *testing.T) {
	source := loadCatalog(t, supported[0])
	for _, tag := range supported[1:] {
		for key, m := range loadCatalog(t, tag) {
			want, ok := source[key]
			if !ok {
				continue
			}
			if got, wantNames := placeholders(m), placeholders(want); !slices.Equal(got, wantNames) {
				t.Errorf("%v %q uses placeholders %v, the Russian one %v", tag, key, got, wantNames)
			}
		}
	}
}

func TestPluralMessagesHaveEveryFormTheirLanguageNeeds(t *testing.T) {
	for _, tag := range supported {
		required, ok := requiredForms[tag]
		if !ok {
			t.Fatalf("requiredForms has no entry for %v: add its CLDR plural forms", tag)
		}
		for key, m := range loadCatalog(t, tag) {
			if m.plain {
				continue
			}
			for _, form := range required {
				if _, ok := m.texts[form]; !ok {
					t.Errorf("%v %q lacks the %q plural form", tag, key, form)
				}
			}
		}
	}
}

func TestEveryMessageRendersInEveryLanguage(t *testing.T) {
	c := newCatalog(t)
	for _, tag := range supported {
		loc := c.For(tag.String())
		for key, m := range loadCatalog(t, tag) {
			if !m.plain {
				continue // a plural message needs a count; the form test covers it
			}
			data := map[string]any{}
			for _, name := range placeholders(m) {
				data[name] = "x"
			}
			if got := loc.T(ID{key: key}, data); got == key {
				t.Errorf("%v %q does not render", tag, key)
			}
		}
	}
}
