package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func parseInputFile(
	inputPath string,
) (startCoords [2]int, splitterCoords [][2]int, maxX int, maxY int) {
	file, err := os.Open(inputPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lineIndex := 0
	for scanner.Scan() {
		// Terminate on empty line
		if scanner.Text() == "" {
			break
		}
		// Line 1: Find the start coordinate.
		if lineIndex == 0 {
			for i, charRune := range scanner.Text() {
				if string(charRune) == "S" {
					startCoords = [2]int{lineIndex, i}
					break
				}
			}
			lineIndex = lineIndex + 1
			// Also assume a square input, and keep track of the max line length
			maxX = len(scanner.Text())
			continue
		}
		// For every next line, find the splitters
		for i, charRune := range scanner.Text() {
			if string(charRune) == "^" {
				splitterCoords = append(splitterCoords, [2]int{lineIndex, i})
			}
		}
		lineIndex = lineIndex + 1
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return startCoords, splitterCoords, maxX, lineIndex
}

func sliceContainsCoords(coordsToCheck [2]int, splitterCoords [][2]int) bool {
	for _, coord := range splitterCoords {
		if coord[0] == coordsToCheck[0] && coord[1] == coordsToCheck[1] {
			return true
		}
	}
	return false
}

func propagateAndCountSplits(
	currentCoord [2]int,
	foundSplits map[[2]int]bool,
	splitterCoords [][2]int,
	maxX int,
	maxY int,
) {
	currentY := currentCoord[0]
	currentX := currentCoord[1]
	// Termination conditions.
	// Case 1: reached max,
	if currentY >= maxY {
		// fmt.Println("Reached maxY, terminating.")
		return
	}
	// Case 2: X is invalid.
	if currentX >= maxX || currentX < 0 {
		// fmt.Println("Hit invalid X", currentCoord, "terminating.")
		return
	}
	if sliceContainsCoords(currentCoord, splitterCoords) {
		// Early terminate
		_, exists := foundSplits[currentCoord]
		if exists {
			return
		}
		foundSplits[currentCoord] = true
		propagateAndCountSplits(
			[2]int{currentY, currentX - 1},
			foundSplits,
			splitterCoords,
			maxX,
			maxY,
		)
		propagateAndCountSplits(
			[2]int{currentY, currentX + 1},
			foundSplits,
			splitterCoords,
			maxX,
			maxY,
		)
		return
	}
	propagateAndCountSplits(
		[2]int{currentY + 1, currentX},
		foundSplits,
		splitterCoords,
		maxX,
		maxY,
	)
}

func MainPart1() {
	// startCoord, splitterCoords, maxX, maxY := parseInputFile("test_inputs/p1_example.txt")
	startCoord, splitterCoords, maxX, maxY := parseInputFile("input.txt")
	// fmt.Println("Start coordinates", startCoord)
	// fmt.Println("Splitter coordinates", splitterCoords)
	// fmt.Println("Max X:", maxX)
	// fmt.Println("Max Y:", maxY)
	foundSplits := make(map[[2]int]bool)
	propagateAndCountSplits(startCoord, foundSplits, splitterCoords, maxX, maxY)
	fmt.Println(
		"Split count (P1):", len(foundSplits),
	)
}
