package core

import "sort"

type Hand []Card

// Returns if the hand contains a given card.
func (h *Hand) Contains(card Card) bool {
	for _, c := range *h {
		if c.Color == card.Color && c.Value == card.Value {
			return true
		}
	}
	return false
}

// Adds a card to the hand and then sorts the hand.
func (h *Hand) Add(card Card) {
	(*h) = append((*h), card)
	(*h).Sort()
}

// Try to remove a card from the hand and then sorts the hand if the card got removed.
func (h *Hand) Remove(card Card) {
	for idx, c := range *h {
		if c.Color == card.Color && c.Value == card.Value {
			*h = append((*h)[:idx], (*h)[idx+1:]...)
			(*h).Sort()
			return
		}
	}
}

// Returns the amount of cards in the hand.
func (h *Hand) Len() int {
	return len(*h)
}

// Sorts the hand cards by color and value.
func (h *Hand) Sort() {
	sort.Slice((*h), func(i, j int) bool {
		if (*h)[i].Color == (*h)[j].Color {
			return (*h)[i].Value < (*h)[j].Value
		}
		return (*h)[i].Color < (*h)[j].Color
	})
}
