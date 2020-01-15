package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDupCheck(t *testing.T) {
	place := []int{2, 4, 5, 2, 5, 6, 1}
	dupMap := make(map[int]int)
	dupMap[2] = 2
	dupMap[4] = 1
	dupMap[5] = 2
	dupMap[6] = 1
	dupMap[1] = 1

	assert.Equal(t, dupMap, DupCheck(place))
}

func TestRemoveCityAliens(t *testing.T) {
	cities := []string{"mar", "foo", "bar", "bee"}
	assert.Equal(t, []string{"mar", "foo", "bee"}, RemoveCityAliens(cities, 2))

	cities = []string{"mar", "foo", "bar", "bee"}
	assert.Equal(t, []string{"mar", "foo", "bar"}, RemoveCityAliens(cities, 3))

	cities = []string{"mar", "foo", "bar", "bee"}
	assert.Equal(t, []string{"mar", "bar", "bee"}, RemoveCityAliens(cities, 1))
}

func TestRemovePlace(t *testing.T) {
	places := []int{3, 4, 3, 2, 5}
	assert.Equal(t, []int{3, 4, 2, 5}, RemovePlaces(places, 2))

	places = []int{4, 3, 2, 5}
	assert.Equal(t, []int{4, 3, 2}, RemovePlaces(places, 3))

	places = []int{9, 3, 4, 3, 2, 5}
	assert.Equal(t, []int{9, 4, 3, 2, 5}, RemovePlaces(places, 1))
}

func TestFind(t *testing.T) {
	paths := []int{3, 4, 5, 2, 3}
	_, exists := Find(paths, 3)
	_, exists1 := Find(paths, 4)
	_, exists2 := Find(paths, 7)

	assert.Equal(t, true, exists)
	assert.Equal(t, true, exists1)
	assert.Equal(t, false, exists2)
}

func TestDeleteEmpty(t *testing.T) {
	slice := []string{"foo", "bar", "", "bee"}
	slice1 := []string{"foo", "bar", "mus", "bee", "", ""}
	slice2 := []string{"foo", "", "mus", "bee", "", ""}

	assert.Equal(t, []string{"foo", "bar", "bee"}, DeleteEmpty(slice))
	assert.Equal(t, []string{"foo", "bar", "mus", "bee"}, DeleteEmpty(slice1))
	assert.Equal(t, []string{"foo", "mus", "bee"}, DeleteEmpty(slice2))
}
