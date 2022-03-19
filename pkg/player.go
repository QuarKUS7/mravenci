package src

type Player interface {
	GetBricks() int
	GetWeapons() int
	GetCrystals() int

	GetBuilders() int
	GetSoldiers() int
	GetMages() int

	GetCastle() int
	GetFence() int

	SetCrystals(value int)
	SetBricks(value int)
	SetWeapons(value int)

	SetMages(value int)
	SetBuilders(value int)
	SetSoldiers(value int)

	SetCastle(value int)
	SetFence(value int)

	Attack(value int)

	GetCards() []Card
	ReplaceCard(pos int, newCard Card)
}

type player struct {
	cards []Card

	crystals, bricks, weapons int
	mages, builders, soldiers int
	castle, fence             int
}

func NewPlayer(cardsPackage CardsPackage) Player {
	player := &player{}
	player.crystals = 5
	player.bricks = 5
	player.weapons = 5

	player.mages = 2
	player.builders = 2
	player.soldiers = 2

	player.castle = 30
	player.fence = 10

	for i := 0; i < 10; i++ {
		player.cards = append(player.cards, cardsPackage.GetNextCard())
	}

	return player
}

func (p *player) SetCrystals(value int) {
	if value >= 0 {
		p.crystals = value
	} else {
		p.crystals = 0
	}
}

func (p *player) SetBricks(value int) {
	if value >= 0 {
		p.bricks = value
	} else {
		p.bricks = 0
	}
}

func (p *player) SetWeapons(value int) {
	if value >= 0 {
		p.weapons = value
	} else {
		p.weapons = 0
	}
}

func (p *player) SetMages(value int) {
	if value >= 0 {
		p.mages = value
	} else {
		p.mages = 0
	}
}

func (p *player) SetBuilders(value int) {
	if value >= 0 {
		p.builders = value
	} else {
		p.builders = 0
	}
}

func (p *player) SetSoldiers(value int) {
	if value >= 0 {
		p.soldiers = value
	} else {
		p.soldiers = 0
	}
}

func (p *player) SetCastle(value int) { p.castle = value }

func (p *player) SetFence(value int) {
	if value >= 0 {
		p.fence = value
	} else {
		p.fence = 0
	}

}

func (p *player) Attack(value int) {
	p.fence -= value

	if p.fence < 0 {
		p.castle += p.fence
		p.fence = 0
	}
}

func (p *player) GetCards() []Card { return p.cards }

func (p *player) ReplaceCard(pos int, newCard Card) {
	p.cards[pos] = newCard
}

func (p *player) GetBricks() int   { return p.bricks }
func (p *player) GetWeapons() int  { return p.weapons }
func (p *player) GetCrystals() int { return p.crystals }

func (p *player) GetBuilders() int { return p.builders }
func (p *player) GetSoldiers() int { return p.soldiers }
func (p *player) GetMages() int    { return p.mages }

func (p *player) GetCastle() int { return p.castle }
func (p *player) GetFence() int  { return p.fence }
