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

func Test_DestroyCity(t *testing.T) {
	foo := []string{"Baz", "Qu-ux"}
	bee := []string{}
	cities := NewCities("../../cities_test")
	cities.DestroyCity("Bar")
	if !reflect.DeepEqual(cities.CityConnections("Bee"), bee) {
		t.Errorf("Bee created wrong, want: %s, got: %s", bee, cities.CityConnections("Bee"))
	}
	if !reflect.DeepEqual(cities.CityConnections("Foo"), foo) {
		t.Errorf("Foo created wrong, want: %s, got: %s", foo, cities.CityConnections("Foo"))
	}
}

func Test_String(t *testing.T) {
	exp := "Foo north=Bar west=Baz south=Qu-ux\nBar south=Foo west=Bee\nBaz east=Foo\nQu-ux north=Foo\nBee east=Bar"
	cities := NewCities("../../cities_test")
	if cities.String() != exp {
		t.Errorf("Printed wrong, want: %s, got: %s", exp, cities.String())
	}
}

func Test_ListCities(t *testing.T) {
	cities := NewCities("../../cities_test")
	cityList := []string{"Bee", "Foo", "Bar", "Baz", "Qu-ux"}
	for _, city := range cityList {
		found := false
		for _, cityCreated := range cities.ListCities() {
			if city == cityCreated {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Listed cities wrongly, want: %s, got: %s", cityList, cities.ListCities())
		}
	}
}
