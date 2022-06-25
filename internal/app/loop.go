package app

import (
	"fmt"

	"github.com/galbino/alien-assignment/internal/implementation"
)

func GameLoop(cityFile string, amount int) {
	game := implementation.NewGame(cityFile, amount)
	for game.NewRound() {
		continue
	}
	fmt.Println(game)
}
