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
	guard Guard
}

func (g *GuardPatrol) GetGuard() Guard {
	for y, row := range g.Map {
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

//go:embed input.txt
var input string

func SolveA() {
	patrolMap := ParseMap(input)
	fmt.Print(utils.RunesToString(patrolMap))
}

// Parsing
func ParseMap(input string) Map {
	var result Map
	for _, line := range strings.Split(input, "\n") {
		result = append(result, []rune(line))
	}
	return result
}
