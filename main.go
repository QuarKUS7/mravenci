package main

import (
	"math/rand"
	"time"

	pkg "github.com/quarkus7/mravenci/pkg"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	game := pkg.NewGame()
	controller := pkg.NewController(game)

	controller.Run()
}
