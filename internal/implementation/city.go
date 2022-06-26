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
	cities := []string{}
	for k := range c {
		cities = append(cities, k)
	}
	return cities
}

func (c cities) CityConnections(cityName string) []string {
	cities := []string{}
	for _, k := range c[cityName] {
		cities = append(cities, strings.Split(k, "|")[0])
	}
	return cities
}

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

func (c cities) DestroyCity(cityName string) {
	for _, conn := range c[cityName] {
		cityData := strings.Split(conn, "|")
		direction := cityData[1]
		cityDestiny := cityData[0]
		toDelete := fmt.Sprintf("%s|%s", cityName, InvertSides(direction))
		for ind, connections := range c[cityDestiny] {
			if connections == toDelete {
				if len(c[cityDestiny]) == 1 {
					c[cityDestiny] = c[cityDestiny][:0]
				} else if ind == len(c[cityDestiny])-1 {
					c[cityDestiny] = c[cityDestiny][ind:]
				} else {
					c[cityDestiny] = append(c[cityDestiny][:ind], c[cityDestiny][ind+1:]...)
				}
			}
		}
	}
	delete(c, cityName)
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
