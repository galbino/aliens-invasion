package implementation

import (
	"bufio"
	"log"
	"os"
	"strings"

	"github.com/galbino/alien-assignment/internal/domain"
	"golang.org/x/exp/slices"
)

func NewCities(cityFile string) domain.GameCities {
	cities := domain.GameCities{}
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
			leadingCity := strings.Split(conn, "=")[1]
			cities[city] = append(cities[city], leadingCity)
			if _, ok := cities[leadingCity]; !ok {
				cities[leadingCity] = []string{city}
			} else {
				if !slices.Contains(cities[leadingCity], city) {
					cities[leadingCity] = append(cities[leadingCity], city)
				}
			}
		}
	}
	return cities
}
