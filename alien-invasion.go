package main

import (
	"flag"
	"fmt"
	"strconv"

	. "github.com/sysdevguru/alien-invasion/step"
)

const (
	// DefaultNumberOfAliens if number of Aliens is not provided
	DefaultNumberOfAliens int = 10
	// DefaultMapFile used if map file is not provided
	DefaultMapFile = "maps/map.txt"
)

var (
	alienArr = []string{}
)

func main() {
	numPtr := flag.Int("aliens", DefaultNumberOfAliens, "aliens number")
	mapPtr := flag.String("map", DefaultMapFile, "map file")
	flag.Parse()

	// get map file path
	MapPath = *mapPtr

	// generate aliens
	for i := 0; i < *numPtr; i++ {
		a := strconv.Itoa(i + 1)
		alienArr = append(alienArr, "a"+a)
	}
	fmt.Printf("alien-invasion: Alien invasion detected\n")

	// initialize first step
	firstStep := Step{
		Count:  1,
		Aliens: alienArr,
	}
	// read map file
	firstStep.ReadCityInfo(true)
	// generate aliens first place
	firstStep.GenPlace()
	// report aliens first positions
	firstStep.Report()
	// simulate aliens movements
	firstStep.Run()
}
