package implementation

import (
	"fmt"

	"github.com/galbino/alien-assignment/internal/domain"
)

type game struct {
	cities domain.Cities
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
	aliens := NewAlien(aliensAmount, cities.ListCities())
	return &game{cities: cities, aliens: aliens, rounds: 0}
}

func (g game) String() string {
	return fmt.Sprintf("%d", g.rounds)
}
