package day6

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolveA(t *testing.T) {
	SolveA()
}

func TestGetGuard(t *testing.T) {
	m := Map([][]rune{
		[]rune("   "),
		[]rune(" ^ "),
		[]rune("   "),
	})
	guard := m.GetGuard()
	assert.Equal(t, 1, guard.Position.X)
	assert.Equal(t, 1, guard.Position.Y)
	assert.Equal(t, 0, guard.Direction.X)
	assert.Equal(t, -1, guard.Direction.Y)
}

func TestTurnRight(t *testing.T) {
	g := Guard{
		Position:  Position{X: 0, Y: 0},
		Direction: Direction{X: 0, Y: -1},
	}
	g.TurnRight()
	assert.Equal(t, 1, g.Direction.X)
	assert.Equal(t, 0, g.Direction.Y)
	g.TurnRight()
	assert.Equal(t, 0, g.Direction.X)
	assert.Equal(t, 1, g.Direction.Y)
	g.TurnRight()
	assert.Equal(t, -1, g.Direction.X)
	assert.Equal(t, 0, g.Direction.Y)
	g.TurnRight()
	assert.Equal(t, 0, g.Direction.X)
	assert.Equal(t, -1, g.Direction.Y)
}

func TestMoveForward(t *testing.T) {
	g := Guard{
		Position:  Position{X: 0, Y: 0},
		Direction: Direction{X: 0, Y: -1},
	}
	g.MoveForward()

	assert.Equal(t, 0, g.Position.X)
	assert.Equal(t, -1, g.Position.Y)
}
