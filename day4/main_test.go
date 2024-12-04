package day4

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wseubring/aoc2024/utils"
)

func TestLoadPuzzle(t *testing.T) {
	puzzle := LoadPuzzle()
	fmt.Println(puzzle)
}

func TestPuzzleFindWordOccurrencesRight(t *testing.T) {
	test := []string{"test"}
	grid := utils.StringsToRunes(test)
	puzzle := Puzzle(grid)

	count := puzzle.FindWordOccurrences("test")
	assert.Equal(t, 1, count)
}

func TestPuzzleFindWordOccurrencesDown(t *testing.T) {
	test := []string{"t", "e", "s", "t"}
	grid := utils.StringsToRunes(test)
	puzzle := Puzzle(grid)

	count := puzzle.FindWordOccurrences("test")
	assert.Equal(t, 1, count)
}

func TestPuzzleFindWordOccurrencesDownRight(t *testing.T) {
	test := []string{"txxx", "xexx", "xxsx", "xxxt"}
	grid := utils.StringsToRunes(test)
	puzzle := Puzzle(grid)

	count := puzzle.FindWordOccurrences("test")
	assert.Equal(t, 1, count)
}

func TestPuzzleFindWordOccurrencesMultiple(t *testing.T) {
	test := []string{"test", "xexe", "xxss", "xxxt"}
	grid := utils.StringsToRunes(test)
	puzzle := Puzzle(grid)

	count := puzzle.FindWordOccurrences("test")
	assert.Equal(t, 3, count)
}

func TestPuzzleGetStringInDirection(t *testing.T) {
	test := []string{"test", "xexe", "xxss", "xxxt"}
	grid := utils.StringsToRunes(test)
	puzzle := Puzzle(grid)

	word := puzzle.GetStringInDirection(0, 0, 1, 0)
	assert.Equal(t, "test", word)

	word = puzzle.GetStringInDirection(0, 0, 0, 1)
	assert.Equal(t, "txxx", word)

	word = puzzle.GetStringInDirection(0, 0, 1, 1)
	assert.Equal(t, "test", word)

	word = puzzle.GetStringInDirection(0, 0, -1, 1)
	assert.Equal(t, "t", word)
}

func TestSolveA(t *testing.T) {
	SolveA()
}
