//go:build !js || !wasm

package app

func defaultLanguage() Language {
	return LanguageRU
}
