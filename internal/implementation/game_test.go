package implementation_test

import (
	"testing"

	. "github.com/galbino/alien-assignment/internal/implementation"
	"github.com/galbino/alien-assignment/internal/mocks"
	"github.com/golang/mock/gomock"
)

func Test_NewRound(t *testing.T) {
	dest := "Bar"
	destination := []string{dest}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	cities := mocks.NewMockCities(ctrl)
	cities.EXPECT().CityConnections(gomock.Any()).Return(destination).AnyTimes()
	cities.EXPECT().ListCities().Return(destination)
	cities.EXPECT().DestroyCity(gomock.Any()).Times(1)
	game := NewGame(cities, 2)
	game.NewRound()
	if game.NewRound() != false {
		t.Errorf("Aliens werent killed by each other, want: %d aliens, got: %d aliens", 0, game.CountAliens())
	}
}

func Test_CountAlien(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	cities := mocks.NewMockCities(ctrl)
	cities.EXPECT().ListCities().Return([]string{"foo"}).AnyTimes()
	game := NewGame(cities, 2)
	if game.CountAliens() != 2 {
		t.Errorf("Aliens count wrong, want: %d aliens, got: %d aliens", 2, game.CountAliens())
	}
}
