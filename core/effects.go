package core

func ApplyCardEffects(gameCtx *Game, card Card) {
	switch card.Value {
	case VSkip:
		gameCtx.NextTurn()
	case VReverse:
		gameCtx.turnDirection *= -1
	case VDrawTwo:
		makeNextPlayerDrawNCards(gameCtx, 2)
		gameCtx.NextTurn()
	case VWildDrawFour:
		makeNextPlayerDrawNCards(gameCtx, 4)
	}
}

func makeNextPlayerDrawNCards(gameCtx *Game, amount int) {
	target := gameCtx.PeekNextPlayer()
	for i := 0; i < amount; i++ {
		card := gameCtx.PopCardFromDrawPile()
		target.Hand.Add(card)
	}
}
