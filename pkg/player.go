package pkg

import "fmt"

type Player struct {
	Cards []Card

	Crystals, Bricks, Weapons int
	Mages, Builders, Soldiers int
	Castle, Fence             int
}

func NewPlayer(cardsPackage CardsPackage) *Player {
	player := &Player{}
	player.Crystals = 5
	player.Bricks = 5
	player.Weapons = 5

	player.Mages = 2
	player.Builders = 2
	player.Soldiers = 2

	player.Castle = 30
	player.Fence = 10

	for i := 0; i < 10; i++ {
		player.Cards = append(player.Cards, cardsPackage.GetNextCard())
	}

	return player
}

func (p *Player) String() string {
	return fmt.Sprintf("Builders: +%d\nBricks: %d\nSoldiers: +%d\nWeapons: %d\nMages: +%d\nCrystals: %d\nCastle: %d\nFence: %d",
		p.Builders, p.Bricks,
		p.Soldiers, p.Weapons,
		p.Mages, p.Crystals,
		p.Castle, p.Fence,
	)
}

func (p *Player) SetCrystals(value int) {
	if value >= 0 {
		p.Crystals = value
	} else {
		p.Crystals = 0
	}
}

func (p *Player) SetBricks(value int) {
	if value >= 0 {
		p.Bricks = value
	} else {
		p.Bricks = 0
	}
}

func (p *Player) SetWeapons(value int) {
	if value >= 0 {
		p.Weapons = value
	} else {
		p.Weapons = 0
	}
}

func (p *Player) SetMages(value int) {
	if value >= 0 {
		p.Mages = value
	} else {
		p.Mages = 0
	}
}

func (p *Player) SetBuilders(value int) {
	if value >= 0 {
		p.Builders = value
	} else {
		p.Builders = 0
	}
}

func (p *Player) SetSoldiers(value int) {
	if value >= 0 {
		p.Soldiers = value
	} else {
		p.Soldiers = 0
	}
}

func (p *Player) SetCastle(value int) { p.Castle = value }

func (p *Player) SetFence(value int) {
	if value >= 0 {
		p.Fence = value
	} else {
		p.Fence = 0
	}

}

func (p *Player) Attack(value int) {
	p.Fence -= value

	if p.Fence < 0 {
		p.Castle += p.Fence
		p.Fence = 0
	}
}

func (p *Player) ReplaceCard(pos int, newCard Card) {
	p.Cards[pos] = newCard
}
