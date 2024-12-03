package day2

import (
	"bufio"
	"log"
	"os"
	"strings"

	"github.com/wseubring/aoc2024/utils"
)

const (
	MAX_SAFE_STEP_SIZE = 3
	MIN_SAFE_STEP_SIZE = 1
)

func Solve() {
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
		levels, err := utils.StringsToInts(levelsAsString)
		if err != nil {
			log.Fatal(err)
		}

		if isReportSave(levels) {
			sum++
		} else {
			for i := 0; i < len(levels); i++ {
				subsetLevels := utils.RemoveIndexFromArr(levels, i)
				if isReportSave(subsetLevels) {
					sum++
					break
				}
			}
		}
	}

	log.Println(sum)
}

func isReportSave(levels []int) bool {
	prevLevel := levels[0]
	trend := computeTrend(prevLevel, levels[1])

	for _, level := range levels[1:] {
		if !isStepSafe(prevLevel, level, trend) {
			return false
		}

		prevLevel = level
	}
	return true
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
	step := utils.AbsInt(prev - current)
	return step >= MIN_SAFE_STEP_SIZE && step <= MAX_SAFE_STEP_SIZE
}

func computeTrend(prev, current int) int {
	return utils.BoolToInt(current > prev) - utils.BoolToInt(prev > current)
}
