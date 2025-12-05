package main

import "fmt"

func MainPart2() {
	// rangeStrings, _, err := parseInputFile("test_inputs/p1_example.txt")
	// rangeStrings, _, err := parseInputFile("test_inputs/self_synth_input.txt")
	rangeStrings, _, err := parseInputFile("input.txt")
	if err != nil {
		panic(err)
	}

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

	validCountSum := 0
	skip := false
	for _, rangePair := range ranges {
		if !skip {
			validCountSum = validCountSum + rangePair[1] - rangePair[0] + 1
			// fmt.Println("Summing", rangePair, "total=", validCountSum)
		}
	}

	fmt.Println("Total Allowed IDs (P2)", validCountSum)
}
