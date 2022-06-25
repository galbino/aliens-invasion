package main

import (
	"flag"

	"github.com/galbino/alien-assignment/internal/app"
)

func main() {
	cityFile := flag.String("input", "cities", "the name of city input file")
	aliens := flag.Int("n", 10, "the amount of aliens to be spawned")
	app.GameLoop(*cityFile, *aliens)
}
