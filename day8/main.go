package day8

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/wseubring/aoc2024/utils"
)

//go:embed input.txt
var input string

type Location struct {
	x int
	y int
}

type Vector struct {
	x int
	y int
}

func SolveA() {
	grid := utils.StringsToRunes(strings.Split(input, "\n"))

	signalLocations := RuneGridToMapOfLocations(grid)
	fmt.Printf("%+v\n", signalLocations)

	for _, locations := range signalLocations {
		for i, locationA := range locations {
			for j := i + 1; j < len(locations); j++ {
				locationB := locations[j]
				vector := Vector{locationB.x - locationA.x, locationB.y - locationA.y}

				for {
					newLocationA := Location{locationA.x - vector.x, locationA.y - vector.y}

					if !IsInGrid(grid, newLocationA) {
						break
					}

					grid[newLocationA.y][newLocationA.x] = '#'
					locationA = newLocationA
				}

				for {
					newLocationB := Location{locationB.x + vector.x, locationB.y + vector.y}

					if !IsInGrid(grid, newLocationB) {
						break
					}

					grid[newLocationB.y][newLocationB.x] = '#'
					locationB = newLocationB
				}

				// Part one
				// newLocationA := Location{locationA.x - vector.x, locationA.y - vector.y}
				// newLocationB := Location{locationB.x + vector.x, locationB.y + vector.y}

				// if IsInGrid(grid, newLocationA) {
				// 	grid[newLocationA.y][newLocationA.x] = '#'
				// }

				// if IsInGrid(grid, newLocationB) {
				// 	grid[newLocationB.y][newLocationB.x] = '#'
				// }
			}
		}
	}
	fmt.Println(utils.RunesToString(grid))
	fmt.Println(len(RuneGridToMapOfLocations(grid)['#']))
}

func IsInGrid(grid [][]rune, location Location) bool {
	return location.x >= 0 && location.y >= 0 && location.x < len(grid[0]) && location.y < len(grid)
}
func RuneGridToMapOfLocations(grid [][]rune) map[rune][]Location {
	locations := make(map[rune][]Location)

	for y, row := range grid {
		for x, r := range row {
			if r == '.' {
				continue
			}

			locations[r] = append(locations[r], Location{x, y})
		}
	}

	return locations
}
