package implementation

import (
	"fmt"
	"math/rand"

	"github.com/galbino/alien-assignment/internal/domain"
)

func NewAlien(amount int, cities []string) []domain.Alien {
	aliens := []domain.Alien{}
	for i := 0; i < amount; i++ {
		fmt.Println(cities)
		rng := rand.Intn(len(cities))
		alien := domain.Alien{Name: fmt.Sprintf("%d", i), Location: cities[rng]}
		aliens = append(aliens, alien)
	}
	return aliens
}
