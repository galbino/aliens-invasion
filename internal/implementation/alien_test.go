package implementation_test

import (
	"testing"

	. "github.com/galbino/alien-assignment/internal/implementation"
	"golang.org/x/exp/slices"
)

func Test_NewAlien(t *testing.T) {
	amount := 15
	cities := []string{"stub1", "stub2"}
	aliens := NewAlien(amount, cities)
	if len(aliens) != amount {
		t.Errorf("The created amount of aliens was incorrect, got: %d, want: %d", len(aliens), amount)
	}
	for _, alien := range aliens {
		if !slices.Contains(cities, alien.Location) {
			t.Errorf("Alien was created with wrong city, got: %s, want: %s", alien.Location, cities)
		}
	}
}
