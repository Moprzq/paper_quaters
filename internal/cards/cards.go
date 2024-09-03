package cards

import "math/rand"

type Card struct {
	Type  string
	Value int
}

type GameDeck struct {
	Stack1 []Card
	Stack2 []Card
	Stack3 []Card
}

var cards = map[string][]int{
	"♒︎": {1, 2, 3, 5, 5, 6, 6, 7, 8, 8, 9, 10, 10, 11, 11, 13, 14, 15},
	"💸":  {1, 2, 4, 5, 5, 6, 7, 7, 8, 8, 9, 9, 10, 11, 11, 12, 14, 15},
	"📭":  {3, 4, 6, 7, 8, 9, 10, 12, 13},
	"🚧":  {3, 4, 6, 7, 8, 9, 10, 12, 13},
	"🌊":  {3, 4, 6, 7, 8, 9, 10, 12, 13},
	"🌳":  {1, 2, 4, 5, 5, 6, 7, 7, 8, 8, 9, 9, 10, 11, 11, 12, 14, 15},
}

func GetGameDeck() GameDeck {
	deck := GameDeck{}
	c := GetShuffledCards()
	deck.Stack1 = c[:27]
	deck.Stack2 = c[27:54]
	deck.Stack3 = c[54:]
	return deck
}

func GetShuffledCards() []Card {
	Cards := make([]Card, 0, 81)
	for t, values := range cards {
		for _, value := range values {
			Cards = append(Cards, Card{
				Type:  t,
				Value: value,
			})
		}
	}
	for i := range Cards {
		j := rand.Intn(i + 1)
		Cards[i], Cards[j] = Cards[j], Cards[i]
	}
	return Cards
}
