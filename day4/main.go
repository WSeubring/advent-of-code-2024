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

// Part B
// First rewrite use a kernel instead
var xmasKernels = [][][]rune{
	{
		{'X', 'M', 'A', 'S'},
	},
	{
		{'S', 'A', 'M', 'X'},
	},
	{
		{'X'},
		{'M'},
		{'A'},
		{'S'},
	},
	{
		{'S'},
		{'A'},
		{'M'},
		{'X'},
	},
	{
		{'X', ' ', ' ', ' '},
		{' ', 'M', ' ', ' '},
		{' ', ' ', 'A', ' '},
		{' ', ' ', ' ', 'S'},
	},
	{
		{' ', ' ', ' ', 'X'},
		{' ', ' ', 'M', ' '},
		{' ', 'A', ' ', ' '},
		{'S', ' ', ' ', ' '},
	},
	{
		{'S', ' ', ' ', ' '},
		{' ', 'A', ' ', ' '},
		{' ', ' ', 'M', ' '},
		{' ', ' ', ' ', 'X'},
	},
	{
		{' ', ' ', ' ', 'S'},
		{' ', ' ', 'A', ' '},
		{' ', 'M', ' ', ' '},
		{'X', ' ', ' ', ' '},
	},
}

func (p Puzzle) FindKernelOccurrences(kernel [][]rune) int {
	occurrences := 0
	for y := 0; y < len(p); y++ {
		for x := 0; x < len(p[y]); x++ {
			if p.MatchKernel(x, y, kernel) {
				occurrences++
			}
		}
	}
	return occurrences
}

func (p Puzzle) MatchKernel(x, y int, kernel [][]rune) bool {
	for ky, row := range kernel {
		for kx, char := range row {
			// Check if we are out of bounds
			if x+kx >= len(p[0]) || y+ky >= len(p) {
				return false
			}

			// Skip spaces
			if char == ' ' {
				continue
			}

			// Check if the character matches
			if p[y+ky][x+kx] != char {
				return false
			}
		}
	}
	return true
}

func (p Puzzle) FindKernelsOccurrences(kernels [][][]rune) int {
	occurrences := 0
	for _, kernel := range kernels {
		occurrences += p.FindKernelOccurrences(kernel)
	}

	return occurrences
}

var masKernels = [][][]rune{
	{
		{'M', ' ', 'M'},
		{' ', 'A', ' '},
		{'S', ' ', 'S'},
	},
	{
		{'S', ' ', 'S'},
		{' ', 'A', ' '},
		{'M', ' ', 'M'},
	},
	{
		{'M', ' ', 'S'},
		{' ', 'A', ' '},
		{'M', ' ', 'S'},
	},
	{
		{'S', ' ', 'M'},
		{' ', 'A', ' '},
		{'S', ' ', 'M'},
	},
}

func SolveB() {
	puzzle := LoadPuzzle()

	// Kernels to match xmas from all directions
	fmt.Println(puzzle.FindKernelsOccurrences(xmasKernels))

	// Second problem
	fmt.Println(puzzle.FindKernelsOccurrences(masKernels))
}

func PrintRuneGrid(grid [][]rune) {
	for _, row := range grid {
		for _, char := range row {
			fmt.Printf("%c", char)
		}
		fmt.Println()
	}
}
