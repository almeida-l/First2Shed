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
	case WildCardPlayedEvent:
		return true
	case SetWinner:
		return true
	default:
		return false
	}
}

func (s *StateResolvingCard) OnEnter(gameCtx *Game) {
	if gameCtx.currentPlayer != nil && gameCtx.currentPlayer.Hand.Len() == 0 {
		gameCtx.Process(SetWinner{Player: gameCtx.currentPlayer})
		return
	}

	if gameCtx.lastPlayedCard.Color == CWild {
		gameCtx.Process(WildCardPlayedEvent{})
		return
	}

	if gameCtx.lastPlayedCard.HasEffect() {
		ApplyCardEffects(gameCtx, gameCtx.lastPlayedCard)
	}
	gameCtx.Process(CardResolvedEvent{})
}

func (s *StateResolvingCard) Next(gameCtx *Game, e Event) State {
	switch e.(type) {
	case CardResolvedEvent:
		return gameCtx.statePlayerTurn
	case WildCardPlayedEvent:
		return gameCtx.stateAwaitingColorChoice
	case SetWinner:
		return gameCtx.stateGameOver
	default:
		return nil
	}
}
