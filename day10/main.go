package day10

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

func FindConnectedPeaks(prevHeight int, grid [][]rune, x int, y int) []Location {
	if x < 0 || y < 0 || y >= len(grid) || x >= len(grid[y]) {
		return []Location{
			{-1, -1},
		}
	}

	value := grid[y][x]
	number := int(value - '0')

	if number != prevHeight+1 {
		return []Location{
			{-1, -1},
		}
	}

	if number == 9 {
		return []Location{{x, y}}
	}

	var locations []Location
	locations = append(locations, FindConnectedPeaks(number, grid, x-1, y)...)
	locations = append(locations, FindConnectedPeaks(number, grid, x+1, y)...)
	locations = append(locations, FindConnectedPeaks(number, grid, x, y-1)...)
	locations = append(locations, FindConnectedPeaks(number, grid, x, y+1)...)
	return locations
}

func SolveA() {
	count := 0
	countB := 0

	grid := utils.StringsToRunes(strings.Split(input, "\n"))

	startingLocations := FindStartingLocations(grid)
	for _, location := range startingLocations {
		locations := FindConnectedPeaks(-1, grid, location.x, location.y)

		// Count the number of unique peaks reached
		uniqueLocations := make(map[Location]bool)
		for _, location := range locations {
			if location.x == -1 {
				continue
			}
			uniqueLocations[location] = true
		}

		for range uniqueLocations {
			count++
		}

		// Count the of peaks reached by unique paths
		for _, location := range locations {
			if location.x == -1 {
				continue
			}
			countB++
		}
	}
	fmt.Println(count)
	fmt.Println(countB)

	fmt.Printf("%+v\n", startingLocations)
}

func FindStartingLocations(grid [][]rune) []Location {
	var startingLocations []Location

	for y, row := range grid {
		for x, cell := range row {
			if cell == '0' {
				startingLocations = append(startingLocations, Location{x, y})
			}
		}
	}

	return startingLocations
}
