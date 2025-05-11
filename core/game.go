package core

type Game struct {
	players    []*Player
	numPlayers int

	drawPile       Pile
	discardPile    Pile
	lastPlayedCard Card

	state                   State
	stateLobby              *StateLobby
	stateDealing            *StateDealing
	stateSettingInitialCard *StateSettingInitialCard
	stateResolvingCard      *StateResolvingCard
}

func (g *Game) Init() {
	g.stateLobby = &StateLobby{}
	g.stateDealing = &StateDealing{}
	g.stateSettingInitialCard = &StateSettingInitialCard{}
	g.stateResolvingCard = &StateResolvingCard{}

	g.state = g.stateLobby
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
