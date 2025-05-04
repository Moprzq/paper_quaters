package main

import (
	"fmt"
	"image/color"
	"math/rand"
	"os"
	"strconv"
	"strings"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	const windowWidth = 1920
	const windowHeight = 1080
	rl.InitWindow(windowWidth, windowHeight, "Пример Raylib")
	rl.ToggleFullscreen()
	defer rl.CloseWindow()

	rl.SetTargetFPS(30)

	deck := GetDeck()
	deck.Shuffle()
	missions := GetMissons()

	turn := 1
	missionTextureFirst := missions.First.MissionTexture
	missionTextureSecond := missions.Second.MissionTexture
	missionTextureThird := missions.Third.MissionTexture

	firstMissionBounds := rl.Rectangle{float32(windowWidth/2 - missionTextureFirst.Width/2 - 700), float32(windowHeight/2 - missionTextureFirst.Height/2 - 350), float32(missionTextureFirst.Width), float32(missionTextureFirst.Height)}
	secondMissionBounds := rl.Rectangle{float32(windowWidth/2 - missionTextureSecond.Width/2 - 700), float32(windowHeight/2 - missionTextureSecond.Height/2), float32(missionTextureSecond.Width), float32(missionTextureSecond.Height)}
	thirdMissionBounds := rl.Rectangle{float32(windowWidth/2 - missionTextureThird.Width/2 - 700), float32(windowHeight/2 - missionTextureThird.Height/2 + 350), float32(missionTextureThird.Width), float32(missionTextureThird.Height)}

	nextTurnButtonBounds := rl.Rectangle{1490, 440, 330, 82}
	backButtonBounds := rl.Rectangle{1560, 590, 160, 74}
	shuffleButtonBounds := rl.Rectangle{1550, 790, 200, 60}
	restartButtonBounds := rl.Rectangle{1550, 890, 200, 60}

	for !rl.WindowShouldClose() {

		mousePoint := rl.GetMousePosition()

		if rl.IsKeyPressed(rl.KeySpace) || rl.IsKeyPressed(rl.KeyRight) || rl.CheckCollisionPointRec(mousePoint, nextTurnButtonBounds) && rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
			if turn < 26 {
				turn += 1
			}
		} else if rl.IsKeyPressed(rl.KeyN) || rl.CheckCollisionPointRec(mousePoint, restartButtonBounds) && rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
			missions = GetMissons()
			turn = 1
			deck.Shuffle()
			missionTextureFirst = missions.First.MissionTexture
			missionTextureSecond = missions.Second.MissionTexture
			missionTextureThird = missions.Third.MissionTexture
		} else if rl.IsKeyPressed(rl.KeyB) || rl.IsKeyPressed(rl.KeyLeft) || rl.CheckCollisionPointRec(mousePoint, backButtonBounds) && rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
			if turn > 1 {
				turn -= 1
			}
		} else if rl.IsKeyPressed(rl.KeyR) || rl.CheckCollisionPointRec(mousePoint, shuffleButtonBounds) && rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
			turn = 1
			deck.Shuffle()
		} else if rl.IsKeyPressed(rl.KeyOne) || rl.CheckCollisionPointRec(mousePoint, firstMissionBounds) && rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
			missionTextureFirst = missions.First.MissionRevTexture
		} else if rl.IsKeyPressed(rl.KeyTwo) || rl.CheckCollisionPointRec(mousePoint, secondMissionBounds) && rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
			missionTextureSecond = missions.Second.MissionRevTexture
		} else if rl.IsKeyPressed(rl.KeyThree) || rl.CheckCollisionPointRec(mousePoint, thirdMissionBounds) && rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
			missionTextureThird = missions.Third.MissionRevTexture
		}

		rl.BeginDrawing()
		// rl.ClearBackground(rl.NewColor(245, 222, 179, 255))
		rl.ClearBackground(rl.NewColor(255, 235, 205, 255))

		rl.DrawTexture(deck.Stack1[turn-1].TypeTexture, windowWidth/2-deck.Stack1[turn-1].TypeTexture.Width/2-300, windowHeight/2-deck.Stack1[turn-1].TypeTexture.Height/2+200, rl.White)
		rl.DrawTexture(deck.Stack2[turn-1].TypeTexture, windowWidth/2-deck.Stack2[turn-1].TypeTexture.Width/2, windowHeight/2-deck.Stack2[turn-1].TypeTexture.Height/2+200, rl.White)
		rl.DrawTexture(deck.Stack3[turn-1].TypeTexture, windowWidth/2-deck.Stack3[turn-1].TypeTexture.Width/2+300, windowHeight/2-deck.Stack3[turn-1].TypeTexture.Height/2+200, rl.White)

		rl.DrawTexture(deck.Stack1[turn].CardTexture, windowWidth/2-deck.Stack1[turn].CardTexture.Width/2-300, windowHeight/2-deck.Stack1[turn].CardTexture.Height/2-200, rl.White)
		rl.DrawTexture(deck.Stack2[turn].CardTexture, windowWidth/2-deck.Stack2[turn].CardTexture.Width/2, windowHeight/2-deck.Stack2[turn].CardTexture.Height/2-200, rl.White)
		rl.DrawTexture(deck.Stack3[turn].CardTexture, windowWidth/2-deck.Stack3[turn].CardTexture.Width/2+300, windowHeight/2-deck.Stack3[turn].CardTexture.Height/2-200, rl.White)

		//misson complite logic

		rl.DrawTexture(missionTextureFirst, windowWidth/2-missionTextureFirst.Width/2-700, windowHeight/2-missionTextureFirst.Height/2-350, rl.White)
		rl.DrawTexture(missionTextureSecond, windowWidth/2-missionTextureSecond.Width/2-700, windowHeight/2-missionTextureSecond.Height/2, rl.White)
		rl.DrawTexture(missionTextureThird, windowWidth/2-missionTextureThird.Width/2-700, windowHeight/2-missionTextureThird.Height/2+350, rl.White)

		rl.DrawText(fmt.Sprintf("Turn: %d", turn), 1575, 100, 40, rl.Black)
		rl.DrawRectangle(1550, 90, 200, 60, color.RGBA{0, 0, 0, 50})

		rl.DrawText("Next turn", 1500, 450, 62, rl.Black)
		rl.DrawRectangle(1490, 440, 330, 82, color.RGBA{0, 0, 0, 50})

		rl.DrawText("Back", 1575, 600, 54, rl.Black)
		rl.DrawRectangle(1560, 590, 160, 74, color.RGBA{0, 0, 0, 50})

		rl.DrawText("Shuffle", 1575, 800, 40, rl.Black)
		rl.DrawRectangle(1550, 790, 200, 60, color.RGBA{0, 0, 0, 50})

		rl.DrawText("Restart", 1575, 900, 40, rl.Black)
		rl.DrawRectangle(1550, 890, 200, 60, color.RGBA{0, 0, 0, 50})

		rl.EndDrawing()
	}

}

