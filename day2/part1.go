package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func parseInputFile(inputPath string) (allParts []string, err error) {
	file, err := os.Open(inputPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}
		parts := strings.Split(scanner.Text(), ",")
		for _, part := range parts {
			if part == "" {
				continue
			}
			allParts = append(allParts, part)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return allParts, nil
}

func splitPartIntoMinMax(part string) (minStr string, maxStr string) {
	components := strings.Split(part, "-")

	if len(components) != 2 {
		panic("Invalid number of components on input")
	}

	return components[0], components[1]
}

func roundToClosestEvenLength(valStr string, roundUp bool) (roundedVal string) {
	// count length of string
	lenOfStr := len(valStr)
	if roundUp {
		return "1" + strings.Repeat("0", lenOfStr)
	} else {
		return strings.Repeat("9", lenOfStr-1)
	}
}

func inBoundsCheck(repVal int, minVal int, maxVal int) bool {
	if repVal >= minVal && repVal <= maxVal {
		return true
	} else {
		return false
	}
}
func findAndCountRepetitions(minStr string, maxStr string) (repetitions []int) {
	// Use heuristic-based counting.
	// We know that repetitions are only valid if the same number repeats twice
	// Heuristic 1: If both the minVal and maxVal have an odd number of characters, 0.
	minIsEvenLength := len(minStr)%2 == 0
	maxIsEvenLength := len(maxStr)%2 == 0
	if !minIsEvenLength && !maxIsEvenLength {
		return make([]int, 0)
	}

	// Heuristic 2: round to closest val if needed.
	if !minIsEvenLength {
		minStr = roundToClosestEvenLength(minStr, true)
	}
	if !maxIsEvenLength {
		maxStr = roundToClosestEvenLength(maxStr, false)
	}
	minVal, err := strconv.Atoi(minStr)
	if err != nil {
		panic(err)
	}
	maxVal, err := strconv.Atoi(maxStr)
	if err != nil {
		panic(err)
	}

	// Heuristic-based counting.
	// Since we know the length, and we have a guaranteed even length,
	// Loop through possibilities.
	partLeftMinStr := minStr[:len(minStr)/2]
	partLeftMaxStr := maxStr[:len(maxStr)/2]
	partLeftMinVal, err := strconv.Atoi(partLeftMinStr)
	if err != nil {
		panic(err)
	}
	partLeftMaxVal, err := strconv.Atoi(partLeftMaxStr)
	if err != nil {
		panic(err)
	}
	// Loop through possibilities.
	for partLeftVal := partLeftMinVal; partLeftVal <= partLeftMaxVal; partLeftVal++ {
		partLeftStringVal := strconv.Itoa(partLeftVal)
		partVal, err := strconv.Atoi(partLeftStringVal + partLeftStringVal)
		if err != nil {
			panic(err)
		}
		if inBoundsCheck(partVal, minVal, maxVal) {
			repetitions = append(repetitions, partVal)
		}
	}
	return repetitions
}

func sumSlice(toSum []int) int {
	sum := 0
	for _, val := range toSum {
		sum += val
	}
	return sum
}

func MainPart1() {
	// parts, err := parseInputFile("test_inputs/p1_example.txt")
	parts, err := parseInputFile("input.txt")
	if err != nil {
		panic(err)
	}

	repetitionsToSum := make([]int, 0)
	for _, part := range parts {
		minStr, maxStr := splitPartIntoMinMax(part)
		repetitions := findAndCountRepetitions(minStr, maxStr)
		repetitionsToSum = append(repetitionsToSum, repetitions...)
	}
	fmt.Println("Total repetitions (P1):", len(repetitionsToSum))
	fmt.Println("Sum of IDs (P1):", sumSlice(repetitionsToSum))
}
