package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func recursiveFind(combination []int, sliceToSearch []int, lengthOfCombination int, searchNum int, sourceSliceToSearch []int) ([]int, bool) {
	if len(combination) == lengthOfCombination {
		return combination, true
	}
	if len(sliceToSearch) == 0 {
		return make([]int, 0), false
	}
	if searchNum < 0 {
		return make([]int, 0), false
	}

	index := slices.Index(sliceToSearch, searchNum)

	// Value found
	if index != -1 {
		// Here, only accept the value if the slice is still long enough.
		if len(sliceToSearch)-index < lengthOfCombination-len(combination) {
			return recursiveFind(combination, sliceToSearch, lengthOfCombination, searchNum-1, sourceSliceToSearch)
		} else {
			sliceToSearch = sliceToSearch[index+1:]
			combination = append(combination, searchNum)
			return recursiveFind(combination, sliceToSearch, lengthOfCombination, 9, sourceSliceToSearch)
		}
	} else {
		return recursiveFind(combination, sliceToSearch, lengthOfCombination, searchNum-1, sourceSliceToSearch)
	}
}

func findMaxCombinationP2(line string, lengthOfCombination int) int {
	lineChars := strings.Split(line, "")
	lineInts := make([]int, 0)
	for _, lineChar := range lineChars {
		lineInt, err := strconv.Atoi(lineChar)
		if err != nil {
			panic(err)
		}
		lineInts = append(lineInts, lineInt)
	}

	// Find max combination.
	// Just go by heuristic.
	// Since it's a 2-number value every time, just check if 9 is available.
	// Then check if it is the rightmost number, if not, find 8, and so on.
	startNum := 9
	combinationFound := false
	proposal := make([]int, 0)
	for !combinationFound {
		proposal, combinationFound = recursiveFind(proposal, lineInts, lengthOfCombination, startNum, lineInts)
		startNum = startNum - 1
	}

	// Composing combined number.
	combinedNumberStr := ""
	for _, val := range proposal {
		valStr := strconv.Itoa(val)
		combinedNumberStr = combinedNumberStr + valStr
	}
	combinedNumber, err := strconv.Atoi(combinedNumberStr)
	if err != nil {
		panic(err)
	}

	return combinedNumber
}
func MainPart2() {
	// parts, err := parseInputFile("test_inputs/p1_example.txt")
	parts, err := parseInputFile("input.txt")
	if err != nil {
		panic(err)
	}

	intsToSum := make([]int, 0)
	for _, part := range parts {
		// fmt.Println(part)
		intsToSum = append(intsToSum, findMaxCombinationP2(part, 12))
	}
	fmt.Println("Sum of joltage (P2):", sumSlice(intsToSum))
}
