package implementation_test

import (
	"testing"

	. "github.com/galbino/alien-assignment/internal/implementation"
	"github.com/galbino/alien-assignment/internal/mocks"
	"github.com/golang/mock/gomock"
)

func Test_NewRound(t *testing.T) {
	game := NewGame("cities_test", 1)
	game.NewRound()
	destination := []string{"Bar"}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	cities := mocks.NewMockCities(ctrl)
	aliens := mocks.NewMockAlien(ctrl)
	cities.EXPECT().CityConnections(gomock.Any()).Return(destination)
	aliens.EXPECT().Location().Return(destination)
}
