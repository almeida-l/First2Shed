package core

import "log"

type Game struct {
	players    []*Player
	numPlayers int

	currentPlayer    *Player
	currentPlayerIdx int

	turnDirection int

	drawPile       Pile
	discardPile    Pile
	lastPlayedCard Card

	isLobbyOpen bool

	state                   State
	stateLobby              *StateLobby
	stateDealing            *StateDealing
	stateSettingInitialCard *StateSettingInitialCard
	stateResolvingCard      *StateResolvingCard
	statePlayerTurn         *StatePlayerTurn
}

func (g *Game) Init() {
	g.stateLobby = &StateLobby{}
	g.stateDealing = &StateDealing{}
	g.stateSettingInitialCard = &StateSettingInitialCard{}
	g.stateResolvingCard = &StateResolvingCard{}

	g.state = g.stateLobby

	g.turnDirection = 1

	g.isLobbyOpen = true
}

func (g *Game) Process(e Event) {
	log.Printf("[PROCESS EVENT] Event %T: %+v", e, e)

	if _, ok := e.(GlobalEvent); ok {
		g.handleGlobalEvent(e)
		return
	}

	if g.state.CanHandle(g, e) {
		if state := g.state.Next(g, e); state != nil {
			g.state.OnEnter(g)
		}
	}
}

func (g *Game) PlayCard(player *Player, card Card) error {
	g.discardPile.Push(card)
	g.lastPlayedCard = card
	return nil
}

func (g *Game) ResetDrawPile() {
	g.drawPile = append(g.drawPile, g.discardPile[:g.discardPile.Len()-1]...)
	g.discardPile = g.discardPile[:g.discardPile.Len()-1]
	g.drawPile.Shuffle()
}

func (g *Game) NextTurn() {
	g.currentPlayerIdx += g.turnDirection
	switch {
	case g.currentPlayerIdx >= len(g.players):
		g.currentPlayerIdx = 0
	case g.currentPlayerIdx < 0:
		g.currentPlayerIdx = len(g.players) - 1
	}

	g.currentPlayer = g.players[g.currentPlayerIdx]
}

func (g *Game) handleGlobalEvent(e Event) {
	switch ev := e.(type) {
	case PlayerJoinCommand:
		g.handlePlayerJoin(ev)
	}
}

func (g *Game) handlePlayerJoin(playerJoinCommand PlayerJoinCommand) {
	if !g.isLobbyOpen {
		// TODO: raise an event that notifies the player that the lobby is closed
		log.Printf("player ID %d can't join, lobby is closed", playerJoinCommand.ID)
		return
	}

	for _, p := range g.players {
		if p.ID == playerJoinCommand.ID {
			// TODO: raise an event that notifies the player that he is already in the game
			log.Printf("player ID %d already in the game", playerJoinCommand.ID)
			return
		}
	}

	g.players = append(g.players, &Player{ID: playerJoinCommand.ID})
	// TODO: raise an event to annouce that a new player joined
	log.Printf("player ID %d joined", playerJoinCommand.ID)
}
