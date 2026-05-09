package app

import (
	"bytes"
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/font/gofont/goregular"
)

const (
	screenWidth  = 1920
	screenHeight = 1080

	cardsInDeck     = 81
	cardsInStack    = 27
	lastTurn        = 26
	missionsInLevel = 6

	missionCardWidth  = 200
	missionCardHeight = 300
	turnCardWidth     = 240
	turnCardHeight    = 360

	controlFontSize   = 32
	hintFontSize      = 24
	cardCornerRadius  = 14
	labelCornerRadius = 10
)

var controlFaceSource = mustLoadControlFaceSource()

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
	cache          *imageCache
	deck           Deck
	missions       GameMissions
	turn           int
	housesBuilt    int
	maxCountedTurn int
	roundedCards   map[roundedCardKey]*ebiten.Image
}

type roundedCardKey struct {
	img    *ebiten.Image
	width  int
	height int
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
		cache:          cache,
		deck:           deck,
		missions:       missions,
		turn:           1,
		maxCountedTurn: 1,
		roundedCards:   make(map[roundedCardKey]*ebiten.Image),
	}, nil
}

func (g *Game) Update() error {
	mouseX, mouseY := ebiten.CursorPosition()

	switch {
	case inpututil.IsKeyJustPressed(ebiten.KeyEscape),
		inpututil.IsKeyJustPressed(ebiten.KeyQ):
		return ebiten.Termination
	case inpututil.IsKeyJustPressed(ebiten.KeyF11):
		ebiten.SetFullscreen(!ebiten.IsFullscreen())
	case inpututil.IsKeyJustPressed(ebiten.KeySpace),
		inpututil.IsKeyJustPressed(ebiten.KeyArrowRight),
		clicked(nextTurnButton(), mouseX, mouseY):
		g.nextTurn()
	case inpututil.IsKeyJustPressed(ebiten.KeyR),
		clicked(restartButton(), mouseX, mouseY):
		if err := g.restart(); err != nil {
			return err
		}
	case inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft),
		clicked(backButton(), mouseX, mouseY):
		g.previousTurn()
	case inpututil.IsKeyJustPressed(ebiten.KeyS),
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
		if g.turn > g.maxCountedTurn {
			g.housesBuilt++
			g.maxCountedTurn = g.turn
		}
	}
}

func (g *Game) previousTurn() {
	if g.turn > 1 {
		g.turn--
	}
}

func (g *Game) shuffleDeck() {
	g.turn = 1
	g.maxCountedTurn = 1
	g.deck.Shuffle()
}

func (g *Game) restart() error {
	missions, err := newMissions(g.cache)
	if err != nil {
		return err
	}

	g.turn = 1
	g.housesBuilt = 0
	g.maxCountedTurn = 1
	g.deck.Shuffle()
	g.missions = missions
	return nil
}

func (g *Game) drawTurnCards(screen *ebiten.Image) {
	typeY := float64(screenHeight/2 - turnCardHeight/2 + 210)
	cardY := float64(screenHeight/2 - turnCardHeight/2 - 210)

	g.drawCardImage(screen, g.deck.Stack1[g.turn-1].TypeImage, float64(screenWidth/2-turnCardWidth/2-310), typeY, turnCardWidth, turnCardHeight)
	g.drawCardImage(screen, g.deck.Stack2[g.turn-1].TypeImage, float64(screenWidth/2-turnCardWidth/2), typeY, turnCardWidth, turnCardHeight)
	g.drawCardImage(screen, g.deck.Stack3[g.turn-1].TypeImage, float64(screenWidth/2-turnCardWidth/2+310), typeY, turnCardWidth, turnCardHeight)

	g.drawCardImage(screen, g.deck.Stack1[g.turn].CardImage, float64(screenWidth/2-turnCardWidth/2-310), cardY, turnCardWidth, turnCardHeight)
	g.drawCardImage(screen, g.deck.Stack2[g.turn].CardImage, float64(screenWidth/2-turnCardWidth/2), cardY, turnCardWidth, turnCardHeight)
	g.drawCardImage(screen, g.deck.Stack3[g.turn].CardImage, float64(screenWidth/2-turnCardWidth/2+310), cardY, turnCardWidth, turnCardHeight)
}

func (g *Game) drawMissions(screen *ebiten.Image) {
	g.drawMission(screen, g.missions.First, firstMissionRect())
	g.drawMission(screen, g.missions.Second, secondMissionRect())
	g.drawMission(screen, g.missions.Third, thirdMissionRect())
}

func (g *Game) drawControls(screen *ebiten.Image) {
	drawLabel(screen, fmt.Sprintf("Turn: %d/%d", g.turn, lastTurn), Rect{X: 1550, Y: 90, W: 200, H: 60})
	drawLabel(screen, fmt.Sprintf("Домов построенно: %d", g.housesBuilt), Rect{X: 1440, Y: 190, W: 420, H: 60})
	drawButton(screen, "Next turn", nextTurnButton())
	drawButton(screen, "Back", backButton())
	drawButton(screen, "Shuffle", shuffleButton())
	drawButton(screen, "Restart", restartButton())
	drawLabelWithSize(screen, "Space/Right: Next turn   Left: Back   S: Shuffle   R: Restart   Q/Esc: Exit   F11: Fullscreen", shortcutHintRect(), hintFontSize)
}

