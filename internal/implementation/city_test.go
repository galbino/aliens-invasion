package implementation_test

import (
	"fmt"
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
	cities := NewCities("../../cities_test")
	fmt.Println(cities)
	if len(cities) != amountCities {
		t.Errorf("Cities created wrong, want: %d, got: %d", amountCities, len(cities))
	}
	if !reflect.DeepEqual(cities["Bar"], bar) {
		t.Errorf("Bar created wrong, want: %s, got: %s", bar, cities["Bar"])
	}
	if !reflect.DeepEqual(cities["Foo"], foo) {
		t.Errorf("Foo created wrong, want: %s, got: %s", foo, cities["Foo"])
	}
	if !reflect.DeepEqual(cities["Baz"], baz) {
		t.Errorf("Baz created wrong, want: %s, got: %s", baz, cities["Baz"])
	}
	if !reflect.DeepEqual(cities["Bee"], bee) {
		t.Errorf("Bee created wrong, want: %s, got: %s", bee, cities["Bee"])
	}
	if !reflect.DeepEqual(cities["Qu-ux"], bee) {
		t.Errorf("Qu-ux created wrong, want: %s, got: %s", qu, cities["Qu-ux"])
	}
}
