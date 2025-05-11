package core

import (
	"testing"
)

func TestHand_AddAndContains(t *testing.T) {
	hand := Hand{}
	card := Card{Color: 1, Value: 3}
	hand.Add(card)

	if !hand.Contains(card) {
		t.Errorf("Expected hand to contain card %+v", card)
	}

	if hand.Len() != 1 {
		t.Errorf("Expected hand length 1, got %d", hand.Len())
	}
}

func TestHand_Remove(t *testing.T) {
	hand := Hand{
		{Color: 1, Value: 3},
		{Color: 2, Value: 5},
	}
	hand.Remove(Card{Color: 1, Value: 3})

	if hand.Len() != 1 {
		t.Errorf("Expected hand length 1 after remove, got %d", hand.Len())
	}
	if hand.Contains(Card{Color: 1, Value: 3}) {
		t.Errorf("Card should have been removed")
	}
}

func TestHand_Sort(t *testing.T) {
	hand := Hand{
		{Color: 2, Value: 7},
		{Color: 1, Value: 4},
		{Color: 3, Value: 0},
		{Color: 2, Value: 3},
		{Color: 1, Value: 2},
	}

	hand.Sort() // Sorts the hand

	expected := Hand{
		{Color: 1, Value: 2},
		{Color: 1, Value: 4},
		{Color: 2, Value: 3},
		{Color: 2, Value: 7},
		{Color: 3, Value: 0},
	}
	for i := range hand {
		if hand[i] != expected[i] {
			t.Errorf("Expected card %v at index %d, got %v", expected[i], i, hand[i])
		}
	}
}