func (g *Game) drawMission(screen *ebiten.Image, mission Mission, rect Rect) {
	img := mission.Front
	if mission.Done {
		img = mission.Back
	}
	g.drawCardImage(screen, img, rect.X, rect.Y, int(rect.W), int(rect.H))
}

func (g *Game) drawCardImage(screen *ebiten.Image, img *ebiten.Image, x, y float64, width, height int) {
	card := g.roundedCardImage(img, width, height)
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(x, y)
	screen.DrawImage(card, opts)
}

func (g *Game) roundedCardImage(img *ebiten.Image, width, height int) *ebiten.Image {
	key := roundedCardKey{img: img, width: width, height: height}
	if rounded := g.roundedCards[key]; rounded != nil {
		return rounded
	}

	card := ebiten.NewImage(width, height)
	bounds := img.Bounds()
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Scale(float64(width)/float64(bounds.Dx()), float64(height)/float64(bounds.Dy()))
	card.DrawImage(img, opts)

	mask := ebiten.NewImage(width, height)
	drawRoundedRect(mask, Rect{W: float64(width), H: float64(height)}, cardCornerRadius, color.White)
	maskOpts := &ebiten.DrawImageOptions{
		Blend: destinationInBlend(),
	}
	card.DrawImage(mask, maskOpts)

	g.roundedCards[key] = card
	return card
}

func drawButton(screen *ebiten.Image, label string, rect Rect) {
	drawLabel(screen, label, rect)
}

func drawLabel(screen *ebiten.Image, label string, rect Rect) {
	drawLabelWithSize(screen, label, rect, controlFontSize)
}

func drawLabelWithSize(screen *ebiten.Image, label string, rect Rect, fontSize float64) {
	drawRoundedRect(screen, rect, labelCornerRadius, color.RGBA{A: 50})

	face := &text.GoTextFace{
		Source: controlFaceSource,
		Size:   fontSize,
	}
	textWidth, textHeight := text.Measure(label, face, 0)

	opts := &text.DrawOptions{}
	opts.GeoM.Translate(
		rect.X+(rect.W-textWidth)/2,
		rect.Y+(rect.H-textHeight)/2,
	)
	opts.ColorScale.ScaleWithColor(color.Black)
	text.Draw(screen, label, face, opts)
}

func drawRoundedRect(dst *ebiten.Image, rect Rect, radius float64, clr color.Color) {
	r := float32(radius)
	x := float32(rect.X)
	y := float32(rect.Y)
	w := float32(rect.W)
	h := float32(rect.H)

	var path vector.Path
	path.MoveTo(x+r, y)
	path.LineTo(x+w-r, y)
	path.QuadTo(x+w, y, x+w, y+r)
	path.LineTo(x+w, y+h-r)
	path.QuadTo(x+w, y+h, x+w-r, y+h)
	path.LineTo(x+r, y+h)
	path.QuadTo(x, y+h, x, y+h-r)
	path.LineTo(x, y+r)
	path.QuadTo(x, y, x+r, y)
	path.Close()

	var colorScale ebiten.ColorScale
	colorScale.ScaleWithColor(clr)
	vector.FillPath(dst, &path, &vector.FillOptions{}, &vector.DrawPathOptions{
		AntiAlias:  true,
		ColorScale: colorScale,
	})
}

func destinationInBlend() ebiten.Blend {
	return ebiten.Blend{
		BlendFactorSourceRGB:        ebiten.BlendFactorZero,
		BlendFactorSourceAlpha:      ebiten.BlendFactorZero,
		BlendFactorDestinationRGB:   ebiten.BlendFactorSourceAlpha,
		BlendFactorDestinationAlpha: ebiten.BlendFactorSourceAlpha,
		BlendOperationRGB:           ebiten.BlendOperationAdd,
		BlendOperationAlpha:         ebiten.BlendOperationAdd,
	}
}

func mustLoadControlFaceSource() *text.GoTextFaceSource {
	source, err := text.NewGoTextFaceSource(bytes.NewReader(goregular.TTF))
	if err != nil {
		log.Fatalf("load control font: %v", err)
	}
	return source
}

func clicked(rect Rect, mouseX, mouseY int) bool {
	return inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) && rect.Contains(mouseX, mouseY)
}

func firstMissionRect() Rect {
	return Rect{X: float64(screenWidth/2 - missionCardWidth/2 - 700), Y: float64(screenHeight/2 - missionCardHeight/2 - 350), W: missionCardWidth, H: missionCardHeight}
}

func secondMissionRect() Rect {
	return Rect{X: float64(screenWidth/2 - missionCardWidth/2 - 700), Y: float64(screenHeight/2 - missionCardHeight/2), W: missionCardWidth, H: missionCardHeight}
}

func thirdMissionRect() Rect {
	return Rect{X: float64(screenWidth/2 - missionCardWidth/2 - 700), Y: float64(screenHeight/2 - missionCardHeight/2 + 350), W: missionCardWidth, H: missionCardHeight}
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

func shortcutHintRect() Rect {
	return Rect{X: 385, Y: 1006, W: 1150, H: 42}
}

func Run() error {
	game, err := newGame()
	if err != nil {
		return err
	}

	ebiten.SetWindowSize(1280, 720)
	ebiten.SetWindowTitle("Paper Quarters")
	ebiten.SetFullscreen(true)
	return ebiten.RunGame(game)
}
