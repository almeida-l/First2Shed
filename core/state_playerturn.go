package core

/* `StatePlayerTurn`
- Waits for current player to:
  - Play a card (must match color or value, or be wild)
  - Or draw and play it or pass.
- Transitions to `StateResolvingCard` after a play.
- Transitions back to `SatePlayerTurn` when player pass (after changing turn)*/

type StatePlayerTurn struct {
	hasDrew bool
}

func (s *StatePlayerTurn) CanHandle(gameCtx *Game, e Event) bool {
	switch event := e.(type) {
	case PlayCardCommand:
		return canPlayCard(gameCtx, event)
	case DrawCardCommand:
		return canDrawCard(gameCtx, event)
	case PassCommand:
		return canPass(gameCtx, event)
	default:
		return false
	}
}

func (s *StatePlayerTurn) OnEnter(gameCtx *Game) {
	s.hasDrew = false
}

func (s *StatePlayerTurn) Next(gameCtx *Game, e Event) State {
	switch event := e.(type) {
	case PlayCardCommand:
		event.player.PlayCard(event.card)
		return gameCtx.stateResolvingCard
	case DrawCardCommand:
		card, err := gameCtx.drawPile.Pop()
		if err == ErrEmptyPile {
			// that is kinda sus but i think its ok for now
			gameCtx.ResetDrawPile()
			card, _ = gameCtx.drawPile.Pop()
		}
		event.player.Hand.Add(card)
		s.hasDrew = true
		return nil
	case PassCommand:
		gameCtx.NextTurn()
		return gameCtx.statePlayerTurn
	}

	return nil
}

func canPlayCard(gameCtx *Game, playCardCommand PlayCardCommand) bool {
	if playCardCommand.player != gameCtx.currentPlayer {
		return false
	}
	if !playCardCommand.card.CanPlayOn(gameCtx.lastPlayedCard) {
		return false
	}
	if !playCardCommand.player.Hand.Contains(playCardCommand.card) {
		return false
	}

	return true
}

func canDrawCard(gameCtx *Game, drawCardCommand DrawCardCommand) bool {
	if drawCardCommand.player != gameCtx.currentPlayer {
		return false
	}

	return !gameCtx.statePlayerTurn.hasDrew
}

func canPass(gameCtx *Game, passCommand PassCommand) bool {
	if passCommand.player != gameCtx.currentPlayer {
		return false
	}

	return gameCtx.statePlayerTurn.hasDrew
}
