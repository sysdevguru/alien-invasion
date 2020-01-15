package step

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/sysdevguru/alien-invasion/util"
)

var (
	cityNum  = 0
	stepNum  = 0
	cityMap  = make(map[string]int)
	numMap   = make(map[int]string)
	pathMap  = make(map[int]int)
	moveMap  = make(map[int][]string)
	movePath = []int{}
	cityArr  = []string{}
	MapPath  = ""
)

// step represents status of Cities
// and Aliens after each Alien movement
type Step struct {
	Count     int
	PrevPlace []int
	CurPlace  []int
	Cities    []string
	Aliens    []string
}

// report prints current Aliens locations
func (s *Step) Report() {
	fmt.Printf("%d:alien-invasion: current Alien locations report ===> ", s.Count)
	for i := 0; i < len(s.CurPlace); i++ {
		fmt.Printf("%v:%v ", s.Aliens[i], numMap[s.CurPlace[i]])
	}
	fmt.Printf("\n")
}

// genPlace generates potential next city
// to which Aliens can move
func (s *Step) GenPlace() {
	// in case first step
	if len(s.PrevPlace) == 0 {
		place := []int{}
		for i := 0; i < len(s.Aliens); i++ {
			p := util.GenRanInt(cityNum)
			place = append(place, p)
		}
		s.CurPlace = place
		return
	}

	// non-first step
	place := []int{}
	for i := 0; i < len(s.Aliens); i++ {
		// check random path validation
		for {
			// next potential place
			p := util.GenRanInt(cityNum)
			np := s.PrevPlace[i]*10 + p

			// check next place is valid
			_, found := util.Find(movePath, np)
			if found {
				place = append(place, p)
				break
			}
		}
	}
	s.CurPlace = place
}

// simulate aliens movements
func (s *Step) Run() {
	dupMap := make(map[int]int)
	newStep := Step{}
	for {
		// aliens are all dead or reached movement limit
		if (s.Count == 10000) || len(s.Aliens) == 0 {
			break
		}

		// check aliens potential confliction
		isFight := false
		for _, v := range dupMap {
			if v > 1 {
				isFight = true
				break
			}
		}

		// aliens are fighting
		if isFight {
			// update movePath
			newStep.updatePath(dupMap)
			// remove city from map
			newStep.removeCity(dupMap)
			// kill aliens
			newStep.killAliens(dupMap)
		}

		// generate new step
		newStep.PrevPlace = s.CurPlace
		newStep.Count = s.Count + 1
		c := util.DeleteEmpty(s.Cities)
		newStep.Cities = c
		a := util.DeleteEmpty(s.Aliens)
		newStep.Aliens = a
		newStep.GenPlace()

		if len(newStep.Aliens) != 0 {
			// newStep.Report()
		}

		// get aliens confliction places info
		dupMap = util.DupCheck(newStep.CurPlace)

		// process with new step
		s.PrevPlace = newStep.PrevPlace
		s.CurPlace = newStep.CurPlace
		s.Cities = newStep.Cities
		s.Aliens = newStep.Aliens
		s.Count = newStep.Count
	}

	s.ReadCityInfo(false)
}

// read map file
func (s *Step) ReadCityInfo(initial bool) {
	file, err := os.Open(MapPath)
	if err != nil {
		fmt.Printf("alien-invasion: unexpected Reading file failure: %v\n", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if initial {
			cityNum++

			splitArr := strings.Split(strings.TrimSpace(scanner.Text()), " ")
			cityMap[splitArr[0]] = cityNum
			numMap[cityNum] = splitArr[0]
			cityArr = append(cityArr, splitArr[0])
			for i := 1; i < len(splitArr); i++ {
				nameArr := strings.Split(strings.TrimSpace(splitArr[i]), "=")
				moveMap[cityMap[splitArr[0]]] = append(moveMap[cityMap[splitArr[0]]], nameArr[1])
			}
			s.Cities = cityArr

			for k, v := range moveMap {
				for i := 0; i < len(v); i++ {
					s := k
					d := cityMap[v[i]]
					path := s*10 + d
					movePath = append(movePath, path)
				}
			}
		} else {
			// check remaining cities
			splitArr := strings.Split(scanner.Text(), " ")
			if len(s.Cities) == 0 {
				fmt.Printf("alien-invasion: no remaining cities\n")
				return
			}
			for _, v := range s.Cities {
				if v == splitArr[0] {
					fmt.Printf("%v\n", scanner.Text())
				}
			}
		}
	}
}

// updatePath updates loads status
func (s *Step) updatePath(dupMap map[int]int) {
	city := findCitiesToRemove(dupMap)

	newPath := []int{}
	for _, v1 := range movePath {
		flag := false
		for _, v := range city {
			if ((v+1)*10 > v1 && v1 > v*10) || (v1%10 == v) {
				flag = true
				break
			}
		}
		if flag {
			continue
		}
		newPath = append(newPath, v1)
	}
	movePath = newPath
}

// removeCity removes cities after aliens movement and fighting
func (s *Step) removeCity(dupMap map[int]int) {
	alienMap := s.findAliensInCity(dupMap)
	indexes := []int{}
	for _, v := range alienMap {
		indexes = append(indexes, s.CurPlace[v])
	}

	sort.Ints(indexes[:])
	for i := len(indexes) - 1; i >= 0; i-- {
		for j, v := range s.Cities {
			if v == numMap[indexes[i]] {
				util.RemoveCityAliens(s.Cities, j)
			}
		}
	}
}

// killAliens kills fighting aliens
func (s *Step) killAliens(dupMap map[int]int) {
	indexes := []int{}
	for i := 0; i < len(s.CurPlace); i++ {
		flag := false
		for _, v := range movePath {
			if ((s.CurPlace[i]+1)*10 > v && v > (s.CurPlace[i]*10)) || (v%10 == (s.CurPlace[i])) {
				flag = true
			}
		}
		if !flag {
			indexes = append(indexes, i)
		}
	}

	alienMap := s.findAliensInCity(dupMap)
	for _, v := range alienMap {
		// report
		cityName := numMap[s.CurPlace[v]]
		aliens := []string{}
		for i, v := range s.CurPlace {
			if v == cityMap[cityName] {
				aliens = append(aliens, s.Aliens[i])
			}
		}
		if len(aliens) <= 1 {
			continue
		}
		fmt.Printf("alien-invasion: %v has been destroyed by alien ", cityName)
		for i, v := range aliens {
			if i == len(aliens)-1 {
				fmt.Printf("%v\n", v)
				break
			}
			fmt.Printf("%v and alien ", v)
		}
	}

	sort.Ints(indexes[:])
	for i := len(indexes) - 1; i >= 0; i-- {
		util.RemoveCityAliens(s.Aliens, indexes[i])
		util.RemovePlaces(s.CurPlace, indexes[i])
	}
}

func (s *Step) findAliensInCity(dupMap map[int]int) map[int]int {
	aliens := make(map[int]int)
	for k, v := range dupMap {
		if v != 1 {
			for index, value := range s.CurPlace {
				if value == k {
					aliens[k] = index
				}
			}
		}
	}
	return aliens
}

func findCitiesToRemove(dupMap map[int]int) []int {
	city := []int{}
	for k, v := range dupMap {
		if v != 1 {
			city = append(city, k)
		}
	}
	return city
}
