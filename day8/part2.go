package main

import "fmt"

func MainPart2() {
	// positions, err := parseInputFile("test_inputs/p1_example.txt")
	positions, err := parseInputFile("input.txt")
	if err != nil {
		panic(err)
	}

	distanceMat := make([][]int, len(positions))
	for indA, positionA := range positions {
		distanceVect := make([]int, len(positions))
		for indB, positionB := range positions {
			if indA <= indB {
				distanceVect[indB] = 0
			} else {
				distanceVect[indB] = computeDistance(positionA, positionB)
			}
		}
		distanceMat[indA] = distanceVect
	}
	// fmt.Println(distanceMat)

	// Find minimum distances
	circuits := make([][]int, 0)
	matchedCircuitInd := -1
	loopInCircuit := false
	finalMinIndY := -1
	finalMinIndX := -1
	for {
		if len(circuits) == 1 && len(circuits[0]) == len(positions) {
			break
		}
		minIndY, minIndX := findMinimumIndex(distanceMat)
		if minIndY < 0 || minIndX < 0 {
			break
		}
		finalMinIndY = minIndY
		finalMinIndX = minIndX
		fmt.Println("Last positions:", positions[finalMinIndX], positions[finalMinIndY])
		fmt.Println("Found connection between", minIndY, "and", minIndX)
		// Adding to circuits
		circuits, matchedCircuitInd, loopInCircuit = addToCircuits(circuits, minIndY, minIndX)
		if !loopInCircuit {
			distanceMat[minIndY][minIndX] = 0
		} else {
			fmt.Println("Loop detected in existing circuit, i.e. DO NOTHING")
			for i := range distanceMat[minIndY] {
				distanceMat[minIndY][i] = 0
			}
		}
		// fmt.Println(circuits, matchedCircuitInd)
		if matchedCircuitInd == -1 {
			// fmt.Println("No match found")
			if minIndY < minIndX {
				circuits = append(circuits, []int{minIndY, minIndX})
			} else {
				circuits = append(circuits, []int{minIndX, minIndY})
			}
			continue
		}
		// Compacting circuits
		for {
			prevLen := len(circuits)
			circuits = compactCircuits(circuits)
			if len(circuits) == prevLen {
				break
			}
		}
	}
	fmt.Println("Final Indices", finalMinIndX, finalMinIndY)
	fmt.Println("Final Positions", positions[finalMinIndX], positions[finalMinIndY])
	fmt.Println("Product of X positions (P2):", positions[finalMinIndX][0]*positions[finalMinIndY][0])
}
