package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func roundToClosestEvenLengthP2(valStr string, roundUp bool, repCount int) (roundedVal string) {
	lenOfStr := len(valStr)
	targetLength := 0
	if roundUp {
		remainder := lenOfStr % repCount
		targetLength = lenOfStr + (repCount - remainder)
		return "1" + strings.Repeat("0", targetLength-1)
	} else {
		targetLength = (lenOfStr / repCount) * repCount
		return strings.Repeat("9", targetLength)
	}
}

func findAndCountRepetitionsP2(minStr string, maxStr string, repCount int) (repetitions []int) {
	// Use heuristic-based counting.
	// We know that repetitions are only valid if the same number repeats twice
	// Heuristic 1: If both the minVal and maxVal have an odd number of characters, 0.
	minIsEvenLength := len(minStr)%repCount == 0
	maxIsEvenLength := len(maxStr)%repCount == 0
	if !minIsEvenLength && !maxIsEvenLength {
		return make([]int, 0)
	}

	// Heuristic 2: round to closest val if needed.
	if !minIsEvenLength {
		minStr = roundToClosestEvenLengthP2(minStr, true, repCount)
	}
	if !maxIsEvenLength {
		maxStr = roundToClosestEvenLengthP2(maxStr, false, repCount)
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
	partLeftMinStr := minStr[:len(minStr)/repCount]
	partLeftMaxStr := maxStr[:len(maxStr)/repCount]
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
		partVal, err := strconv.Atoi(strings.Repeat(partLeftStringVal, repCount))
		if err != nil {
			panic(err)
		}
		if inBoundsCheck(partVal, minVal, maxVal) {
			repetitions = append(repetitions, partVal)
		}
	}
	return repetitions
}

func MainPart2() {
	// parts, err := parseInputFile("test_inputs/p1_example.txt")
	parts, err := parseInputFile("input.txt")
	if err != nil {
		panic(err)
	}

	repetitionsToSum := make([]int, 0)
	for _, part := range parts {
		minStr, maxStr := splitPartIntoMinMax(part)
		currentPartRepetitions := make([]int, 0)
		for i := 2; i <= len(maxStr); i++ {
			repetitions := findAndCountRepetitionsP2(minStr, maxStr, i)
			for _, repetition := range repetitions {
				if !slices.Contains(currentPartRepetitions, repetition) {
					currentPartRepetitions = append(currentPartRepetitions, repetition)
				}
			}
		}
		repetitionsToSum = append(repetitionsToSum, currentPartRepetitions...)
	}
	fmt.Println("Total repetitions (P2):", len(repetitionsToSum))
	fmt.Println("Sum of IDs (P2):", sumSlice(repetitionsToSum))
}
