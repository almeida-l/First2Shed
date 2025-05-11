package core

import "errors"

var (
	ErrCardNotInHand = errors.New("card not in hand")
)

type Player struct {
	ID int

	game *Game

	Hand *Hand
}
