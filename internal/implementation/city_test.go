package implementation_test

import (
	"reflect"
	"testing"

	. "github.com/galbino/alien-assignment/internal/implementation"
)

func Test_NewCity(t *testing.T) {
	amountCities := 5
	bar := []string{"Foo", "Bee"}
	foo := []string{"Bar", "Baz", "Qu-ux"}
	baz := []string{"Foo"}
	bee := []string{"Bar"}
	qu := []string{"Foo"}
	cityInt := NewCities("../../cities_test")
	cities := cityInt.ListCities()
	if len(cities) != amountCities {
		t.Errorf("Cities created wrong, want: %d, got: %d", amountCities, len(cities))
	}
	if !reflect.DeepEqual(cityInt.CityConnections("Bar"), bar) {
		t.Errorf("Bar created wrong, want: %s, got: %s", bar, cityInt.CityConnections("Bar"))
	}
	if !reflect.DeepEqual(cityInt.CityConnections("Foo"), foo) {
		t.Errorf("Foo created wrong, want: %s, got: %s", foo, cityInt.CityConnections("Foo"))
	}
	if !reflect.DeepEqual(cityInt.CityConnections("Baz"), baz) {
		t.Errorf("Baz created wrong, want: %s, got: %s", baz, cityInt.CityConnections("Baz"))
	}
	if !reflect.DeepEqual(cityInt.CityConnections("Bee"), bee) {
		t.Errorf("Bee created wrong, want: %s, got: %s", bee, cityInt.CityConnections("Bee"))
	}
	if !reflect.DeepEqual(cityInt.CityConnections("Qu-ux"), qu) {
		t.Errorf("Qu-ux created wrong, want: %s, got: %s", qu, cityInt.CityConnections("Qu-ux"))
	}
}
