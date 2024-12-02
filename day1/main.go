package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

const inputFile = "day1/input.txt"

func main() {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var listA []int
	var listB []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Fields(line)
		a, err := strconv.Atoi(parts[0])
		if err != nil {
			log.Fatal(err)
		}
		listA = append(listA, a)

		b, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Fatal(err)
		}
		listB = append(listB, b)
	}

	fmt.Println(solveA(listA, listB))
	fmt.Println(solveB(listA, listB))
}

func solveA(listA, listB []int) int {
	slices.Sort(listA)
	slices.Sort(listB)

	sum := 0
	for i := 0; i < len(listA); i++ {
		a := listA[i]
		b := listB[i]

		sum += int(math.Abs(float64(a) - float64(b)))
	}

	return sum
}

func solveB(listA, listB []int) int {
	mapA := make(map[int]int)
	mapB := make(map[int]int)

	for i := 0; i < len(listA); i++ {
		mapA[listA[i]]++
		mapB[listB[i]]++
	}

	sum := 0
	for k, v := range mapA {
		smallestCount := int(math.Min(float64(v), float64(mapB[k])))
		sum += k * mapB[k] * smallestCount
	}

	return sum
}
