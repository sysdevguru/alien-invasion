package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	. "github.com/sysdevguru/alien-invasion/step"
)

var (
	alienArr = []string{}
)

func usage() {
	fmt.Printf("alien-invasion: usage\n")
	fmt.Printf("\t\t./alien-invasion -aliens=[number of aliens] -map=[map file]\n")
}

func main() {
	if len(os.Args[1:]) == 0 || len(os.Args[1:]) > 2 {
		usage()
		return
	}

	numPtr := flag.Int("aliens", 3, "aliens number")
	mapPtr := flag.String("map", "maps/map.txt", "map file")
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
