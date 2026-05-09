//go:build js && wasm

package app

import (
	"net/url"
	"strings"
	"syscall/js"
)

func defaultLanguage() Language {
	location := js.Global().Get("location")
	if location.IsUndefined() || location.IsNull() {
		return LanguageRU
	}

	values, err := url.ParseQuery(strings.TrimPrefix(location.Get("search").String(), "?"))
	if err != nil {
		return LanguageRU
	}

	language, err := normalizeLanguage(values.Get("lang"))
	if err != nil {
		return LanguageRU
	}
	return language
}
