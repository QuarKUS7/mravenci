package pkg

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strconv"
)

type Controller interface {
	Run()
}

type controller struct {
	game *Game
	view View
}

func NewController(game *Game) Controller {
	view := View{game}
	return &controller{game, view}
}

func clearTerminal() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}

func (controller *controller) Run() {
	for {
		clearTerminal()
		controller.game.StartRound()
		fmt.Println(controller.view.RenderGame())

		if ok := controller.game.HasWon(); ok {
			controller.view.RenderWinner()
			return
		}

		promptResult := controller.prompt()
		for promptResult != promptSuccess {
			if promptResult == promptFailed {
				fmt.Println("Invalid move")
			}
			promptResult = controller.prompt()
		}

		controller.game.EndRound()
	}

	panic("Error occured, please try a new game.")
}

type promptState int

const (
	promptSuccess promptState = iota
	promptFailed
	promptHelp
)

func (controller *controller) prompt() promptState {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Select card to play: ")
	text, _ := reader.ReadString('\n')
	text = text[:len(text)-1]

	if text == "help" {
		controller.showHelp()
		return promptHelp
	} else {
		numberMatched, _ := regexp.MatchString(`^[0-9]$`, text)
		if numberMatched {
			card_index, _ := strconv.Atoi(text)
			return controller.playCard(card_index)
		}
	}

	return promptFailed
}

func (controller *controller) showHelp() {
	fmt.Printf(" - type the number of the card you desire to play\n\n")
}

func (controller *controller) playCard(card_index int) promptState {
	if !controller.game.IsValidToPlayNthCard(card_index) {
		controller.discardCard(card_index)
		return promptSuccess
	}

	controller.game.PlayNthCard(card_index)
	return promptSuccess
}

func (controller *controller) discardCard(card_index int) {
	controller.game.DiscardNthCard(card_index)
}
