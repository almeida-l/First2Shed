package core

import "errors"

var (
	ErrCardNotInHand = errors.New("card not in hand")
)

type Player struct {
	ID int

	Game *Game

	Hand *Hand
}

func (p *Player) PlayCard(card Card) error {
	// Validate that the card is in the player's hand
	if !p.Hand.Contains(card) {
		return ErrCardNotInHand
	}

	// Pass the responsibility of actually playing the card to the Game
	return p.Game.PlayCard(p, card)
}
