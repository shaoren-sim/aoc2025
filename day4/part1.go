package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

func getMultidimensionalRepresentation(rows []string) [][]string {
	multidim := make([][]string, len(rows))
	for rowInd, row := range rows {
		rowVals := make([]string, len(row))
		for colInd, rowValRune := range row {
			rowVals[colInd] = string(rowValRune)
		}
		multidim[rowInd] = rowVals
	}
	return multidim
}

func getSliceOfAdjacentCoords(rowInd int, colInd int) [][2]int {
	adjacentCoords := make([][2]int, 0)

	offsets := []int{-1, 0, 1}
	for _, rowOffset := range offsets {
		rowIndToAdd := rowInd + rowOffset
		if rowIndToAdd < 0 {
			continue
		}
		for _, colOffset := range offsets {
			// Skip the central pixel.
			if rowOffset == 0 && colOffset == 0 {
				continue
			}
			colIndToAdd := colInd + colOffset
			if colIndToAdd < 0 {
				continue
			} else {
				adjacentCoords = append(adjacentCoords, [2]int{rowIndToAdd, colIndToAdd})
			}
		}
	}

	return adjacentCoords
}

func countAdjacentHits(multidim [][]string, rowInd int, colInd int) (adjacentHits int) {
	// Given the coordinates, find adjacent values.
	numRows := len(multidim)
	numCols := len(multidim[0]) // Here. assuming all cols are same length
	adjacentCoords := getSliceOfAdjacentCoords(rowInd, colInd)

	for _, coordCandidate := range adjacentCoords {
		if coordCandidate[0] >= numRows {
			continue
		}
		if coordCandidate[1] >= numCols {
			continue
		}
		if multidim[coordCandidate[0]][coordCandidate[1]] == "@" {
			adjacentHits += 1
		}
	}
	return adjacentHits
}

func MainPart1() {
	// parts, err := parseInputFile("test_inputs/p1_example.txt")
	parts, err := parseInputFile("input.txt")
	if err != nil {
		panic(err)
	}

	multidim := getMultidimensionalRepresentation(parts)

	// Effectively copying the original map into the visualizer.
	visualization := make([][]string, 0)
	for _, row := range multidim {
		vizRow := make([]string, len(row))
		copy(vizRow, row)
		visualization = append(visualization, vizRow)
	}

	validCount := 0
	for rowInd, row := range multidim {
		for colInd, rowVal := range row {
			if rowVal == "@" {
				adjacentHits := countAdjacentHits(multidim, rowInd, colInd)
				// fmt.Println("Found @ at", rowInd, colInd, "with", adjacentHits, "adjacent @s")
				if adjacentHits < 4 {
					// fmt.Println("Found valid at", rowInd, colInd)
					validCount += 1
					visualization[rowInd][colInd] = "x"
				}
			}
		}
	}

	// fmt.Println("Viz")
	// for _, row := range visualization {
	// 	fmt.Println(row)
	// }
	// fmt.Println("==endviz")
	fmt.Println("Total counts (P1):", validCount)
}
