package pkg

import (
	"math/rand"
)

type Card interface {
	Play(me *Player, opponent *Player)
	String() string
	GetConsumption() CardConsumption
}

type Material int

const (
	Bricks Material = iota
	Weapons
	Crystals
)

type CardConsumption struct {
	Amount   int
	Material Material
}

type CardsPackage interface {
	GetNextCard() Card
}

type randomCardsPackage struct {
	possibleCards []Card
}

func NewCardsPackage() CardsPackage {
	cardsPackage := &randomCardsPackage{}
	cardsPackage.possibleCards = append(cardsPackage.possibleCards,
		&BansheeCard{},
		&DragonCard{},
		&AttackCard{},
		&SwatCard{},
		&PlatoonCard{},
		&RaiderCard{},
		&ArcherCard{},

		&FenceCard{},
		&DefenceCard{},
		&WallCard{},

		&BabylonCard{},
		&PixiesCard{},
		&FortCard{},
		&WainCard{},
		&ReservesCard{},
		&TowerCard{},
		&BaseCard{},

		&DestroyBricksCard{},
		&DestroyWeaponsCard{},
		&DestroyCrystalsCard{},

		&ConjureBricksCard{},
		&ConjureCrystalsCard{},
		&ConjureWeaponsCard{},

		&SorcererCard{},
		&RecruitingCard{},
		&SchoolCard{},

		&CurseCard{},
		&SaboteurCard{},
	)
	return cardsPackage
}

func (p *randomCardsPackage) GetNextCard() Card {
	return p.possibleCards[rand.Intn(len(p.possibleCards))]
}

// Attack Cards

type BansheeCard struct{}

func (_ *BansheeCard) Play(me *Player, opponent *Player) { opponent.Attack(32) }
func (_ *BansheeCard) String() string                    { return "Banshee - attack +32" }
func (_ *BansheeCard) GetConsumption() CardConsumption   { return CardConsumption{28, Weapons} }

type DragonCard struct{}

func (_ *DragonCard) Play(me *Player, opponent *Player) { opponent.Attack(25) }
func (_ *DragonCard) String() string                    { return "Dragon - attack +25" }
func (_ *DragonCard) GetConsumption() CardConsumption   { return CardConsumption{21, Crystals} }

type AttackCard struct{}

func (_ *AttackCard) Play(me *Player, opponent *Player) { opponent.Attack(12) }
func (_ *AttackCard) String() string                    { return "Attack - attack +12" }
func (_ *AttackCard) GetConsumption() CardConsumption   { return CardConsumption{10, Weapons} }

type SwatCard struct{}

func (_ *SwatCard) Play(me *Player, opponent *Player) { opponent.SetCastle(opponent.Castle - 10) }
func (_ *SwatCard) String() string                    { return "SWAT - Castle of your enemy -10" }
func (_ *SwatCard) GetConsumption() CardConsumption   { return CardConsumption{18, Weapons} }

type PlatoonCard struct{}

func (_ *PlatoonCard) Play(me *Player, opponent *Player) { opponent.Attack(6) }
func (_ *PlatoonCard) String() string                    { return "Platoon - attack +6" }
func (_ *PlatoonCard) GetConsumption() CardConsumption   { return CardConsumption{4, Weapons} }

type RaiderCard struct{}

func (_ *RaiderCard) Play(me *Player, opponent *Player) { opponent.Attack(4) }
func (_ *RaiderCard) String() string                    { return "Raider - attack +4" }
func (_ *RaiderCard) GetConsumption() CardConsumption   { return CardConsumption{2, Weapons} }

type ArcherCard struct{}

func (_ *ArcherCard) Play(me *Player, opponent *Player) { opponent.Attack(2) }
func (_ *ArcherCard) String() string                    { return "Archer - attack +2" }
func (_ *ArcherCard) GetConsumption() CardConsumption   { return CardConsumption{1, Weapons} }

// Fence Card

type FenceCard struct{}

func (_ *FenceCard) Play(me *Player, opponent *Player) { me.SetFence(me.Fence + 22) }
func (_ *FenceCard) String() string                    { return "Fence - Fence +22" }
func (_ *FenceCard) GetConsumption() CardConsumption   { return CardConsumption{12, Bricks} }

type DefenceCard struct{}

