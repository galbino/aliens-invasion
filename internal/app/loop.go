package app

import (
	"fmt"

	"github.com/galbino/alien-assignment/internal/implementation"
)

func GameLoop(cityFile string, amount int) {
	cities := implementation.NewCities(cityFile)
	game := implementation.NewGame(cities, amount)
	for game.NewRound() {
		continue
	}
	fmt.Println(game)
}
