package core

/*
### `StateLobby`
- Waiting for players to join.
- Transitions to `StateDealing` when the game starts.

So, on enter we need to annouce that the lobby is open (add a notification)

on handle we need to handle the start command
	if we have more than 1 player in the lobby, we can transite to StateDealing

*/

type StateLobby struct {
}

func (s *StateLobby) CanHandle(gameCtx *Game, e Event) bool {
	switch e.(type) {
	case StartGameCommand:
		return len(gameCtx.players) > 1
	}
	return false
}

func (s *StateLobby) OnEnter(gameCtx *Game) {}

func (s *StateLobby) Next(gameCtx *Game, e Event) State {
	switch e.(type) {
	case StartGameCommand:
		return gameCtx.stateDealing
	default:
		return nil
	}
}
