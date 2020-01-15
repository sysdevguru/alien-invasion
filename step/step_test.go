package step

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sysdevguru/alien-invasion/util"
)

var (
	step   = Step{}
	dupMap = make(map[int]int)
)

func init() {
	step.Count = 1
	step.PrevPlace = []int{3, 2, 1, 3, 2}
	step.CurPlace = []int{3, 2, 1, 3, 2}
	step.Cities = []string{"foo", "bar", "bee", "mar", "jus", "baz"}
	step.Aliens = []string{"a1", "a2", "a3", "a4", "a5"}

	dupMap[3] = 2
	dupMap[2] = 2
	dupMap[1] = 1

	cityMap["foo"] = 1
	cityMap["bar"] = 2
	cityMap["bee"] = 3
	cityMap["mar"] = 4
	cityMap["jus"] = 5
	cityMap["baz"] = 6
}

func TestFindCitiesToRemove(t *testing.T) {
	dups := findCitiesToRemove(dupMap)
	values := []int{}
	for _, v := range dups {
		values = append(values, v)
	}
	sort.Ints(values)

	assert.Equal(t, []int{2, 3}, values)
}

func TestFindAliensInCity(t *testing.T) {
	expected := make(map[int]int)
	expected[2] = 4
	expected[3] = 3

	assert.Equal(t, expected, step.findAliensInCity(dupMap))
}

func TestUpdatePath(t *testing.T) {
	movePath = []int{12, 13, 14, 21, 24, 25, 31, 36, 41, 42, 52, 63}
	step.updatePath(dupMap)
	assert.Equal(t, []int{14, 41}, movePath)

	dupMap[1] = 2
	movePath = []int{12, 13, 14, 21, 24, 25, 31, 36, 41, 42, 52, 63}
	step.updatePath(dupMap)
	assert.Equal(t, []int{}, movePath)
}

func TestRemoveCity(t *testing.T) {
	step.Count = 1
	step.PrevPlace = []int{4, 2, 5, 2, 6}
	step.CurPlace = []int{3, 3, 1, 3, 2}
	step.Cities = []string{"foo", "bar", "bee", "mar", "jus", "baz"}
	step.Aliens = []string{"a1", "a2", "a3", "a4", "a5"}

	dupMap[3] = 3
	dupMap[1] = 1
	dupMap[2] = 1

	numMap[1] = "foo"
	numMap[2] = "bar"
	numMap[3] = "bee"
	numMap[4] = "mar"
	numMap[5] = "jus"
	numMap[6] = "baz"

	step.removeCity(dupMap)
	assert.Equal(t, []string{"foo", "bar", "mar", "jus", "baz", ""}, step.Cities)

	cities := util.DeleteEmpty(step.Cities)
	assert.Equal(t, []string{"foo", "bar", "mar", "jus", "baz"}, cities)
}

func TestKillAlien(t *testing.T) {
	movePath = []int{12, 13, 21, 31, 36, 63}
	step.killAliens(dupMap)
	assert.Equal(t, []string{"a1", "a2", "a3", "a4", "a5"}, step.Aliens)
}
