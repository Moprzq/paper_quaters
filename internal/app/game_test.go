package app

import "testing"

func TestNextTurnCountsOnlyNewTurns(t *testing.T) {
	g := &Game{turn: 1, maxCountedTurn: 1}

	g.nextTurn()
	if g.housesBuilt != 1 {
		t.Fatalf("housesBuilt after first nextTurn = %d, want 1", g.housesBuilt)
	}

	g.previousTurn()
	g.nextTurn()
	if g.housesBuilt != 1 {
		t.Fatalf("housesBuilt after returning to counted turn = %d, want 1", g.housesBuilt)
	}

	g.nextTurn()
	if g.housesBuilt != 2 {
		t.Fatalf("housesBuilt after new turn = %d, want 2", g.housesBuilt)
	}
}

func TestShuffleKeepsCounterAndStartsNewCountedRun(t *testing.T) {
	g := &Game{deck: testDeck(), turn: 3, housesBuilt: 2, maxCountedTurn: 3}

	g.shuffleDeck()
	if g.housesBuilt != 2 {
		t.Fatalf("housesBuilt after shuffle = %d, want 2", g.housesBuilt)
	}
	if g.turn != 1 {
		t.Fatalf("turn after shuffle = %d, want 1", g.turn)
	}

	g.nextTurn()
	if g.housesBuilt != 3 {
		t.Fatalf("housesBuilt after nextTurn on shuffled deck = %d, want 3", g.housesBuilt)
	}
}

func TestRestartResetsCounter(t *testing.T) {
	g := &Game{
		cache:          newImageCache(),
		deck:           testDeck(),
		turn:           3,
		housesBuilt:    2,
		maxCountedTurn: 3,
	}

	if err := g.restart(); err != nil {
		t.Fatalf("restart() error = %v", err)
	}
	if g.housesBuilt != 0 {
		t.Fatalf("housesBuilt after restart = %d, want 0", g.housesBuilt)
	}
	if g.maxCountedTurn != 1 {
		t.Fatalf("maxCountedTurn after restart = %d, want 1", g.maxCountedTurn)
	}
}

func testDeck() Deck {
	return Deck{Cards: make([]Card, cardsInDeck)}
}
