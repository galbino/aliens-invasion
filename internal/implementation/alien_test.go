package implementation_test

import (
	"testing"

	. "github.com/galbino/alien-assignment/internal/implementation"
	"github.com/galbino/alien-assignment/internal/mocks"
	"github.com/golang/mock/gomock"
	"golang.org/x/exp/slices"
)

var citiesStry = []string{"stub1", "stub2"}

func Test_NewAlien(t *testing.T) {
	amount := 15
	aliens := NewAlien(amount, citiesStry)
	if len(aliens) != amount {
		t.Errorf("The created amount of aliens was incorrect, got: %d, want: %d", len(aliens), amount)
	}
	for _, alien := range aliens {
		if !slices.Contains(citiesStry, alien.Location()) {
			t.Errorf("Alien was created with wrong city, got: %s, want: %s", alien.Location(), citiesStry)
		}
	}
}

func Test_AlienWalk(t *testing.T) {
	amount := 1
	destination := []string{"Bar"}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	alien := NewAlien(amount, citiesStry)[0]

	cities := mocks.NewMockCities(ctrl)
	cities.EXPECT().CityConnections(alien.Location()).Return(destination)
	alien.Walk(cities)
	if alien.Location() != destination[0] {
		t.Errorf("Alien moved to wrong city, got: %s, want: %s", alien.Location(), destination[0])
	}
}
