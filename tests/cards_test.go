package card_test

import (
	"testing"

	. "github.com/quarkus7/mravenci/pkg"
)

func TestSorcererCard(t *testing.T) {
	cards := NewCardsPackage()
	me, opponent := NewPlayer(cards), NewPlayer(cards)

	card := SorcererCard{}
	card.Play(me, opponent)

	if me.Mages != 3 {
		t.Errorf("Unexpected result after Sorcerer play.")
	}
}

func TestCurseCard(t *testing.T) {
	cards := NewCardsPackage()
	me, opponent := NewPlayer(cards), NewPlayer(cards)

	card := CurseCard{}
	card.Play(me, opponent)

	if me.Builders != 3 || me.Mages != 3 || me.Soldiers != 3 || opponent.Builders != 1 || opponent.Mages != 1 || opponent.Soldiers != 1 {
		t.Errorf("Curse card play failed!")
	}
}
