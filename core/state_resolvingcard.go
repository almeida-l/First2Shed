package core

/*
### `StateResolvingCard`
- Centralized logic to apply card effects:
  - Skip
  - Reverse
  - Draw Two
  - Wild (waits for player to choose color)
  - Wild Draw Four
- Also called for normal cards to centralize resolution and turn advancement.
- Transitions to:
  - `StateGameOver` if the last player has emptied their hand.
  - `StatePlayerTurn` otherwise.

*/

type StateResolvingCard struct {
}

func (s *StateResolvingCard) CanHandle(gameCtx *Game, e Event) bool {
	switch e.(type) {
	case CardResolvedEvent:
		return true
	default:
		return false
	}
}

func (s *StateResolvingCard) OnEnter(gameCtx *Game) {
	if gameCtx.lastPlayedCard.HasEffect() {
		ApplyCardEffects(gameCtx, gameCtx.lastPlayedCard)
	}
	gameCtx.Process(CardResolvedEvent{})
}

func (s *StateResolvingCard) Next(gameCtx *Game, e Event) State {
	switch e.(type) {
	case CardResolvedEvent:
		return gameCtx.statePlayerTurn
	default:
		return nil
	}
}
