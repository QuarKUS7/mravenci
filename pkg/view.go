package src

import "fmt"

type View struct {
	game Game
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

func (view *View) renderPlayer(player Player) string {
	return fmt.Sprintf("Builders: +%d\nBricks 🧱: %d\nSoldiers: +%d\nWeapons 🗡: %d\nMages: +%d\nCrystals 💎: %d\nCastle 🏰: %d\nFence: %d",
		player.GetBuilders(), player.GetBricks(),
		player.GetSoldiers(), player.GetWeapons(),
		player.GetMages(), player.GetCrystals(),
		player.GetCastle(), player.GetFence(),
	)
}

func (view *View) RenderGame() string {
	currentPlayer := view.game.GetCurrentPlayer()
	redPlayer, blackPlayer := view.game.GetRedPlayer(), view.game.GetBlackPlayer()

	board := fmt.Sprintf("Red player: \n\n%s\n\n---------\n\nBlack player:\n\n%s\n\n",
		view.renderPlayer(redPlayer),
		view.renderPlayer(blackPlayer))

	if currentPlayer == blackPlayer {
		board += "Black player's move:"
	} else {
		board += "Red player's move:"
	}

	board += "\n\n----------\n\nCards:\n"
	for i, card := range currentPlayer.GetCards() {
		consumption := card.GetConsumption()

		formatStart := "\033[0m"
		formatEnd := "\033[0m"

		if !view.game.IsValidToPlayNthCard(i) {
			formatStart = "\033[2m"
		}

		board += fmt.Sprintf("%s(%d) %s [%d %s]%s\n",
			formatStart, i, card.Render(), consumption.Amount, view.renderMaterial(consumption.Material), formatEnd)
	}

	return board
}

func (view *View) RenderWinner() {
	currentPlayer := view.game.GetCurrentPlayer()
	redPlayer := view.game.GetRedPlayer()

	var winner string
	if currentPlayer == redPlayer {
		winner = "Black player has won!"
	} else {
		winner = "Red player has won!"
	}

	fmt.Println(winner)
}
