package util

import (
	"math/rand"
	"time"
)

func DupCheck(places []int) map[int]int {
	dupMap := make(map[int]int)
	for _, place := range places {
		_, exist := dupMap[place]
		if exist {
			dupMap[place] += 1
		} else {
			dupMap[place] = 1
		}
	}
	return dupMap
}

func RemoveCityAliens(ca []string, i int) []string {
	copy(ca[i:], ca[i+1:])
	ca[len(ca)-1] = ""
	ca = ca[:len(ca)-1]
	return ca
}

func RemovePlaces(p []int, i int) []int {
	copy(p[i:], p[i+1:])
	p[len(p)-1] = 0
	p = p[:len(p)-1]
	return p
}

func GenRanInt(max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	i := 1 + rand.Intn(max)
	return i
}

func Find(paths []int, p int) (int, bool) {
	for i, v := range paths {
		if v == p {
			return i, true
		}
	}
	return -1, false
}

func DeleteEmpty(s []string) []string {
	r := []string{}
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}
