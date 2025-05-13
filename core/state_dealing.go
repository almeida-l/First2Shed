package core

import "first2shed/core/rule"

type StateDealing struct{}

/* ### `StateDealing`
- Initializes the draw pile.
- Deals 7 cards to each player.
- Transitions to `StateSettingInitialCard`.
*/

func (s *StateDealing) OnEnter(gameCtx *Game) {
	gameCtx.drawPile = generateFullDeck()
	gameCtx.drawPile.Shuffle()

	for i := 0; i < rule.InitialHandSize; i++ {
		for _, p := range gameCtx.players {
			card, _ := gameCtx.drawPile.Pop()
			p.Hand.Add(card)
		}
	}

	gameCtx.Process(DealingFinishedEvent{})
}

func (s *StateDealing) CanHandle(gameCtx *Game, e Event) bool {
	switch e.(type) {
	case DealingFinishedEvent:
		return true
	default:
		return false
	}
}

func (s *StateDealing) Next(gameCtx *Game, e Event) State {
	switch e.(type) {
	case DealingFinishedEvent:
		return gameCtx.stateSettingInitialCard
	}
	return nil
}

func generateFullDeck() Pile {
	var drawPile Pile

	colors := []Color{CRed, CGreen, CBlue, CYellow}

	// Number cards
	for _, color := range colors {
		// One zero per color
		c := Card{Color: color, Value: VZero}
		drawPile.Push(c)

		// Two of each 1–9 per color
		for v := VOne; v <= VNine; v++ {
			for i := 0; i < 2; i++ {
				c := Card{Color: color, Value: v}
				drawPile.Push(c)
			}
		}
	}

	// Action cards (Skip, Reverse, Draw Two) - 2 per color
	for _, color := range colors {
		for _, v := range []Value{VSkip, VReverse, VDrawTwo} {
			for i := 0; i < 2; i++ {
				c := Card{Color: color, Value: v}
				drawPile.Push(c)
			}
		}
	}

	// Wild cards
	for i := 0; i < 4; i++ {
		wild := Card{CWild, VWild}
		drawFour := Card{CWild, VWildDrawFour}
		drawPile.Push(wild)
		drawPile.Push(drawFour)
	}

	return drawPile
}
