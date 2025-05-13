package core

import "log"

func (g *Game) DebugGetCurrentPlayer() *Player {
	return g.currentPlayer
}

func (g *Game) DebugGetLastPlayedCard() Card {
	return g.lastPlayedCard
}

type Game struct {
	players []*Player

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
	g.statePlayerTurn = &StatePlayerTurn{}

	g.state = g.stateLobby

	g.turnDirection = 1
	g.currentPlayerIdx = -1 // initializes to -1 so when the game start the NextTurn() method will be called and the first players will be set to idx 0

	g.isLobbyOpen = true
}

func (g *Game) Process(e Event) {
	log.Printf("[PROCESS EVENT] Event %T: %+v", e, e)

	if _, ok := e.(GlobalEvent); ok {
		g.handleGlobalEvent(e)
		return
	}

	if g.state.CanHandle(g, e) {
		log.Printf("[PROCESS EVENT] State %T can handle the event %T", g.state, e)
		if state := g.state.Next(g, e); state != nil {
			log.Printf("[PROCESS EVENT] Transitioning to state %T", state)
			g.state = state
			g.state.OnEnter(g)
		}
	} else {
		log.Printf("[PROCESS EVENT] State %T cannot handle the event %T", g.state, e)
	}
}

func (g *Game) PlayCard(player *Player, card Card) error {
	g.discardPile.Push(card)
	g.lastPlayedCard = card
	if player != nil { // when player is nil, it's the initial card
		player.Hand.Remove(card)
	}
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

func (g *Game) GetPlayerFromID(playerID int) *Player {
	for idx, p := range g.players {
		if p.ID == playerID {
			return g.players[idx]
		}
	}
	return nil
}

func (g *Game) PeekNextPlayer() *Player {
	nextIdx := g.currentPlayerIdx + g.turnDirection
	switch {
	case nextIdx >= len(g.players):
		nextIdx = 0
	case nextIdx < 0:
		nextIdx = len(g.players) - 1
	}

	return g.players[nextIdx]
}

func (g *Game) PopCardFromDrawPile() Card {
	card, err := g.drawPile.Pop()
	if err == ErrEmptyPile {
		g.ResetDrawPile()
		card, _ = g.drawPile.Pop()
	}

	return card
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

	if player := g.GetPlayerFromID(playerJoinCommand.ID); player != nil {
		// TODO: raise an event that notifies the player that he is already in the game
		log.Printf("player ID %d already in the game", playerJoinCommand.ID)
		return
	}
	newPlayer := &Player{
		ID:   playerJoinCommand.ID,
		Hand: &Hand{},
	}

	g.players = append(g.players, newPlayer)
	// TODO: raise an event to annouce that a new player joined
	log.Printf("player ID %d joined", playerJoinCommand.ID)
}
