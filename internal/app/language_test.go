package app

import "testing"

func TestNormalizeLanguage(t *testing.T) {
	tests := map[string]Language{
		"":    LanguageRU,
		"ru":  LanguageRU,
		"eng": LanguageENG,
		"en":  LanguageENG,
		"ENG": LanguageENG,
	}

	for value, want := range tests {
		got, err := normalizeLanguage(value)
		if err != nil {
			t.Fatalf("normalizeLanguage(%q) error = %v", value, err)
		}
		if got != want {
			t.Fatalf("normalizeLanguage(%q) = %q, want %q", value, got, want)
		}
	}
}

func TestNormalizeLanguageRejectsUnsupportedLanguage(t *testing.T) {
	if _, err := normalizeLanguage("de"); err == nil {
		t.Fatal("normalizeLanguage(\"de\") error = nil, want error")
	}
}
