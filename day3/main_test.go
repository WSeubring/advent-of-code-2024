package day3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var example = "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"
var example2 = "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"

func TestFindAllOperations(t *testing.T) {
	operations := CorruptedProgram(example).FindAllOpperation()
	if len(operations) != 4 {
		assert.Equal(t, 4, len(operations))
	}
}

func TestRunOperation(t *testing.T) {
	operation := Operation("mul(2,4)")
	result := operation.Run()
	assert.Equal(t, 8, result)
}

func TestGetValues(t *testing.T) {
	operation := Operation("mul(2,4)")
	values := operation.GetValues()
	assert.Equal(t, 2, values[0])
	assert.Equal(t, 4, values[1])
}

func TestRun(t *testing.T) {
	result := CorruptedProgram(example).Run()
	assert.Equal(t, 161, result)
}

func TestReducedToParts(t *testing.T) {
	c := CorruptedProgram(example2)
	c.ReducToDoParts()

	assert.Equal(t, "xmul(2,4)&mul[3,7]!^don't()do()?mul(8,5))", string(c))
}

func TestRunReduced(t *testing.T) {
	c := CorruptedProgram(example2)
	c.ReducToDoParts()
	result := c.Run()
	assert.Equal(t, 48, result)
}
