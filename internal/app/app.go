package main

import (
	"fmt"
	"os"

	"github.com/eiannone/keyboard"

	"paper_quarters/internal/cards"
	"paper_quarters/internal/pkg/clear"
)

func main() {
	deck := cards.GetGameDeck()
	turn := 1
	missons := cards.GetMissons()

	keysEvents, err := keyboard.GetKeys(10)
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()
	clear.CallClearTerminal()
	printTurn(missons, deck, turn, "")
	for {
		msg := ""
		event := <-keysEvents
		if event.Err != nil {
			panic(event.Err)
		}
		if event.Key != keyboard.KeyEsc && event.Key != keyboard.KeySpace &&
			event.Rune != 'r' && event.Rune != 'R' &&
			event.Rune != 'B' && event.Rune != 'b' &&
			event.Rune != 'N' && event.Rune != 'n' &&
			event.Rune != '1' && event.Rune != '2' && event.Rune != '3' {
			continue
		}
		if event.Key == keyboard.KeySpace {
			if turn < 26 {
				turn += 1
			} else {
				msg = "For next turn you must shuffle the deck"
			}
		} else if event.Rune == 'r' || event.Rune == 'R' {
			msg = "The deck is shuffled"
			deck = cards.GetGameDeck()
			turn = 1
		} else if event.Rune == '1' || event.Rune == '2' || event.Rune == '3' {
			if event.Rune == '1' {
				missons.First.IsAlreadyDone = true
			} else if event.Rune == '2' {
				missons.Second.IsAlreadyDone = true
			} else {
				missons.Third.IsAlreadyDone = true
			}
		} else if event.Rune == 'B' || event.Rune == 'b' {
			if turn > 1 {
				turn -= 1
			}
			msg = "One move back"
		} else if event.Key == keyboard.KeyEsc {
			os.Exit(1)
		} else if event.Rune == 'N' || event.Rune == 'n' {
			deck = cards.GetGameDeck()
			turn = 1
			missons = cards.GetMissons()
		}
		clear.CallClearTerminal()
		printTurn(missons, deck, turn, msg)
	}
}

var IsAlreadyDone = map[bool]string{
	true: "🚫",
}

func printTurn(missons cards.GameMissons, deck cards.GameDeck, turn int, msg string) {
	fmt.Printf("Turn: %d/%d\nMissons:\n", turn, 26)
	fmt.Printf("\t%d. Buildings: %v \n\tPoints: %d%s %d\n",
		missons.First.MissionNumber,
		missons.First.Mission(),
		missons.First.ScoreFirst,
		IsAlreadyDone[missons.First.IsAlreadyDone],
		missons.First.ScoreSecond)
	fmt.Printf("\t%d. Buildings: %v \n\tPoints: %d%s %d\n",
		missons.Second.MissionNumber,
		missons.Second.Mission(),
		missons.Second.ScoreFirst,
		IsAlreadyDone[missons.Second.IsAlreadyDone],
		missons.Second.ScoreSecond)
	fmt.Printf("\t%d. Buildings: %v \n\tPoints: %d%s %d\n",
		missons.Third.MissionNumber,
		missons.Third.Mission(),
		missons.Third.ScoreFirst,
		IsAlreadyDone[missons.Third.IsAlreadyDone],
		missons.Third.ScoreSecond)

	fmt.Printf("\nCurrent turn:\t%d%s\t%d%s\t%d%s\n\n",
		deck.Stack1[turn].Value, deck.Stack1[turn-1].Type,
		deck.Stack2[turn].Value, deck.Stack2[turn-1].Type,
		deck.Stack3[turn].Value, deck.Stack3[turn-1].Type)
	if turn < 26 {
		fmt.Printf("Next types: \t%s \t%s \t%s\n", deck.Stack1[turn].Type, deck.Stack2[turn].Type, deck.Stack3[turn].Type)
	}
	if msg != "" {
		fmt.Println(msg)
	}
}
