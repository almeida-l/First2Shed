package core

type Game struct {
	players    []*Player
	numPlayers int

	currentPlayer    *Player
	currentPlayerIdx int

	turnDirection int

	drawPile       Pile
	discardPile    Pile
	lastPlayedCard Card

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
}

func (g *Game) Process(e Event) {
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
