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
	if g.checkState() {
		for i := 0; i < len(g.aliens); i++ {
			alien := g.aliens[i]
			alien.Walk(g.cities)
		}
		g.checkAliens()
		g.rounds += 1
		return true
	}
	return false
}

func (g game) checkState() bool {
	if g.rounds == 10001 || len(g.aliens) == 0 {
		return false
	}
	return true
}

func (g *game) checkAliens() {
	locations := map[string][]int{}
	for _, alien := range g.aliens {
		if val, ok := locations[alien.Location()]; ok {
			if len(val) == 1 {
				fmt.Printf("%s has been destroyed by alien %d and alien %d!\n", alien.Location(), val[0], alien.Name())
			}
			locations[alien.Location()] = append(locations[alien.Location()], alien.Name())
		} else {
			locations[alien.Location()] = []int{alien.Name()}
		}
	}
	for k := range locations {
		if len(locations[k]) > 1 {
			for _, val := range locations[k] {
				for ind, al := range g.aliens {
					if al.Name() == val {
						if len(g.aliens) == 1 {
							g.aliens = g.aliens[:0]
						} else if ind == len(g.aliens)-1 {
							g.aliens = g.aliens[ind:]
						} else {
							g.aliens = append(g.aliens[:ind], g.aliens[ind+1:]...)
						}
						break
					}
				}
			}
			g.cities.DestroyCity(k)
		}
	}
}

func NewGame(cities domain.Cities, amount int) domain.Game {
	aliens := NewAlien(amount, cities.ListCities())
	return &game{cities: cities, aliens: aliens, rounds: 0}
}

func (g game) String() string {
	return fmt.Sprintf("%s", g.cities)
}

func (g game) CountAliens() int {
	return len(g.aliens)
}
