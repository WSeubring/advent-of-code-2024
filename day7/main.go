package day7

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/wseubring/aoc2024/utils"
)

//go:embed input.txt
var input string

func SolveA() {
	sum := 0

	lines := strings.Split(input, "\n")
	for _, line := range lines {
		calculation := ParseCalculationLine(line)
		if calculation.IsResultPosible() {
			sum += calculation.result
		}
	}

	fmt.Println("Sum of possible results:", sum)
}

type Calculation struct {
	result int
	values []int
}

func ParseCalculationLine(line string) Calculation {
	parts := strings.Split(line, ": ")

	result, err := utils.StringToInt(parts[0])
	if err != nil {
		panic(err)
	}

	values, err := utils.StringsToInts(strings.Split(parts[1], " "))
	if err != nil {
		panic(err)
	}

	return Calculation{
		result: result,
		values: values,
	}
}

func (c Calculation) IsResultPosible() bool {
	firstValue := c.values[0]
	results := c.findResultsRec(c.values[1:], firstValue)
	for _, r := range results {
		if r == c.result {
			return true
		}
	}

	return false
}

func (c Calculation) findResultsRec(remaining []int, sum int) []int {
	if len(remaining) == 0 {
		return []int{sum}
	}

	nextValue := remaining[0]
	remaining = remaining[1:]

	stringConcat := fmt.Sprintf("%d%d", sum, nextValue)
	stringConcatAsNumber, err := utils.StringToInt(stringConcat)
	if err != nil {
		panic(err)
	}
	// Part 1
	result := append(
		c.findResultsRec(remaining, sum+nextValue),
		c.findResultsRec(remaining, sum*nextValue)...,
	)
	// Part 2
	result = append(result, c.findResultsRec(remaining, stringConcatAsNumber)...)

	return result
}
