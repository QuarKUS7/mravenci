package main

import (
	src "github.com/quarkus7/mravenci/pkg"
)

func main() {
	game := src.NewGame()
	controller := src.NewController(game)

	controller.Run()
}
