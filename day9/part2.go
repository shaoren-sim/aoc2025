package main

import (
	"fmt"
	"math"
	"slices"
)

func _isRectangleComplete(positions [][2]int, filledPositions [][2]int, a [2]int, b [2]int) bool {
	if len(a) != len(b) {
		panic("a and b need to have same length for distance calculation")
	}
	xA := a[0]
	yA := a[1]
	xB := b[0]
	yB := b[1]

	if xA == xB || yA == yB {
		fmt.Println("Handling comparison of", a, "and", b, "and they are on the same axis, i.e COMPLETE.")
		return true
	}

	// Early return if there are in-between values, i.e. this rectangle is invalid.
	if xB > xA {
		for x := xA + 1; x < xB; x++ {
			// Case 2: Checking if all intermediate values are filled.
			// Done crudely, looping over each of the corners.
			// We do not need to assume the same-horizontal/vertical case due to early return.
			candidatePositionA := [2]int{x, yA}
			if slices.Contains(positions, candidatePositionA) || slices.Contains(filledPositions, candidatePositionA) {
				fmt.Println("Handling comparison of", a, "and", b, "and it is NOT COMPLETE with unfilled coord", candidatePositionA)
				return false
			}
			candidatePositionB := [2]int{x, yB}
			if slices.Contains(positions, candidatePositionB) || slices.Contains(filledPositions, candidatePositionB) {
				fmt.Println("Handling comparison of", a, "and", b, "and it is NOT COMPLETE with unfilled coord", candidatePositionB)
				return false
			}
			// Case 1: Staircase.
			if yB > yA {
				for y := yA + 1; y < yB; y++ {
					if slices.Contains(positions, [2]int{x, y}) {
						fmt.Println("Handling comparison of", a, "and", b, "and it is NOT COMPLETE.")
						return false
					}
				}
			} else {
				for y := yB + 1; y < yA; y++ {
					if slices.Contains(positions, [2]int{x, y}) {
						fmt.Println("Handling comparison of", a, "and", b, "and it is NOT COMPLETE.")
						return false
					}
				}
			}
		}
	} else {
		for x := xB + 1; x < xA; x++ {
			// Case 2: Checking if all intermediate values are filled.
			// Done crudely, looping over each of the corners.
			// We do not need to assume the same-horizontal/vertical case due to early return.
			candidatePositionA := [2]int{x, yA}
			if slices.Contains(positions, candidatePositionA) || slices.Contains(filledPositions, candidatePositionA) {
				fmt.Println("Handling comparison of", a, "and", b, "and it is NOT COMPLETE with unfilled coord", candidatePositionA)
				return false
			}
			candidatePositionB := [2]int{x, yB}
			if slices.Contains(positions, candidatePositionB) || slices.Contains(filledPositions, candidatePositionB) {
				fmt.Println("Handling comparison of", a, "and", b, "and it is NOT COMPLETE with unfilled coord", candidatePositionB)
				return false
			}
			if yB > yA {
				for y := yA + 1; y < yB; y++ {
					candidatePositionA := [2]int{xA, y}
					if slices.Contains(positions, candidatePositionA) || slices.Contains(filledPositions, candidatePositionA) {
						fmt.Println("Handling comparison of", a, "and", b, "and it is NOT COMPLETE with unfilled coord", candidatePositionA)
						return false
					}
					candidatePositionB := [2]int{xB, y}
					if slices.Contains(positions, candidatePositionB) || slices.Contains(filledPositions, candidatePositionB) {
						fmt.Println("Handling comparison of", a, "and", b, "and it is NOT COMPLETE with unfilled coord", candidatePositionB)
						return false
					}
					if slices.Contains(positions, [2]int{x, y}) {
						fmt.Println("Handling comparison of", a, "and", b, "and it is NOT COMPLETE.")
						return false
					}
				}
			} else {
				for y := yB + 1; y < yA; y++ {
					candidatePositionA := [2]int{xA, y}
					if slices.Contains(positions, candidatePositionA) || slices.Contains(filledPositions, candidatePositionA) {
						fmt.Println("Handling comparison of", a, "and", b, "and it is NOT COMPLETE with unfilled coord", candidatePositionA)
						return false
					}
					candidatePositionB := [2]int{xB, y}
					if slices.Contains(positions, candidatePositionB) || slices.Contains(filledPositions, candidatePositionB) {
						fmt.Println("Handling comparison of", a, "and", b, "and it is NOT COMPLETE with unfilled coord", candidatePositionB)
						return false
					}
					if slices.Contains(positions, [2]int{x, y}) {
						fmt.Println("Handling comparison of", a, "and", b, "and it is NOT COMPLETE.")
						return false
					}
				}
			}
		}
	}

	fmt.Println("Handling comparison of", a, "and", b, "and it is complete.")
	return true
}

