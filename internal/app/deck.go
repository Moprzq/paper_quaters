package app

import (
	"fmt"
	"math/rand"
	"path"

	"github.com/hajimehoshi/ebiten/v2"
)

type Card struct {
	Value     int
	Type      string
	CardImage *ebiten.Image
	TypeImage *ebiten.Image
}

type Deck struct {
	Cards  []Card
	Stack1 []Card
	Stack2 []Card
	Stack3 []Card
}

func newDeck(cache *imageCache) (Deck, error) {
	cards, err := loadCards(cache)
	if err != nil {
		return Deck{}, err
	}
	if len(cards) != cardsInDeck {
		return Deck{}, fmt.Errorf("expected %d cards, got %d", cardsInDeck, len(cards))
	}

	deck := Deck{Cards: cards}
	deck.Shuffle()
	return deck, nil
}

func loadCards(cache *imageCache) ([]Card, error) {
	types := []string{"park", "price", "stroiteli", "water", "workers", "zabor"}

	cards := make([]Card, 0, cardsInDeck)
	for _, cardType := range types {
		dir := "cards/" + cardType
		files, err := readDirSorted(dir)
		if err != nil {
			return nil, err
		}

		typeImage, err := cache.load(path.Join(dir, cardType+".jpg"))
		if err != nil {
			return nil, err
		}

		for _, file := range files {
			if file.IsDir() || file.Name() == cardType+".jpg" {
				continue
			}

			value, err := cardValueFromFileName(file.Name())
			if err != nil {
				return nil, err
			}

			cardImage, err := cache.load(path.Join(dir, file.Name()))
			if err != nil {
				return nil, err
			}

			cards = append(cards, Card{
				Value:     value,
				Type:      cardType,
				CardImage: cardImage,
				TypeImage: typeImage,
			})
		}
	}

	return cards, nil
}

func (d *Deck) Shuffle() {
	for i := range d.Cards {
		j := rand.Intn(i + 1)
		d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
	}

	d.Stack1 = d.Cards[:cardsInStack]
	d.Stack2 = d.Cards[cardsInStack : cardsInStack*2]
	d.Stack3 = d.Cards[cardsInStack*2:]
}
