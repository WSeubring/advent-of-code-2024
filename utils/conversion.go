package utils

import "strconv"

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
	i, err := strconv.Atoi(s)
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
