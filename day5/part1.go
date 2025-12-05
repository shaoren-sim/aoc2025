package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func parseInputFile(inputPath string) (rangeStrings []string, idsToCheck []string, err error) {
	file, err := os.Open(inputPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	newLineCount := 0
	for scanner.Scan() {
		if scanner.Text() == "" {
			newLineCount = newLineCount + 1
			continue
		}
		if newLineCount == 0 {
			rangeStrings = append(rangeStrings, scanner.Text())
		} else if newLineCount == 1 {
			idsToCheck = append(idsToCheck, scanner.Text())
		} else {
			break
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}

	return rangeStrings, idsToCheck, nil
}

func generateInitialRanges(rangeStrings []string) [][2]int {
	ranges := make([][2]int, 0)
	for _, rangeString := range rangeStrings {
		rangePartStr := strings.Split(rangeString, "-")
		rangePartLeft, err := strconv.Atoi(rangePartStr[0])
		if err != nil {
			panic(err)
		}
		rangePartRight, err := strconv.Atoi(rangePartStr[1])
		if err != nil {
			panic(err)
		}
		ranges = append(ranges, [2]int{rangePartLeft, rangePartRight})
	}
	return ranges
}

func compactRanges(uncompactedRanges [][2]int) ([][2]int, bool) {
	compactedRanges := make([][2]int, 0)
	compactionFound := false
	for _, rangePartCandidate := range uncompactedRanges {
		rangePartLeft := rangePartCandidate[0]
		rangePartRight := rangePartCandidate[1]
		if rangePartLeft > rangePartRight {
			fmt.Println(rangePartCandidate)
			panic("Left is > right part!")
		}
		if len(compactedRanges) == 0 {
			compactedRanges = append(compactedRanges, [2]int{rangePartLeft, rangePartRight})
			continue
		}
		overlapsFound := false
		for indExisting, rangeExisting := range compactedRanges {
			// fmt.Println("Evaluating", rangePartCandidate, "against", rangeExisting)
			rangeExistingLeft := rangeExisting[0]
			rangeExistingRight := rangeExisting[1]
			if rangePartLeft >= rangeExistingLeft && rangePartLeft <= rangeExistingRight {
				overlapsFound = true
				// Case 1: Extend rightwards
				if rangePartRight >= rangeExistingRight {
					compactedRanges[indExisting] = [2]int{rangeExistingLeft, rangePartRight}
					compactionFound = true
				}
				continue
			} else if rangePartRight <= rangeExistingRight && rangePartRight >= rangeExistingLeft {
				// Case 2: Extend leftwards
				overlapsFound = true
				if rangePartLeft <= rangeExistingLeft {
					compactedRanges[indExisting] = [2]int{rangePartLeft, rangeExistingRight}
					compactionFound = true
				}
				continue
			} else if rangePartLeft <= rangeExistingLeft && rangePartRight >= rangeExistingRight {
				// Case 3: The range encompasses the existing range entirely.
				compactedRanges[indExisting] = [2]int{rangePartLeft, rangePartRight}
				compactionFound = true
			} else {
				// Case 4: Equal left and right, i.e. single-int range
				if rangeExistingLeft == rangeExistingRight {
					// Only compact if the new candidate range encompasses this single one.
					if rangePartLeft <= rangeExistingLeft && rangePartRight >= rangeExistingRight {
						compactedRanges[indExisting] = [2]int{rangePartLeft, rangePartRight}
					}
				}
			}
		}
		if !overlapsFound {
			// fmt.Println("No overlaps, adding new entry", [2]int{rangePartLeft, rangePartRight})
			compactedRanges = append(compactedRanges, [2]int{rangePartLeft, rangePartRight})
		}

	}
	return compactedRanges, compactionFound
}

func isIdSpoilt(idString string, ranges [][2]int) bool {
	for _, rangePair := range ranges {
		// fmt.Println("Check id", idString, "in", rangePair)
		idVal, err := strconv.Atoi(idString)
		if err != nil {
			panic(err)
		}
		if idVal >= rangePair[0] && idVal <= rangePair[1] {
			return true
		}
	}
	return false
}

func MainPart1() {
	// rangeStrings, idStringsToCheck, err := parseInputFile("test_inputs/p1_example.txt")
	rangeStrings, idStringsToCheck, err := parseInputFile("input.txt")
	if err != nil {
		panic(err)
	}
	// fmt.Println(rangeStrings)
	// fmt.Println(idStringsToCheck)

	rangesInitial := generateInitialRanges(rangeStrings)
	// Copy the initial slice.
	ranges := make([][2]int, len(rangesInitial))
	copy(ranges, rangesInitial)
	for {
		rangesNew, compactionFound := compactRanges(ranges)
		ranges = rangesNew
		if !compactionFound {
			break
		}
	}

	// Check IDs and count spoiled
	spoiltCount := 0
	for _, idString := range idStringsToCheck {
		if isIdSpoilt(idString, ranges) {
			spoiltCount = spoiltCount + 1
		}
	}
	fmt.Println("Spoilt count (P1):", spoiltCount)
}