func (_ *DefenceCard) Play(me *Player, opponent *Player) { me.SetFence(me.Fence + 6) }
func (_ *DefenceCard) String() string                    { return "Fence - Fence +6" }
func (_ *DefenceCard) GetConsumption() CardConsumption   { return CardConsumption{3, Bricks} }

type WallCard struct{}

func (_ *WallCard) Play(me *Player, opponent *Player) { me.SetFence(me.Fence + 3) }
func (_ *WallCard) String() string                    { return "Fence - Fence +3" }
func (_ *WallCard) GetConsumption() CardConsumption   { return CardConsumption{1, Bricks} }

// Castle Cards

type BabylonCard struct{}

func (_ *BabylonCard) Play(me *Player, opponent *Player) { me.SetCastle(me.Castle + 32) }
func (_ *BabylonCard) String() string                    { return "Babylon - Castle +32" }
func (_ *BabylonCard) GetConsumption() CardConsumption   { return CardConsumption{39, Bricks} }

type PixiesCard struct{}

func (_ *PixiesCard) Play(me *Player, opponent *Player) { me.SetCastle(me.Castle + 22) }
func (_ *PixiesCard) String() string                    { return "Pixies - Castle +22" }
func (_ *PixiesCard) GetConsumption() CardConsumption   { return CardConsumption{22, Crystals} }

type FortCard struct{}

func (_ *FortCard) Play(me *Player, opponent *Player) { me.SetCastle(me.Castle + 20) }
func (_ *FortCard) String() string                    { return "Fort - Castle +20" }
func (_ *FortCard) GetConsumption() CardConsumption   { return CardConsumption{18, Bricks} }

type WainCard struct{}

func (_ *WainCard) Play(me *Player, opponent *Player) {
	me.SetCastle(me.Castle + 8)
	opponent.SetCastle(opponent.Castle - 4)
}
func (_ *WainCard) String() string                  { return "Wain - Castle +8, Castle of your enemy -4" }
func (_ *WainCard) GetConsumption() CardConsumption { return CardConsumption{10, Bricks} }

type ReservesCard struct{}

func (_ *ReservesCard) Play(me *Player, opponent *Player) {
	me.SetCastle(me.Castle + 8)
	me.SetFence(me.Fence - 4)
}
func (_ *ReservesCard) String() string                  { return "Reserves - Castle +8, Fence -4" }
func (_ *ReservesCard) GetConsumption() CardConsumption { return CardConsumption{3, Bricks} }

type TowerCard struct{}

func (_ *TowerCard) Play(me *Player, opponent *Player) { me.SetCastle(me.Castle + 5) }
func (_ *TowerCard) String() string                    { return "Tower - Castle +5" }
func (_ *TowerCard) GetConsumption() CardConsumption   { return CardConsumption{5, Bricks} }

type BaseCard struct{}

func (_ *BaseCard) Play(me *Player, opponent *Player) { me.SetCastle(me.Castle + 2) }
func (_ *BaseCard) String() string                    { return "Base - Castle +2" }
func (_ *BaseCard) GetConsumption() CardConsumption   { return CardConsumption{1, Bricks} }

// Destroy Cards

type DestroyWeaponsCard struct{}

func (_ *DestroyWeaponsCard) Play(me *Player, opponent *Player) {
	opponent.SetWeapons(opponent.Weapons - 8)
}
func (_ *DestroyWeaponsCard) String() string                  { return "Destroy Weapons - Weapons of your enemy -8" }
func (_ *DestroyWeaponsCard) GetConsumption() CardConsumption { return CardConsumption{4, Crystals} }

type DestroyBricksCard struct{}

func (_ *DestroyBricksCard) Play(me *Player, opponent *Player) {
	opponent.SetBricks(opponent.Bricks - 8)
}
func (_ *DestroyBricksCard) String() string                  { return "Destroy Bricks - Bricks of your enemy -8" }
func (_ *DestroyBricksCard) GetConsumption() CardConsumption { return CardConsumption{4, Crystals} }

type DestroyCrystalsCard struct{}

func (_ *DestroyCrystalsCard) Play(me *Player, opponent *Player) {
	opponent.SetCrystals(opponent.Crystals - 8)
}
func (_ *DestroyCrystalsCard) String() string                  { return "Destroy Crystals - Crystals of your enemy -8" }
func (_ *DestroyCrystalsCard) GetConsumption() CardConsumption { return CardConsumption{4, Crystals} }

