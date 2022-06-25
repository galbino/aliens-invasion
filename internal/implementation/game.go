package implementation

import (
	"fmt"

	"github.com/galbino/alien-assignment/internal/domain"
)

type game struct {
	cities domain.GameCities
	aliens []domain.Alien
	rounds int
}

func (g *game) NewRound() bool {
	if g.rounds == 10001 {
		return false
	}
	g.rounds += 1
	return true
}

func NewGame(cityFile string, aliensAmount int) domain.Game {
	cities := NewCities(cityFile)
	cityArray := []string{}
	for k := range cities {
		cityArray = append(cityArray, k)
	}
	aliens := NewAlien(aliensAmount, cityArray)
	return &game{cities: cities, aliens: aliens, rounds: 0}
}

func (g game) String() string {
	return fmt.Sprintf("%d", g.rounds)
}
