package core

/*
### `StateSettingInitialCard`
- Flips the first card from the draw pile.
- If the card is invalid (e.g., Wild Draw Four), it is replaced.
- Transitions to `StateResolvingCard`.


*/

type StateSettingInitialCard struct {
}

func (s *StateSettingInitialCard) CanHandle(gameCtx *Game, e Event) bool {
	switch e.(type) {
	case InitialCardSetEvent:
		return true
	default:
		return false
	}
}

func (s *StateSettingInitialCard) OnEnter(gameCtx *Game) {
	var card Card

	for {
		card, _ = gameCtx.drawPile.Pop()
		if !card.IsWild() {
			break
		}

		// its a wild card, so push it back to the draw pile and shuffles again
		gameCtx.drawPile.Push(card)
		gameCtx.drawPile.Shuffle()
	}

	gameCtx.PlayCard(nil, card)
	gameCtx.Process(InitialCardSetEvent{})
}

func (s *StateSettingInitialCard) Next(gameCtx *Game, e Event) State {
	switch e.(type) {
	case InitialCardSetEvent:
		return gameCtx.stateResolvingCard
	default:
		return nil
	}
}
