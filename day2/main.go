package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

const (
	MAX_SAFE_STEP_SIZE = 3
	MIN_SAFE_STEP_SIZE = 1
)

func main() {
	path := "day2/input.txt"

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sum := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		levelsAsString := strings.Fields(scanner.Text())
		levels := stringsToInts(levelsAsString)

		ok, _ := isReportSave(levels)

		if ok {
			sum++
		} else {
			for i := 0; i < len(levels); i++ {
				left := levels[:i]
				right := levels[i+1:]
				subsetLevels := append(left, right...)
				ok, _ := isReportSave(subsetLevels)
				if ok {
					sum++
					break
				}
			}
		}
	}

	log.Println(sum)
}

func stringsToInts(strings []string) []int {
	ints := make([]int, len(strings))
	for i, s := range strings {
		ints[i], _ = strconv.Atoi(s)
	}
	return ints
}

func isReportSave(levels []int) (bool, int) {
	prevLevel := levels[0]
	trend := computeTrend(prevLevel, levels[1])

	for i, level := range levels[1:] {
		if !isStepSafe(prevLevel, level, trend) {
			return false, i + 1
		}

		prevLevel = level
	}
	return true, 0
}

func isStepSafe(prev, current, trend int) bool {
	if !isStepSizeSafe(prev, current) {
		return false
	}
	if computeTrend(prev, current) != trend || trend == 0 {
		return false
	}
	return true
}

func isStepSizeSafe(prev, current int) bool {
	step := int(math.Abs(float64(current) - float64(prev)))
	return step >= MIN_SAFE_STEP_SIZE && step <= MAX_SAFE_STEP_SIZE
}

func computeTrend(prev, current int) int {
	return boolToInt(current > prev) - boolToInt(prev > current)
}

func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}
