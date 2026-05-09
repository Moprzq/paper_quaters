package app

import (
	"fmt"
	"strings"
)

type Language string

const (
	LanguageRU  Language = "ru"
	LanguageENG Language = "eng"
)

type uiText struct {
	windowTitle string
	turn        string
	housesBuilt string
	nextTurn    string
	back        string
	shuffle     string
	restart     string
	hint        string
}

func normalizeLanguage(value string) (Language, error) {
	switch strings.ToLower(strings.TrimSpace(value)) {
	case "", string(LanguageRU):
		return LanguageRU, nil
	case string(LanguageENG), "en":
		return LanguageENG, nil
	default:
		return "", fmt.Errorf("unsupported language %q: use ru or eng", value)
	}
}

func uiTextFor(language Language) uiText {
	switch language {
	case LanguageENG:
		return uiText{
			windowTitle: "Paper Quarters",
			turn:        "Turn: %d/%d",
			housesBuilt: "Houses built: %d",
			nextTurn:    "Next turn",
			back:        "Back",
			shuffle:     "Shuffle",
			restart:     "Restart",
			hint:        "Space/Right: Next turn   Left: Back   S: Shuffle   R: Restart   Q/Esc: Exit   F11: Fullscreen",
		}
	default:
		return uiText{
			windowTitle: "Бумажные кварталы",
			turn:        "Ход: %d/%d",
			housesBuilt: "Домов построено: %d",
			nextTurn:    "Следующий ход",
			back:        "Назад",
			shuffle:     "Перемешать",
			restart:     "Рестарт",
			hint:        "Пробел/Вправо: ход   Влево: назад   S: перемешать   R: рестарт   Q/Esc: выход   F11: экран",
		}
	}
}
