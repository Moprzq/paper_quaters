package app

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	screenWidth  = 1920
	screenHeight = 1080

	cardsInDeck     = 81
	cardsInStack    = 27
	lastTurn        = 26
	missionsInLevel = 6

	cardWidth  = 200
	cardHeight = 300
)

type Rect struct {
	X float64
	Y float64
	W float64
	H float64
}

func (r Rect) Contains(x, y int) bool {
	return float64(x) >= r.X &&
		float64(x) <= r.X+r.W &&
		float64(y) >= r.Y &&
		float64(y) <= r.Y+r.H
}

type Game struct {
	cache    *imageCache
	deck     Deck
	missions GameMissions
	turn     int
}

func newGame() (*Game, error) {
	cache := newImageCache()
	deck, err := newDeck(cache)
	if err != nil {
		return nil, err
	}

	missions, err := newMissions(cache)
	if err != nil {
		return nil, err
	}

	return &Game{
		cache:    cache,
		deck:     deck,
		missions: missions,
		turn:     1,
	}, nil
}

func (g *Game) Update() error {
	mouseX, mouseY := ebiten.CursorPosition()

	switch {
	case inpututil.IsKeyJustPressed(ebiten.KeySpace),
		inpututil.IsKeyJustPressed(ebiten.KeyArrowRight),
		clicked(nextTurnButton(), mouseX, mouseY):
		g.nextTurn()
	case inpututil.IsKeyJustPressed(ebiten.KeyN),
		clicked(restartButton(), mouseX, mouseY):
		if err := g.restart(); err != nil {
			return err
		}
	case inpututil.IsKeyJustPressed(ebiten.KeyB),
		inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft),
		clicked(backButton(), mouseX, mouseY):
		g.previousTurn()
	case inpututil.IsKeyJustPressed(ebiten.KeyR),
		clicked(shuffleButton(), mouseX, mouseY):
		g.shuffleDeck()
	case inpututil.IsKeyJustPressed(ebiten.Key1),
		clicked(firstMissionRect(), mouseX, mouseY):
		g.missions.First.Done = true
	case inpututil.IsKeyJustPressed(ebiten.Key2),
		clicked(secondMissionRect(), mouseX, mouseY):
		g.missions.Second.Done = true
	case inpututil.IsKeyJustPressed(ebiten.Key3),
		clicked(thirdMissionRect(), mouseX, mouseY):
		g.missions.Third.Done = true
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{R: 255, G: 235, B: 205, A: 255})

	g.drawTurnCards(screen)
	g.drawMissions(screen)
	g.drawControls(screen)
}

func (g *Game) Layout(_, _ int) (int, int) {
	return screenWidth, screenHeight
}

func (g *Game) nextTurn() {
	if g.turn < lastTurn {
		g.turn++
	}
}

func (g *Game) previousTurn() {
	if g.turn > 1 {
		g.turn--
	}
}

func (g *Game) shuffleDeck() {
	g.turn = 1
	g.deck.Shuffle()
}

func (g *Game) restart() error {
	missions, err := newMissions(g.cache)
	if err != nil {
		return err
	}

	g.turn = 1
	g.deck.Shuffle()
	g.missions = missions
	return nil
}

func (g *Game) drawTurnCards(screen *ebiten.Image) {
	typeY := float64(screenHeight/2 - cardHeight/2 + 200)
	cardY := float64(screenHeight/2 - cardHeight/2 - 200)

	drawCardImage(screen, g.deck.Stack1[g.turn-1].TypeImage, float64(screenWidth/2-cardWidth/2-300), typeY)
	drawCardImage(screen, g.deck.Stack2[g.turn-1].TypeImage, float64(screenWidth/2-cardWidth/2), typeY)
	drawCardImage(screen, g.deck.Stack3[g.turn-1].TypeImage, float64(screenWidth/2-cardWidth/2+300), typeY)

	drawCardImage(screen, g.deck.Stack1[g.turn].CardImage, float64(screenWidth/2-cardWidth/2-300), cardY)
	drawCardImage(screen, g.deck.Stack2[g.turn].CardImage, float64(screenWidth/2-cardWidth/2), cardY)
	drawCardImage(screen, g.deck.Stack3[g.turn].CardImage, float64(screenWidth/2-cardWidth/2+300), cardY)
}

func (g *Game) drawMissions(screen *ebiten.Image) {
	drawMission(screen, g.missions.First, firstMissionRect())
	drawMission(screen, g.missions.Second, secondMissionRect())
	drawMission(screen, g.missions.Third, thirdMissionRect())
}

func (g *Game) drawControls(screen *ebiten.Image) {
	drawLabel(screen, fmt.Sprintf("Turn: %d/%d", g.turn, lastTurn), Rect{X: 1550, Y: 90, W: 200, H: 60})
	drawButton(screen, "Next turn", nextTurnButton())
	drawButton(screen, "Back", backButton())
	drawButton(screen, "Shuffle", shuffleButton())
	drawButton(screen, "Restart", restartButton())
}

func drawMission(screen *ebiten.Image, mission Mission, rect Rect) {
	img := mission.Front
	if mission.Done {
		img = mission.Back
	}
	drawCardImage(screen, img, rect.X, rect.Y)
}

func drawCardImage(screen *ebiten.Image, img *ebiten.Image, x, y float64) {
	bounds := img.Bounds()
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Scale(float64(cardWidth)/float64(bounds.Dx()), float64(cardHeight)/float64(bounds.Dy()))
	opts.GeoM.Translate(x, y)
	screen.DrawImage(img, opts)
}

func drawButton(screen *ebiten.Image, label string, rect Rect) {
	drawLabel(screen, label, rect)
}

func drawLabel(screen *ebiten.Image, label string, rect Rect) {
	ebitenutil.DrawRect(screen, rect.X, rect.Y, rect.W, rect.H, color.RGBA{A: 50})
	ebitenutil.DebugPrintAt(screen, label, int(rect.X)+16, int(rect.Y)+22)
}

func clicked(rect Rect, mouseX, mouseY int) bool {
	return inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) && rect.Contains(mouseX, mouseY)
}

func firstMissionRect() Rect {
	return Rect{X: float64(screenWidth/2 - cardWidth/2 - 700), Y: float64(screenHeight/2 - cardHeight/2 - 350), W: cardWidth, H: cardHeight}
}

func secondMissionRect() Rect {
	return Rect{X: float64(screenWidth/2 - cardWidth/2 - 700), Y: float64(screenHeight/2 - cardHeight/2), W: cardWidth, H: cardHeight}
}

func thirdMissionRect() Rect {
	return Rect{X: float64(screenWidth/2 - cardWidth/2 - 700), Y: float64(screenHeight/2 - cardHeight/2 + 350), W: cardWidth, H: cardHeight}
}

func nextTurnButton() Rect {
	return Rect{X: 1490, Y: 440, W: 330, H: 82}
}

func backButton() Rect {
	return Rect{X: 1560, Y: 590, W: 160, H: 74}
}

func shuffleButton() Rect {
	return Rect{X: 1550, Y: 790, W: 200, H: 60}
}

func restartButton() Rect {
	return Rect{X: 1550, Y: 890, W: 200, H: 60}
}

func Run() error {
	game, err := newGame()
	if err != nil {
		return err
	}

	ebiten.SetWindowSize(1280, 720)
	ebiten.SetWindowTitle("Paper Quarters")
	return ebiten.RunGame(game)
}
