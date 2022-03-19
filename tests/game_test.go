package game_test

import (
	"testing"

	. "github.com/quarkus7/mravenci/pkg"
)

func TestGameStart(t *testing.T) {
	game := game.NewGame()
	me := game.CurrentPlayer

	initialBricks, initialWeapons, initialCrystals := me.Bricks, me.Weapons, me.Crystals

	if initialBricks != 5 || initialWeapons != 5 || initialCrystals != 5 {
		t.Errorf("Initial material should be 5 each, [%d, %d, %d] found", initialBricks, initialWeapons, initialCrystals)
	}

	game.StartRound()
	game.EndRound()

	if me != game.OpponentPlayer {
		t.Errorf("After first round players must change")
	}

	if me.Bricks != initialBricks || me.Weapons != initialWeapons || me.Crystals != initialCrystals {
		t.Errorf("After first round 1st player's materials must remain the same")
	}
}

func TestGameSecondRound(t *testing.T) {
	game := game.NewGame()
	me, opponent := game.CurrentPlayer, game.OpponentPlayer

	initialBricks, initialWeapons, initialCrystals := me.Bricks, me.Weapons, me.Crystals

	game.StartRound()
	game.EndRound()

	game.StartRound()

	if opponent.Bricks == initialBricks || opponent.Weapons == initialWeapons || opponent.Crystals == initialCrystals {
		t.Errorf("When second round starts 2nd player's materials must change")
	}
}
