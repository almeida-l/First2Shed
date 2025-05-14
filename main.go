package main

import (
	"first2shed/core"
	"fmt"
)

func main() {
	game := core.Game{}
	game.Init()

	game.Process(core.PlayerJoinCommand{ID: 0})
	game.Process(core.PlayerJoinCommand{ID: 1})

	game.Process(core.StartGameCommand{})

	// print the initial card
	fmt.Println("Initial card:", game.DebugGetLastPlayedCard())

	// loop to choose the next command
	for {
		fmt.Printf("Last played card: %s\n", game.DebugGetLastPlayedCard())
		helperShowPlayerHand(&game)
		fmt.Println("Input a card code or a command:")
		fmt.Println("1. Draw card")
		fmt.Println("2. Pass")
		fmt.Println("3. Choose Color")

		var command string
		fmt.Scanln(&command)

		switch command[0:1] {
		case "1":
			helperDrawCardAsPlayer(&game)
		case "2":
			helperPassAsPlayer(&game)
		case "3":
			helperChooseColor(&game)
		default:
			helperPlayCardAsPlayer(&game, command)
		}
	}
}

func helperPlayCardAsPlayer(gameCtx *core.Game, cardCode string) {
	playerHandler := gameCtx.DebugGetCurrentPlayer()

	card := core.Card{}
	card.FromString(cardCode)

	gameCtx.Process(core.PlayCardCommand{Card: card, Player: playerHandler})
}

func helperDrawCardAsPlayer(gameCtx *core.Game) {
	gameCtx.Process(core.DrawCardCommand{Player: gameCtx.DebugGetCurrentPlayer()})
}

func helperPassAsPlayer(gameCtx *core.Game) {
	gameCtx.Process(core.PassCommand{Player: gameCtx.DebugGetCurrentPlayer()})
}

func helperShowPlayerHand(gameCtx *core.Game) {
	playerHandler := gameCtx.DebugGetCurrentPlayer()
	fmt.Printf("Player ID %d hand: ", gameCtx.DebugGetCurrentPlayer().ID)
	for _, card := range *playerHandler.Hand {
		fmt.Printf("%s ", card.String())
	}
	fmt.Println()
}

func helperChooseColor(gameCtx *core.Game) {
	var colorStr string
	fmt.Printf("Enter color letter: ")
	fmt.Scanln(&colorStr)

	// hack to parse the color
	colorStr = colorStr[0:1] + "0"
	cardTmp := core.Card{}
	err := cardTmp.FromString(colorStr)
	if err != nil {
		fmt.Println(err)
		return
	}

	event := core.SetWildColorCommand{
		Player: gameCtx.DebugGetCurrentPlayer(),
		Color:  cardTmp.Color,
	}

	gameCtx.Process(event)
}