func isRectangleComplete(positions map[[2]int]bool, a [2]int, b [2]int) bool {
	if len(a) != len(b) {
		panic("a and b need to have same length for distance calculation")
	}
	xA := a[0]
	yA := a[1]
	xB := b[0]
	yB := b[1]

	if xA == xB || yA == yB {
		fmt.Println("Handling comparison of", a, "and", b, "and they are on the same axis, i.e COMPLETE.")
		return true
	}

	if xB > xA {
		for x := xA; x < xB; x++ {
			if yB > yA {
				for y := yA; y < yB; y++ {
					_, exists := positions[[2]int{x, y}]
					if !exists {
						fmt.Println(x, y)
						fmt.Println("Handling comparison of", a, "and", b, "and it is NOT COMPLETE.")
						return false
					}
				}
			} else {
				for y := yB; y < yA; y++ {
					_, exists := positions[[2]int{x, y}]
					if !exists {
						fmt.Println("Handling comparison of", a, "and", b, "and it is NOT COMPLETE.")
						return false
					}
				}
			}
		}
	} else {
		for x := xB; x < xA; x++ {
			if yB > yA {
				for y := yA; y < yB; y++ {
					_, exists := positions[[2]int{x, y}]
					if !exists {
						fmt.Println("Handling comparison of", a, "and", b, "and it is NOT COMPLETE.")
						return false
					}
				}
			} else {
				for y := yB; y < yA; y++ {
					_, exists := positions[[2]int{x, y}]
					if !exists {
						fmt.Println("Handling comparison of", a, "and", b, "and it is NOT COMPLETE.")
						return false
					}
				}
			}
		}
	}

	fmt.Println("Handling comparison of", a, "and", b, "and it is complete.")
	return true
}
func createFilledPositions(redCoordMap map[[2]int]bool) map[[2]int]bool {
	filledCoordMap := make(map[[2]int]bool)
	for positionA := range redCoordMap {
		for positionB := range redCoordMap {
			if positionA == positionB {
				continue
			}
			xA := positionA[0]
			yA := positionA[1]
			xB := positionB[0]
			yB := positionB[1]

			if xA == xB {
				// fill along y-axis.
				if yB > yA {
					for y := yA + 1; y < yB; y++ {
						candidatePos := [2]int{xA, y}
						_, exists := filledCoordMap[candidatePos]
						if !exists {
							filledCoordMap[candidatePos] = true
						}
					}
				} else {
					for y := yB + 1; y < yA; y++ {
						candidatePos := [2]int{xA, y}
						filledCoordMap[candidatePos] = true
					}
				}
			} else if yA == yB {
				// fill along x-axis.
				if xB > xA {
					for x := xA + 1; x < xB; x++ {
						candidatePos := [2]int{x, yA}
						filledCoordMap[candidatePos] = true
					}
				} else {
					for x := xB + 1; x < xA; x++ {
						candidatePos := [2]int{x, yA}
						filledCoordMap[candidatePos] = true
					}
				}
			} else {
				// If neither position lies on the same axis, skip
				continue
			}
		}
	}
	return filledCoordMap
}

func _bucketFill(existingPositions map[[2]int]bool, startCoords [2]int, bucketFilledPositions map[[2]int]bool) {
	fmt.Println("Bucket filling", startCoords)
	currentX := startCoords[0]
	currentY := startCoords[1]

	if currentX < 0 || currentY < 0 {
		fmt.Println("Out of bounds.")
		return
	}
	_, exists := existingPositions[startCoords]
	if exists {
		fmt.Println(startCoords, "already in initial positions.")
		return
	}
	_, exists = bucketFilledPositions[startCoords]
	if exists {
		fmt.Println("Previously already bucketfilled.")
		return
	}
	fmt.Println("Seems valid, adding to bucket filled positions.")
	bucketFilledPositions[startCoords] = true
	candidatePositions := [][2]int{
		{currentX, currentY - 1},
		{currentX, currentY + 1},
		{currentX - 1, currentY},
		{currentX + 1, currentY},
	}
	for _, candidatePosition := range candidatePositions {
		_, exists = bucketFilledPositions[candidatePosition]
		if exists {
			fmt.Println(candidatePosition, "already bucketfilled.")
			continue
		}
		_, exists = existingPositions[candidatePosition]
		if exists {
			fmt.Println(candidatePosition, "is already part of the edges/corners.")
			continue
		}
		_bucketFill(existingPositions, candidatePosition, bucketFilledPositions)
	}
}

