package day6

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/wseubring/aoc2024/utils"
)

type Map [][]rune

type Position struct {
	X int
	Y int
}

type Direction struct {
	X int
	Y int
}

type GuardPatrol struct {
	Map   Map
	Guard Guard
}

func (m Map) GetGuard() Guard {
	for y, row := range m {
		for x, cell := range row {
			if cell == '^' {
				return Guard{
					Position:  Position{X: x, Y: y},
					Direction: Direction{X: 0, Y: -1},
				}
			}
		}
	}

	panic("No guard found")
}

type Guard struct {
	Position  Position
	Direction Direction
}

func (g *Guard) TurnRight() {
	newX := -g.Direction.Y
	newY := g.Direction.X
	g.Direction.X = newX
	g.Direction.Y = newY
}

func (g *Guard) MoveForward() {
	g.Position.X += g.Direction.X
	g.Position.Y += g.Direction.Y
}

func (g *Guard) GetNextLocation() Position {
	nextPosition := Position{
		X: g.Position.X + g.Direction.X,
		Y: g.Position.Y + g.Direction.Y,
	}

	return nextPosition
}

func (m Map) IsObstacle(position Position) bool {
	return m[position.Y][position.X] == '#'
}

func (m Map) IsInBounds(position Position) bool {
	return position.X >= 0 && position.Y >= 0 && position.Y < len(m) && position.X < len(m[position.Y])
}

func (m Map) Mark(position Position) {
	m[position.Y][position.X] = 'X'
}

func (gp GuardPatrol) MarkPatrol() {
	gp.Map.Mark(gp.Guard.Position)

	for {
		nextPosition := gp.Guard.GetNextLocation()

		if !gp.Map.IsInBounds(nextPosition) {
			break
		}

		if !gp.Map.IsObstacle(nextPosition) {
			gp.Guard.MoveForward()
		} else if gp.Map.IsObstacle(nextPosition) {
			gp.Guard.TurnRight()
		}

		gp.Map.Mark(gp.Guard.Position)
	}
}

func (m Map) CountMarked() int {
	count := 0
	for _, row := range m {
		for _, cell := range row {
			if cell == 'X' {
				count++
			}
		}
	}
	return count
}

//go:embed input.txt
var input string

func SolveA() {
	patrolMap := ParseMap(input)
	guard := patrolMap.GetGuard()

	patrol := GuardPatrol{
		Map:   patrolMap,
		Guard: guard,
	}

	patrol.MarkPatrol()

	fmt.Println(utils.RunesToString(patrol.Map))
	fmt.Println(patrol.Map.CountMarked())
}

// Parsing
func ParseMap(input string) Map {
	var result Map
	for _, line := range strings.Split(input, "\n") {
		result = append(result, []rune(line))
	}
	return result
}
