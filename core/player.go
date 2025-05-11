package core

import "errors"

var (
	ErrCardNotInHand = errors.New("card not in hand")
)

type Player struct {
	ID int

	Hand *Hand
}
