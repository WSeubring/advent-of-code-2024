package day4

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/wseubring/aoc2024/utils"
)

type Puzzle [][]rune

//go:embed input.txt
var input string

func LoadPuzzle() Puzzle {
	lines := strings.Split(input, "\n")
	grid := utils.StringsToRunes(lines)
	return Puzzle(grid)
}

var directions = []struct {
	dx, dy int
}{
	{1, 0},   // right
	{0, 1},   // down
	{-1, 0},  // left
	{0, -1},  // up
	{1, 1},   // down-right
	{-1, 1},  // down-left
	{1, -1},  // up-right
	{-1, -1}, // up-left
}

func (p Puzzle) GetStringInDirection(x, y, dx, dy int) string {
	var builder strings.Builder
	for {
		if x < 0 || x >= len(p[0]) || y < 0 || y >= len(p) {
			break
		}
		builder.WriteRune(p[y][x])
		x += dx
		y += dy
	}
	return builder.String()
}

func (p Puzzle) FindWordOccurrences(word string) int {
	occurrences := 0
	for y := 0; y < len(p); y++ {
		for x := 0; x < len(p[y]); x++ {
			for _, dir := range directions {
				// For each point in the grid, check all directions
				w := p.GetStringInDirection(x, y, dir.dx, dir.dy)
				if strings.HasPrefix(w, word) {
					occurrences++
				}
			}
		}
	}
	return occurrences
}

func SolveA() {
	puzzle := LoadPuzzle()
	fmt.Println(puzzle.FindWordOccurrences("XMAS"))
}