// Conjure Cards

type ConjureBricksCard struct{}

func (_ *ConjureBricksCard) Play(me *Player, opponent *Player) { me.SetBricks(me.Bricks + 8) }
func (_ *ConjureBricksCard) String() string                    { return "Conjure Bricks - Bricks +8" }
func (_ *ConjureBricksCard) GetConsumption() CardConsumption   { return CardConsumption{4, Crystals} }

type ConjureWeaponsCard struct{}

func (_ *ConjureWeaponsCard) Play(me *Player, opponent *Player) { me.SetWeapons(me.Weapons + 8) }
func (_ *ConjureWeaponsCard) String() string                    { return "Conjure Weapons - Weapons +8" }
func (_ *ConjureWeaponsCard) GetConsumption() CardConsumption   { return CardConsumption{4, Crystals} }

type ConjureCrystalsCard struct{}

func (_ *ConjureCrystalsCard) Play(me *Player, opponent *Player) { me.SetCrystals(me.Crystals + 8) }
func (_ *ConjureCrystalsCard) String() string                    { return "Conjure Crystals - Crystals +8" }
func (_ *ConjureCrystalsCard) GetConsumption() CardConsumption   { return CardConsumption{4, Crystals} }

// Production cards

type SchoolCard struct{}

func (_ *SchoolCard) Play(me *Player, opponent *Player) { me.SetBuilders(me.Builders + 1) }
func (_ *SchoolCard) String() string                    { return "School - Builders +1" }
func (_ *SchoolCard) GetConsumption() CardConsumption   { return CardConsumption{8, Bricks} }

type SorcererCard struct{}

func (_ *SorcererCard) Play(me *Player, opponent *Player) { me.SetMages(me.Mages + 1) }
func (_ *SorcererCard) String() string                    { return "Sorcerer - Soldiers +1" }
func (_ *SorcererCard) GetConsumption() CardConsumption   { return CardConsumption{8, Crystals} }

type RecruitingCard struct{}

func (_ *RecruitingCard) Play(me *Player, opponent *Player) { me.SetSoldiers(me.Mages + 1) }
func (_ *RecruitingCard) String() string                    { return "Sorcerer - Mages +1" }
func (_ *RecruitingCard) GetConsumption() CardConsumption   { return CardConsumption{8, Weapons} }

// Miscellaneous cards

type CurseCard struct{}

func (_ *CurseCard) Play(me *Player, opponent *Player) {
	me.SetBuilders(me.Builders + 1)
	me.SetMages(me.Mages + 1)
	me.SetSoldiers(me.Soldiers + 1)

	opponent.SetBuilders(opponent.Builders - 1)
	opponent.SetMages(opponent.Mages - 1)
	opponent.SetSoldiers(opponent.Soldiers - 1)

}
func (_ *CurseCard) String() string                  { return "Sorcerer - All +1, All of your enemy -1" }
func (_ *CurseCard) GetConsumption() CardConsumption { return CardConsumption{25, Crystals} }

// type ThiefCard struct{}

// func (_ *ThiefCard) Play(me *Player, opponent *Player) {
// 	me.SetBuilders(me.GetBuilders() + 1)
// 	me.SetMages(me.GetMages() + 1)
// 	me.SetSoldiers(me.GetMages() + 1)

// 	opponent.SetBuilders(opponent.GetBuilders() - 1)
// 	opponent.SetMages(opponent.GetMages() - 1)
// 	opponent.SetSoldiers(opponent.GetMages() - 1)

// }
// func (_ *ThiefCard) String() string                  { return "Sorcerer - All +1, All of your enemy -1" }
// func (_ *ThiefCard) GetConsumption() CardConsumption { return CardConsumption{25, Crystals} }

type SaboteurCard struct{}

func (_ *SaboteurCard) Play(me *Player, opponent *Player) {
	opponent.SetBricks(opponent.Bricks - 4)
	opponent.SetWeapons(opponent.Weapons - 4)
	opponent.SetCrystals(opponent.Crystals - 4)
}
func (_ *SaboteurCard) String() string                  { return "Saboteur - Reserves of your enemy -4" }
func (_ *SaboteurCard) GetConsumption() CardConsumption { return CardConsumption{12, Weapons} }
