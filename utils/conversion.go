package utils

import (
	"strconv"
	"strings"
)

func StringsToInts(strings []string) ([]int, error) {
	ints := make([]int, len(strings))
	for i, s := range strings {
		result, err := StringToInt(s)
		if err != nil {
			return nil, err
		}

		ints[i] = result
	}
	return ints, nil
}

func StringToInt(s string) (int, error) {
	i, err := strconv.Atoi(strings.TrimSpace(s))
	if err != nil {
		return 0, err
	}
	return i, nil
}

func BoolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

func StringToRunes(s string) []rune {
	runes := make([]rune, len(s))
	for i, r := range s {
		runes[i] = r
	}
	return runes
}

func StringsToRunes(strings []string) [][]rune {
	runes := make([][]rune, len(strings))
	for i, s := range strings {
		runes[i] = StringToRunes(s)
	}
	return runes
}

func RunesToString(runes [][]rune) string {
	var result string
	for _, row := range runes {
		result += string(row) + "\n"
	}
	return result
}
