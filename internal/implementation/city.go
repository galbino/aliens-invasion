package implementation

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/galbino/alien-assignment/internal/domain"
	"golang.org/x/exp/slices"
)

type cities map[string][]string

func (c cities) ListCities() []string {
	cities := make([]string, len(c))
	i := 0
	for k := range c {
		cities[i] = k
		i++
	}
	return cities
}

// Getting all connections from one city, excluding the direction
func (c cities) CityConnections(cityName string) []string {
	cities := []string{}
	for _, k := range c[cityName] {
		cities = append(cities, strings.Split(k, "|")[0])
	}
	return cities
}

// Function to facilitate printing the final result on screen
func (c cities) String() string {
	resp := []string{}
	for city := range c {
		tmp := city
		for _, conn := range c[city] {
			sp := strings.Split(conn, "|")
			tmp += fmt.Sprintf(" %s=%s", sp[1], sp[0])
		}
		resp = append(resp, tmp)
	}
	return strings.Join(resp, "\n")
}

// Function to delete a city from the list and all connections to it
func (c cities) DestroyCity(cityName string) {
	delete(c, cityName)
	for key := range c {
		for ind, conn := range c[key] {
			if strings.Contains(conn, cityName) {
				c[key] = slices.Delete(c[key], ind, ind+1)
				break
			}
		}
	}
}

func NewCities(cityFile string) domain.Cities {
	cities := cities{}
	f, err := os.Open(cityFile)
	if err != nil {
		log.Fatalf("error city.go: %v", err)
	}
	defer f.Close()
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		data := sc.Text()
		dataSplitted := strings.Split(data, " ")
		city := dataSplitted[0]
		cities[city] = []string{}
		for _, conn := range dataSplitted[1:] {
			cityData := strings.Split(conn, "=")
			location := cityData[0]
			leadingCity := cityData[1]
			if !slices.Contains(cities[city], leadingCity) {
				cities[city] = append(cities[city], fmt.Sprintf("%s|%s", leadingCity, location))
			}
			if _, ok := cities[leadingCity]; !ok {
				cities[leadingCity] = []string{fmt.Sprintf("%s|%s", city, InvertSides(location))}
			} else {
				if !slices.Contains(cities[leadingCity], fmt.Sprintf("%s|%s", city, InvertSides(location))) {
					cities[leadingCity] = append(cities[leadingCity], fmt.Sprintf("%s|%s", city, InvertSides(location)))
				}
			}
		}
	}
	return &cities
}

func InvertSides(side string) string {
	switch side {
	case "north":
		return "south"
	case "south":
		return "north"
	case "east":
		return "west"
	case "west":
		return "east"
	default:
		log.Fatal("Side doesn't exist.")
	}
	return ""
}
