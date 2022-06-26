package implementation

import (
	"fmt"
	"math/rand"

	"github.com/galbino/alien-assignment/internal/domain"
)

type alien struct {
	name     int
	location string
}

func (a alien) Location() string {
	return a.location
}

func (a alien) Name() int {
	return a.name
}

func (a alien) String() string {
	return fmt.Sprintf("%d - %s", a.name, a.location)
}

func (a *alien) Walk(cities domain.Cities) {
	connectionList := cities.CityConnections(a.location)
	if len(connectionList) > 0 {
		rng := rand.Intn(len(connectionList))
		moveTo := connectionList[rng]
		a.location = moveTo
	}
}

func NewAlien(amount int, cities []string) []domain.Alien {
	aliens := []domain.Alien{}
	for i := 0; i < amount; i++ {
		rng := rand.Intn(len(cities))
		alien := alien{name: i, location: cities[rng]}
		aliens = append(aliens, &alien)
	}
	return aliens
}
