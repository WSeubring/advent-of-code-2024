package day7

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolveA(t *testing.T) {
	SolveA()
}

func TestParseCalculationLine(t *testing.T) {
	line := "10: 2 5"
	calculation := ParseCalculationLine(line)

	assert.Equal(t, 10, calculation.result)
	assert.Equal(t, []int{2, 5}, calculation.values)
}

func TestIsResultPosibleTrueMultiply(t *testing.T) {
	calculation := Calculation{
		result: 10,
		values: []int{2, 5},
	}

	assert.True(t, calculation.IsResultPosible())
}

func TestIsResultPosibleTrueAdd(t *testing.T) {
	calculation := Calculation{
		result: 7,
		values: []int{2, 5},
	}

	assert.True(t, calculation.IsResultPosible())
}

func TestIsResultPosibleFalse(t *testing.T) {
	calculation := Calculation{
		result: 7,
		values: []int{2, 4},
	}

	assert.False(t, calculation.IsResultPosible())
}

func TestIsResultPosibleMoreNumbers(t *testing.T) {
	calculation := Calculation{
		result: 22,
		values: []int{5, 4, 5, 6, 2},
	}

	assert.True(t, calculation.IsResultPosible())
}
