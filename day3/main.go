package day3

import (
	_ "embed"
	"fmt"
	"regexp"

	"github.com/wseubring/aoc2024/utils"
)

type CorruptedProgram string
type Operation string

//go:embed input.txt
var program CorruptedProgram

func (c CorruptedProgram) Run() int {
	opperations := c.FindAllOpperation()
	sum := 0
	for _, opperation := range opperations {
		sum += opperation.Run()
	}

	return sum
}

// Regex find all that match the pattern: mul(1,2)
func (c CorruptedProgram) FindAllOpperation() []Operation {
	regex := `mul\(\d+,\d+\)`
	re := regexp.MustCompile(regex)

	opperations := []Operation{}
	for _, match := range re.FindAllString(string(c), -1) {
		opperations = append(opperations, Operation(match))
	}

	return opperations
}

func (o Operation) Run() int {
	values := o.GetValues()
	return values[0] * values[1]
}

func (o Operation) GetValues() []int {
	regex := `\d+`
	re := regexp.MustCompile(regex)

	values := []int{}
	for _, match := range re.FindAllString(string(o), -1) {
		num, _ := utils.StringToInt(match)
		values = append(values, num)
	}

	return values
}

func (c *CorruptedProgram) ReducToDoParts() {
	regex := `(do\(\)|^)(.*?)(don't\(\)|$)`
	re := regexp.MustCompile(regex)

	newProgram := ""
	for _, match := range re.FindAllString(string(*c), -1) {
		newProgram += match
	}

	*c = CorruptedProgram(newProgram)
}

func SolveA() {
	fmt.Println(program.Run())
}

func SolveB() {
	program.ReducToDoParts()
	fmt.Println(program.Run())
}
