package core

import (
	"testing"
)

func TestCard_IsWild(t *testing.T) {
	tests := []struct {
		card     Card
		expected bool
	}{
		{Card{Color: CWild, Value: VWild}, true},
		{Card{Color: CBlue, Value: VOne}, false},
		{Card{Color: CRed, Value: VWildDrawFour}, true},
	}

	for _, test := range tests {
		t.Run(test.card.String(), func(t *testing.T) {
			if got := test.card.IsWild(); got != test.expected {
				t.Errorf("Expected %v, got %v", test.expected, got)
			}
		})
	}
}

func TestCard_CanPlayOn(t *testing.T) {
	tests := []struct {
		card       Card
		bottomCard Card
		canPlay    bool
	}{
		{Card{Color: CBlue, Value: VOne}, Card{Color: CBlue, Value: VTwo}, true},
		{Card{Color: CRed, Value: VOne}, Card{Color: CGreen, Value: VOne}, true},
		{Card{Color: CYellow, Value: VSkip}, Card{Color: CRed, Value: VOne}, false},
		{Card{Color: CWild, Value: VWild}, Card{Color: CGreen, Value: VTwo}, true},
	}

	for _, test := range tests {
		t.Run(test.card.String(), func(t *testing.T) {
			if got := test.card.CanPlayOn(test.bottomCard); got != test.canPlay {
				t.Errorf("Expected CanPlayOn to be %v, got %v", test.canPlay, got)
			}
		})
	}
}

func TestCard_String(t *testing.T) {
	tests := []struct {
		card     Card
		expected string
	}{
		{Card{Color: CRed, Value: VOne}, "R1"},
		{Card{Color: CGreen, Value: VDrawTwo}, "GT"},
		{Card{Color: CWild, Value: VWildDrawFour}, "WF"},
	}

	for _, test := range tests {
		t.Run(test.card.String(), func(t *testing.T) {
			if got := test.card.String(); got != test.expected {
				t.Errorf("Expected %v, got %v", test.expected, got)
			}
		})
	}
}
