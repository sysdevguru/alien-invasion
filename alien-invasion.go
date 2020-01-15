package main

import (
	"fmt"
	"os"
	"strconv"

	."github.com/sysdevguru/alien-invasion/step"
)

var (
	alienArr = []string{}
)

func usage() {
	fmt.Printf("alien-invasion: usage\n")
	fmt.Printf("\t\t./alien-invasion [number of aliens]\n")
}

func main() {
	if len(os.Args[1:]) == 0 || len(os.Args[1:]) > 1 {
		usage()
		return
	}

	// generate aliens
	numbAliens := os.Args[1:][0]
	num, err := strconv.Atoi(numbAliens)
	if err != nil {
		fmt.Printf("alien-invasion: invalid aliens number\n")
		usage()
		return
	}
	for i := 0; i < num; i++ {
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
