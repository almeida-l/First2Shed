package core

/*
### `StateGameOver`
- Ends the game.

*/

type StateGameOver struct {
}

func (s *StateGameOver) CanHandle(gameCtx *Game, e Event) bool {
	return false
}

func (s *StateGameOver) OnEnter(gameCtx *Game) {

}

func (s *StateGameOver) Next(gameCtx *Game, e Event) State {
	return nil
}
