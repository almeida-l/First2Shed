package core

import (
	"io"
	"log"
	"testing"
)

func TestFirstPlayerInTurnSkip(t *testing.T) {
	log.SetOutput(io.Discard)
	game := Game{}
	game.Init()

	game.Process(PlayerJoinCommand{ID: 0})
	game.Process(PlayerJoinCommand{ID: 1})
	game.lastPlayedCard = Card{Color: CBlue, Value: VSkip}
	game.state = game.stateResolvingCard
	game.state.OnEnter(&game)

	expected := 1
	got := game.currentPlayerIdx

	if expected != got {
		t.Errorf("Expected the player in turn to be ID %d, got %d", expected, got)
	}
}

func TestFirstPlayerInTurnReverse(t *testing.T) {
	log.SetOutput(io.Discard)
	game := Game{}
	game.Init()

	game.Process(PlayerJoinCommand{ID: 0})
	game.Process(PlayerJoinCommand{ID: 1})
	game.lastPlayedCard = Card{Color: CBlue, Value: VReverse}
	game.state = game.stateResolvingCard
	game.state.OnEnter(&game)

	expected := 1
	got := game.currentPlayerIdx

	if expected != got {
		t.Errorf("Expected the player in turn to be ID %d, got %d", expected, got)
	}
}

func TestFirstPlayerInTurnNormal(t *testing.T) {
	log.SetOutput(io.Discard)

	game := Game{}
	game.Init()

	game.Process(PlayerJoinCommand{ID: 0})
	game.Process(PlayerJoinCommand{ID: 1})
	game.lastPlayedCard = Card{Color: CBlue, Value: VZero}
	game.state = game.stateResolvingCard
	game.state.OnEnter(&game)

	expected := 0
	got := game.currentPlayerIdx

	if expected != got {
		t.Errorf("Expected the player in turn to be ID %d, got %d", expected, got)
	}
}