type Card struct {
	Value       int
	Type        string
	CardTexture rl.Texture2D
	TypeTexture rl.Texture2D
}

type Deck struct {
	Cards  []Card
	Stack1 []Card
	Stack2 []Card
	Stack3 []Card
}

func GetDeck() Deck {
	cards := GetCards()
	return Deck{
		Cards:  cards,
		Stack1: cards[:27],
		Stack2: cards[27:54],
		Stack3: cards[54:],
	}
}

func GetCards() []Card {
	types := []string{"park", "price", "stroiteli", "water", "workers", "zabor"}

	cards := make([]Card, 0, 81)
	for _, t := range types {
		dirPath := "./cards/" + t
		files, err := os.ReadDir(dirPath)
		if err != nil {
			fmt.Println("err to read dir: ", dirPath)
		}

		typeImage := rl.LoadImage(dirPath + "/" + t + ".JPG")
		rl.ImageResize(typeImage, 200, 300)
		typeTexture := rl.LoadTextureFromImage(typeImage)
		rl.UnloadImage(typeImage)

		for _, file := range files {
			if file.Name() == t+".JPG" {
				fmt.Println("type succeful skipped")
				continue
			}
			imagePath := dirPath + "/" + file.Name()
			cardImage := rl.LoadImage(imagePath)
			rl.ImageResize(cardImage, 200, 300)
			cardTexture := rl.LoadTextureFromImage(cardImage)
			rl.UnloadImage(cardImage)

			value, err := strconv.ParseInt(strings.Split(file.Name(), ".")[0], 10, 64)
			if err != nil {
				fmt.Println("err parse file value: ", file.Name())
			}

			cards = append(cards, Card{
				Value:       int(value),
				Type:        t,
				CardTexture: cardTexture,
				TypeTexture: typeTexture,
			})
		}
	}
	return cards
}

func (d *Deck) Shuffle() {
	c := d.Cards
	for i := range c {
		j := rand.Intn(i + 1)
		c[i], c[j] = c[j], c[i]
	}
	d.Stack1 = c[:27]
	d.Stack2 = c[27:54]
	d.Stack3 = c[54:]
}

type Mission struct {
	MissionTexture    rl.Texture2D
	MissionRevTexture rl.Texture2D
}

type GameMissons struct {
	First  Mission
	Second Mission
	Third  Mission
}

func GetMissons() GameMissons {
	mission := []string{"1", "2", "3"}
	missions := make([]Mission, 0, 3)
	for _, m := range mission {

		missionNumber := rand.Intn(6) + 1
		dirPath := "./missons/" + m

		missionImage := rl.LoadImage(dirPath + "/" + strconv.FormatInt(int64(missionNumber), 10) + ".JPG")
		rl.ImageResize(missionImage, 200, 300)
		missionTexture := rl.LoadTextureFromImage(missionImage)
		rl.UnloadImage(missionImage)

		missionRevImage := rl.LoadImage(dirPath + "/" + strconv.FormatInt(int64(missionNumber), 10) + "_rev" + ".JPG")
		rl.ImageResize(missionRevImage, 200, 300)
		missionRevTexture := rl.LoadTextureFromImage(missionRevImage)
		rl.UnloadImage(missionRevImage)

		missions = append(missions, Mission{
			MissionTexture:    missionTexture,
			MissionRevTexture: missionRevTexture,
		})
	}
	return GameMissons{
		First:  missions[0],
		Second: missions[1],
		Third:  missions[2],
	}
}
