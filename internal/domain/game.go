package domain

type Cities interface {
	ListCities() []string
	CityConnections(cityName string) []string
}
type Alien interface {
	Location() string
}
type Game interface {
	NewRound() bool
	String() string
}
