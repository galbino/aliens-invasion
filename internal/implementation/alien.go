package implementation

import (
	"fmt"
	"math/rand"

	"github.com/galbino/alien-assignment/internal/domain"
)

type alien struct {
	name     string
	location string
}

func (a alien) Location() string {
	return a.location
}

func NewAlien(amount int, cities []string) []domain.Alien {
	aliens := []domain.Alien{}
	for i := 0; i < amount; i++ {
		rng := rand.Intn(len(cities))
		alien := alien{name: fmt.Sprintf("%d", i), location: cities[rng]}
		aliens = append(aliens, alien)
	}
	return aliens
}
