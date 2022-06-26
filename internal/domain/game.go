package domain

//go:generate mockgen -destination=../mocks/mock_cities.go -package=mocks github.com/galbino/alien-assignment/internal/domain Cities
//go:generate mockgen -destination=../mocks/mock_alien.go -package=mocks github.com/galbino/alien-assignment/internal/domain Alien
//go:generate mockgen -destination=../mocks/mock_game.go -package=mocks github.com/galbino/alien-assignment/internal/domain Game
type Cities interface {
	ListCities() []string
	CityConnections(cityName string) []string
	String() string
	DestroyCity(cityName string)
}
type Alien interface {
	Name() int
	Location() string
	Walk(cities Cities)
	String() string
}
type Game interface {
	NewRound() bool
	String() string
	CountAliens() int
}
