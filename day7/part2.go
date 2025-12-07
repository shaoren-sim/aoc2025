package main

import "fmt"

func propagate(
	currentCoord [2]int,
	splitTimelinesMap map[[2]int]int,
	splitterCoords [][2]int,
	maxX int,
	maxY int,
) int {
	currentY := currentCoord[0]
	currentX := currentCoord[1]
	// Termination conditions.
	// Case 1: reached max,
	if currentY >= maxY {
		splitTimelinesMap[currentCoord] += 1
		return 1
	}
	// Case 2: X is invalid.
	if currentX >= maxX || currentX < 0 {
		return 0
	}
	for {
		currentY = currentY + 1
		if currentY >= maxY {
			// fmt.Println("Exceed maxY")
			return 1
		}
		if sliceContainsCoords([2]int{currentY, currentX}, splitterCoords) {
			newCoord := [2]int{currentY, currentX}
			_, exists := splitTimelinesMap[newCoord]
			if exists {
				return splitTimelinesMap[newCoord]
			} else {
				splitTimelinesMap[newCoord] += propagate(
					[2]int{currentY, currentX - 1},
					splitTimelinesMap,
					splitterCoords,
					maxX,
					maxY,
				)
				splitTimelinesMap[newCoord] += propagate(
					[2]int{currentY, currentX + 1},
					splitTimelinesMap,
					splitterCoords,
					maxX,
					maxY,
				)
			}
			// Return the value here so the memoized sums propagate upwards.
			return splitTimelinesMap[newCoord]
		}
	}
}

func MainPart2() {
	// startCoord, splitterCoords, maxX, maxY := parseInputFile("test_inputs/p1_example.txt")
	// startCoord, splitterCoords, maxX, maxY := parseInputFile("test_inputs/self_synth_input.txt")
	startCoord, splitterCoords, maxX, maxY := parseInputFile("input.txt")
	splitTimelines := make(map[[2]int]int)
	propagate(
		startCoord,
		splitTimelines,
		splitterCoords,
		maxX,
		maxY,
	)

	// HACK: Just try to find the first split point.
	// Works because there is only a single start point.
	startY := startCoord[0]
	startX := startCoord[1]
	firstSplitY := startY
	for {
		firstSplitY = firstSplitY + 1
		if sliceContainsCoords([2]int{firstSplitY, startX}, splitterCoords) {
			break
		}
	}

	fmt.Println(
		"Timeline count (P2):", splitTimelines[[2]int{firstSplitY, startX}],
	)
}
