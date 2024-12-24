package day9

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/wseubring/aoc2024/utils"
)

type DiskMap struct {
	DenseFormat string
}

//go:embed input.txt
var input string

func LoadInputDiskMap(input string) DiskMap {
	return DiskMap{DenseFormat: input}
}

var mapIndexToFileId = map[int64]int64{}

func (d DiskMap) Decompress() string {
	idCounter := int64(0)
	size := int64(0)
	result := ""
	for i := int64(0); i < int64(len(d.DenseFormat)); i++ {
		value := int64(d.DenseFormat[i] - '0')
		if i%2 == 0 {
			result += strings.Repeat("#", int(value))
			for j := int64(0); j < value; j++ {
				mapIndexToFileId[size+j] = idCounter
			}
			size += value
			idCounter++
		} else {
			result += strings.Repeat(".", int(value))
			for j := int64(0); j < value; j++ {
				mapIndexToFileId[size+j] = 0
			}
			size += value
		}
	}
	return result
}

func ComputeSum(disk string) int64 {
	sum := int64(0)
	for key, value := range mapIndexToFileId {
		sum += key * value
	}
	return sum
}

func SolveA() {
	diskMap := LoadInputDiskMap(input)

	disk := []rune(diskMap.Decompress())
	fmt.Println(mapIndexToFileId)
	for j := int64(len(disk) - 1); j > 0; j-- {
		for i := int64(0); i < j; i++ {
			to := disk[i]
			from := disk[j]

			if from == '.' {
				continue
			}

			if to != '.' {
				continue
			}

			utils.Swap(disk, int(i), int(j))
			// Flip the map values
			mapIndexToFileId[i] = mapIndexToFileId[j]
			mapIndexToFileId[j] = 0
		}
	}
	result := string(disk)

	fmt.Println(result)
	fmt.Println(ComputeSum(result))
}
