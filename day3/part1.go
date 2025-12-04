package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
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
		parts := strings.SplitSeq(scanner.Text(), ",")
		for part := range parts {
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

func sumSlice(toSum []int) int {
	sum := 0
	for _, val := range toSum {
		sum += val
	}
	return sum
}

func findMaxCombination(line string) int {
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
	valuesInt := make([]int, 2)
	for rep := range 2 {
		numToCheck := 9
		matchFound := false
		for !matchFound {
			index := slices.Index(lineInts, numToCheck)
			// Match found
			if index != -1 {
				// If this is the last value, it's okay to be at the end
				if rep == 1 {
					valuesInt[rep] = numToCheck
					lineInts = slices.Delete(lineInts, index, index+1)
					matchFound = true
				} else {
					if index != (len(lineInts) - 1) {
						valuesInt[rep] = numToCheck
						lineInts = slices.Delete(lineInts, 0, index+1)
						matchFound = true
					} else {
						numToCheck -= 1
						if numToCheck < 0 {
							break
						}
						continue
					}
				}
			} else {
				numToCheck -= 1
				if numToCheck < 0 {
					break
				}
				continue
			}
		}
	}

	// Composing combined number.
	combinedNumberStr := ""
	for _, val := range valuesInt {
		valStr := strconv.Itoa(val)
		combinedNumberStr = combinedNumberStr + valStr
	}
	combinedNumber, err := strconv.Atoi(combinedNumberStr)
	if err != nil {
		panic(err)
	}

	return combinedNumber
}

func MainPart1() {
	// parts, err := parseInputFile("test_inputs/p1_example.txt")
	parts, err := parseInputFile("input.txt")
	if err != nil {
		panic(err)
	}

	intsToSum := make([]int, 0)
	for _, part := range parts {
		// fmt.Println(part)
		intsToSum = append(intsToSum, findMaxCombination(part))
	}
	fmt.Println("Sum of joltage (P1):", sumSlice(intsToSum))
}
