package core

/*
### `StateAwaitingColorChoice`
- Waiting for the current player to choice a color
- Transitions back to `StateResolvingCard` when the color is set.

*/

type StateAwaitingColorChoice struct {
}

func (s *StateAwaitingColorChoice) CanHandle(gameCtx *Game, e Event) bool {
	switch ev := e.(type) {
	case SetWildColorCommand:
		return ev.Player.ID == gameCtx.currentPlayer.ID
	default:
		return false
	}
}

func (s *StateAwaitingColorChoice) OnEnter(gameCtx *Game) {

}

func (s *StateAwaitingColorChoice) Next(gameCtx *Game, e Event) State {
	switch ev := e.(type) {
	case SetWildColorCommand:
		if ev.Color == CBlue || ev.Color == CGreen || ev.Color == CRed || ev.Color == CYellow {
			gameCtx.lastPlayedCard.Color = ev.Color
			return gameCtx.stateResolvingCard // returns to the previous state where the cards effects can be handled if needed
		}
		return nil
	default:
		return nil
	}
}
