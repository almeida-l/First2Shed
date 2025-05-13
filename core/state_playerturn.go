package core

import "log"

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
	gameCtx.NextTurn()

	s.hasDrew = false
	// TODO: raise an event that notifies that is the players turn
	log.Printf("PlayerID %d it's yours turn!", gameCtx.currentPlayer.ID)
}

func (s *StatePlayerTurn) Next(gameCtx *Game, e Event) State {
	switch event := e.(type) {
	case PlayCardCommand:
		gameCtx.PlayCard(event.Player, event.Card)
		return gameCtx.stateResolvingCard
	case DrawCardCommand:
		card, err := gameCtx.drawPile.Pop()
		if err == ErrEmptyPile {
			// that is kinda sus but i think its ok for now
			gameCtx.ResetDrawPile()
			card, _ = gameCtx.drawPile.Pop()
		}
		event.Player.Hand.Add(card)
		s.hasDrew = true
		return nil
	case PassCommand:
		gameCtx.NextTurn()
		return gameCtx.statePlayerTurn
	}

	return nil
}

func canPlayCard(gameCtx *Game, playCardCommand PlayCardCommand) bool {
	if playCardCommand.Player != gameCtx.currentPlayer {
		log.Printf("PlayerID %d is not the player in turn", playCardCommand.Player.ID)
		return false
	}
	if !playCardCommand.Card.CanPlayOn(gameCtx.lastPlayedCard) {
		log.Printf("The card %s cannot be played on card %s", playCardCommand.Card, gameCtx.lastPlayedCard)
		return false
	}
	if !playCardCommand.Player.Hand.Contains(playCardCommand.Card) {
		log.Printf("PlayerID %d do not have the card %s in his hand", playCardCommand.Player.ID, playCardCommand.Card)
		return false
	}

	return true
}

func canDrawCard(gameCtx *Game, drawCardCommand DrawCardCommand) bool {
	if drawCardCommand.Player != gameCtx.currentPlayer {
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
