package day8

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolveA(t *testing.T) {
	SolveA()
}

func TestRuneGridToMapOfLocations(t *testing.T) {
	grid := [][]rune{
		[]rune("aba"),
	}

	locations := RuneGridToMapOfLocations(grid)

	assert.Equal(t, 2, len(locations['a']))
	assert.Equal(t, 0, locations['a'][0].x)
	assert.Equal(t, 0, locations['a'][0].y)
	assert.Equal(t, 2, locations['a'][1].x)
	assert.Equal(t, 0, locations['a'][1].y)

	assert.Equal(t, 1, len(locations['b']))
	assert.Equal(t, 1, locations['b'][0].x)
	assert.Equal(t, 0, locations['b'][0].y)
}
