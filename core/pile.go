package core

import (
	"errors"
	"math/rand"
)

type Pile []Card

var ErrEmptyPile = errors.New("empty pile")

// Element access

// Peek returns the top card of the pile, or false if the pile is empty.
func (p *Pile) Peek() (Card, bool) {
	if len(*p) == 0 {
		return Card{}, false
	}

	return (*p)[len(*p)-1], true
}

// Modifiers

// Push adds a card to the top of the pile.
func (p *Pile) Push(card Card) {
	*p = append(*p, card)
}

// Pop removes and returns the card on the top of the pile.
func (p *Pile) Pop() (Card, error) {
	lastIdx := len(*p) - 1

	if lastIdx < 0 {
		return Card{}, ErrEmptyPile
	}

	card := (*p)[lastIdx]
	*p = (*p)[:lastIdx]
	return card, nil
}

// Len returns the quantity of cards on the pile.
func (p *Pile) Len() int {
	return len(*p)
}

// Shuffles the pile.
func (p *Pile) Shuffle() {
	rand.Shuffle(len(*p), func(i, j int) {
		(*p)[i], (*p)[j] = (*p)[j], (*p)[i]
	})
}
