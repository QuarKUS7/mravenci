package pkg

import (
	"fmt"
	"reflect"
)

type View struct {
	game *Game
}

func (view *View) renderMaterial(material Material) string {
	switch material {
	case Bricks:
		return "Bricks"
	case Weapons:
		return "Weapons"
	case Crystals:
		return "Crystals"
	}

	panic("Unknown material")
}

func (view *View) RenderGame() string {
	currentPlayer := view.game.CurrentPlayer
	redPlayer, blackPlayer := view.game.RedPlayer, view.game.BlackPlayer

	board := fmt.Sprintf("Red player: \n\n%s\n\n---------\n\nBlack player:\n\n%s\n\n",
		redPlayer,
		blackPlayer)

	if view.game.LastPlayed != nil {
		board += "----------\n\nLast played card:\n\n"
		if view.game.Discarted {
			board += fmt.Sprintf("Discarded: %s\n", view.game.LastPlayed)
		} else {
			board += fmt.Sprintf("%s\n", view.game.LastPlayed)
		}
		board += "\n----------\n\n"

	}

	if reflect.DeepEqual(currentPlayer, blackPlayer) {

		board += "Black player's move:"
	} else {
		board += "Red player's move:"
	}

	board += "\n\nCards:\n"
	for i, card := range currentPlayer.Cards {
		consumption := card.GetConsumption()

		formatStart := "\033[0m"
		formatEnd := "\033[0m"

		if !view.game.IsValidToPlayNthCard(i) {
			formatStart = "\033[2m"
		}

		board += fmt.Sprintf("%s(%d) %s [%d %s]%s\n",
			formatStart, i, card, consumption.Amount, view.renderMaterial(consumption.Material), formatEnd)
	}

	return board
}

func (view *View) RenderWinner() {
	currentPlayer := view.game.CurrentPlayer
	redPlayer := view.game.RedPlayer
	var winner string
	if reflect.DeepEqual(currentPlayer, redPlayer) {
		winner = "Black player has won!"
	} else {
		winner = "Red player has won!"
	}

	fmt.Println(winner)
}
