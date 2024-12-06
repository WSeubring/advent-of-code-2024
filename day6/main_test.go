package day6

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolveA(t *testing.T) {
	SolveA()
}

func TestGetGuard(t *testing.T) {
	g := GuardPatrol{
		Map: [][]rune{
			[]rune("  ^  "),
			[]rune("     "),
			[]rune("     "),
		},
	}
	guard := g.GetGuard()
	assert.Equal(t, 2, guard.Position.X)
	assert.Equal(t, 0, guard.Position.Y)
	assert.Equal(t, 0, guard.Direction.X)
	assert.Equal(t, -1, guard.Direction.Y)
}