func bucketFill(existingPositions map[[2]int]bool, startCoords [2]int, bucketFilledPositions map[[2]int]bool, maxX int, maxY int) {
	fmt.Println("Bucket filling", startCoords)
	currentX := startCoords[0]
	currentY := startCoords[1]

	if currentX < 0 || currentY < 0 || currentX > maxX || currentY > maxY {
		fmt.Println("Start coordinates out of bounds.")
		return
	}
	stack := [][2]int{startCoords}
	visited := make(map[[2]int]bool)
	visited[startCoords] = true
	for len(stack) != 0 {
		// Pop value from stack.
		candidatePosition := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// fmt.Println("Popped", candidatePosition, "from stack of size", len(stack))

		candidateX := candidatePosition[0]
		candidateY := candidatePosition[1]

		_, exists := existingPositions[candidatePosition]
		if exists {
			// fmt.Println(candidatePosition, "is already part of the edges/corners.")
			continue
		}
		bucketFilledPositions[candidatePosition] = true

		candidatePositions := [][2]int{
			{candidateX, candidateY - 1},
			{candidateX, candidateY + 1},
			{candidateX - 1, candidateY},
			{candidateX + 1, candidateY},
		}
		for _, newCandidatePos := range candidatePositions {
			_, previouslyChecked := visited[newCandidatePos]
			if previouslyChecked {
				continue
			}
			if newCandidatePos[0] < 0 || newCandidatePos[1] < 0 || newCandidatePos[0] > maxX || newCandidatePos[1] > maxY {
				// fmt.Println("Out of bounds.")
				continue
			}
			_, exists = existingPositions[newCandidatePos]
			if exists {
				// fmt.Println(newCandidatePos, "is already part of the edges/corners.")
				continue
			}
			stack = append(stack, newCandidatePos)
			visited[newCandidatePos] = true
		}
	}
}

func visualize(redCoords map[[2]int]bool, greenCoords map[[2]int]bool, maxX int, maxY int) {
	// Initialize empty viz
	viz := make([][]int, maxY+1)
	for i := range viz {
		vizRow := make([]int, maxX+1)
		viz[i] = vizRow
	}
	for redCoord := range redCoords {
		viz[redCoord[1]][redCoord[0]] = 1
	}
	for greenCoord := range greenCoords {
		viz[greenCoord[1]][greenCoord[0]] = 2
	}
	for _, vizRow := range viz {
		fmt.Println(vizRow)
	}
}

func MainPart2() {
	// positions, err := parseInputFile("test_inputs/p1_example.txt")
	positions, err := parseInputFile("input.txt")
	if err != nil {
		panic(err)
	}

	// simplify positions to start at {0, 0}
	minX := math.MaxInt
	minY := math.MaxInt
	maxX := -1
	maxY := -1
	for _, position := range positions {
		if position[0] < minX {
			minX = position[0]
		}
		if position[1] < minY {
			minY = position[1]
		}
		if position[0] > maxX {
			maxX = position[0]
		}
		if position[1] > maxY {
			maxY = position[1]
		}
	}
	// Normalize
	for posInd, position := range positions {
		positions[posInd] = [2]int{position[0] - minX, position[1] - minY}
	}
	// Correct upper bounds too
	maxX = maxX - minX
	maxY = maxY - minY
	fmt.Println("X:", minX, "to", maxX)
	fmt.Println("Y:", minY, "to", maxY)

	// Create list of "green" positions, i.e. filled between all existing positions.
	// fmt.Println(positions)
	fmt.Println("Got initial red coordinates.")
	redCoordMap := make(map[[2]int]bool)
	for _, position := range positions {
		redCoordMap[position] = true
	}
	filledCoordMap := createFilledPositions(redCoordMap)
	// fmt.Println(filledCoordMap)
	// fmt.Println("Created filled positions.")
	// existingPositions := append(positions, filledPositions...)
	existingPositionsMap := make(map[[2]int]bool)
	for redCoord := range redCoordMap {
		existingPositionsMap[redCoord] = true
	}
	for filledCoord := range filledCoordMap {
		existingPositionsMap[filledCoord] = true
	}
	// fmt.Println(existingPositions)
	// visualize(redCoordMap, make(map[[2]int]bool))
	// visualize(redCoordMap, filledCoordMap, maxX, maxY)
	// panic("Check")
	bucketFilledPositions := make(map[[2]int]bool)
	fmt.Println("Doing bucket filling")
	bucketFill(existingPositionsMap, [2]int{maxX / 2, maxY / 2}, bucketFilledPositions, maxX, maxY)
	// fmt.Println(bucketFilledPositions)
	filledPositions := make(map[[2]int]bool)
	for redCoord := range redCoordMap {
		filledPositions[redCoord] = true
	}
	for filledCoord := range filledCoordMap {
		filledPositions[filledCoord] = true
	}
	for bucketFilledPosition := range bucketFilledPositions {
		filledPositions[bucketFilledPosition] = true
	}
	// visualize(redCoordMap, filledPositions, maxX, maxY)
	areaMat := make([][]int, len(positions))
	for indA, positionA := range positions {
		areaRow := make([]int, len(positions))
		for indB, positionB := range positions {
			areaRow[indB] = 0
			if indA <= indB {
				areaRow[indB] = 0
			} else {
				if !isRectangleComplete(filledPositions, positionA, positionB) {
					areaRow[indB] = -1
				} else {
					areaRow[indB] = computeArea(positionA, positionB)
					fmt.Println("For", positionA, "and", positionB, "area=", areaRow[indB])
				}
			}
		}
		areaMat[indA] = areaRow
	}
	fmt.Println(areaMat)
	maxIndY, maxIndX := findLargestIndex(areaMat)
	largestArea := areaMat[maxIndY][maxIndX]
	fmt.Println("Largest area (P2):", largestArea)
}
