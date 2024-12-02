package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

const MAX_SAFE_STEP_SIZE = 3
const MIN_SAFE_STEP_SIZE = 1

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

		levels := make([]int, len(levelsAsString))
		for i, levelAsString := range levelsAsString {
			level, err := strconv.Atoi(levelAsString)
			if err != nil {
				log.Fatal(err)
			}
			levels[i] = level
		}

		if isReportSave(levels) {
			sum += 1
		}
	}

	log.Println(sum)
}

func isReportSave(levels []int) bool {
	prevLevel := levels[0]
	trend := computeTrend(prevLevel, levels[1])
	if trend == 0 {
		return false
	}

	for _, level := range levels[1:] {
		if !isStepSafe(prevLevel, level) {
			return false
		}

		newTrend := computeTrend(prevLevel, level)
		if newTrend != trend {
			return false
		}

		prevLevel = level
	}
	return true
}

func isStepSafe(prev, current int) bool {
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
