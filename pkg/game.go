package pkg

type Game struct {
	BlackPlayer, RedPlayer        *Player
	CurrentPlayer, OpponentPlayer *Player
	cardsPackage                  CardsPackage
	round                         int
	LastPlayed                    Card
	Discarted                     bool
}

func NewGame() *Game {
	game := &Game{}
	game.cardsPackage = NewCardsPackage()
	game.BlackPlayer = NewPlayer(game.cardsPackage)
	game.RedPlayer = NewPlayer(game.cardsPackage)
	game.CurrentPlayer = game.BlackPlayer
	game.OpponentPlayer = game.RedPlayer
	game.round = 0
	return game
}

func consumeMaterial(player *Player, consumption CardConsumption) {
	switch consumption.Material {
	case Bricks:
		player.SetBricks(player.Bricks - consumption.Amount)
	case Weapons:
		player.SetWeapons(player.Weapons - consumption.Amount)
	case Crystals:
		player.SetCrystals(player.Crystals - consumption.Amount)
	}
	return
}

func (game *Game) IsValidToPlayNthCard(card_index int) bool {
	player := game.CurrentPlayer
	card := player.Cards[card_index]
	consumption := card.GetConsumption()

	if consumption.Material == Bricks && player.Bricks < consumption.Amount {
		return false
	}

	if consumption.Material == Weapons && player.Weapons < consumption.Amount {
		return false
	}

	if consumption.Material == Crystals && player.Crystals < consumption.Amount {
		return false
	}

	return true
}

func (game *Game) PlayNthCard(card_index int) {
	card := game.CurrentPlayer.Cards[card_index]
	card.Play(game.CurrentPlayer, game.OpponentPlayer)
	consumeMaterial(game.CurrentPlayer, card.GetConsumption())
	game.CurrentPlayer.ReplaceCard(card_index, game.cardsPackage.GetNextCard())
	game.LastPlayed = card
	game.Discarted = false
}

func (game *Game) DiscardNthCard(card_index int) {
	game.LastPlayed = game.CurrentPlayer.Cards[card_index]
	game.Discarted = true
	game.CurrentPlayer.ReplaceCard(card_index, game.cardsPackage.GetNextCard())

}

func (game *Game) StartRound() {
	me := game.CurrentPlayer

	if game.round > 0 {
		me.SetBricks(me.Bricks + me.Builders)
		me.SetWeapons(me.Weapons + me.Soldiers)
		me.SetCrystals(me.Crystals + me.Mages)
	}
}

func (game *Game) EndRound() {
	game.CurrentPlayer, game.OpponentPlayer = game.OpponentPlayer, game.CurrentPlayer
	game.round++
}

func (game *Game) HasWon() bool {
	return game.OpponentPlayer.Castle <= 0 || game.CurrentPlayer.Castle <= 0 || game.OpponentPlayer.Castle >= 100 || game.CurrentPlayer.Castle >= 100
}
