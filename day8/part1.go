package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func parseInputFile(inputPath string) (positions [][3]int, err error) {
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
		position := [3]int{}
		index := 0
		for part := range parts {
			partVal, err := strconv.Atoi(part)
			if err != nil {
				panic(err)
			}
			position[index] = partVal
			index += 1
		}
		positions = append(positions, position)
	}

	// Last line is the operations
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return positions, nil
}

func computeDistance(a [3]int, b [3]int) int {
	// faux-euclidean distance calculation, no sqrt to keep in int.
	if len(a) != len(b) {
		panic("a and b need to have same length for distance calculation")
	}
	distance := 0
	for i := range a {
		distance += (a[i] - b[i]) * (a[i] - b[i])
	}
	return distance
}

func findMinimumIndex(distanceMat [][]int) (int, int) {
	currentMinVal := math.MaxInt
	minIndY := -1
	minIndX := -1
	for indY, distanceVec := range distanceMat {
		for indX, valX := range distanceVec {
			if indY <= indX {
				continue
			}
			if valX == 0 {
				continue
			}
			if valX < currentMinVal {
				minIndY = indY
				minIndX = indX
				currentMinVal = valX
			}
		}
	}
	return minIndY, minIndX
}

func addToCircuits(circuits [][]int, indFrom int, indTo int) ([][]int, int) {
	// fmt.Println("Checking circuits for", indFrom, "to", indTo, circuits)
	maxCircuitLength := 0
	for _, circuit := range circuits {
		maxCircuitLength = max(maxCircuitLength, len(circuit))
	}
	for i := range maxCircuitLength {
		// fmt.Println(i, "/", maxCircuitLength)
		for indCircuit, circuit := range circuits {
			if i >= len(circuit) {
				break
			}
			if indFrom == circuit[i] || indTo == circuit[i] {
				// fmt.Println("Found match of", indFrom, indTo, "in circuit", circuit)
				if !slices.Contains(circuit, indFrom) {
					circuit = append(circuit, indFrom)
				}
				if !slices.Contains(circuit, indTo) {
					circuit = append(circuit, indTo)
				}
				circuits[indCircuit] = circuit
				return circuits, indCircuit
			}
		}
	}
	// Return -1 for no match found.
	return circuits, -1
}

func MainPart1() {
	positions, err := parseInputFile("test_inputs/p1_example.txt")
	// positions, err := parseInputFile("input.txt")
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
	ind := 0
	for {
		fmt.Println(ind, "/", 499000)
		if ind == 10 {
			break
		}
		if ind == 1001 {
			break
		}
		ind += 1
		minIndY, minIndX := findMinimumIndex(distanceMat)
		if minIndY < 0 || minIndX < 0 {
			break
		}
		fmt.Println("Found connection between", minIndY, "and", minIndX)
		// Adding to circuits.
		circuits, matchedCircuitInd = addToCircuits(circuits, minIndY, minIndX)
		distanceMat[minIndY][minIndX] = 0
		fmt.Println(circuits, matchedCircuitInd)
		if matchedCircuitInd == -1 {
			// fmt.Println("No match found")
			if minIndY < minIndX {
				circuits = append(circuits, []int{minIndY, minIndX})
			} else {
				circuits = append(circuits, []int{minIndX, minIndY})
			}
			continue
		}
		// for _, val := range circuits[matchedCircuitInd] {
		// 	// If a value already exists in a circuit, disable tracking of all linked values.
		// 	// i.e. the "nothing happens" case
		// 	distanceMat[minIndY][val] = 0
		// }
	}
	fmt.Println(circuits)
	fmt.Println("Circuit length check")
	circuitLengths := make([]int, 0)
	totalProduct := 1
	allInds := make([]int, len(positions))
	for i := range positions {
		allInds[i] = i
	}
	unusedInds := make([]int, 0)
	usedInds := make([]int, 0)
	for _, circuit := range circuits {
		usedInds = append(usedInds, circuit...)
		circuitLengths = append(circuitLengths, len(circuit))
	}
	for _, ind := range allInds {
		if !slices.Contains(usedInds, ind) {
			unusedInds = append(unusedInds, ind)
		}
	}
	// fmt.Println(circuitLengths)
	slices.Sort(circuitLengths)
	fmt.Println(circuitLengths)
	fmt.Println(unusedInds)
	for i := len(circuitLengths) - 1; i >= 0; i-- {
		circuitLength := circuitLengths[i]
		if i < len(circuitLengths)-3 {
			break
		}
		totalProduct *= circuitLength
	}

	fmt.Println("Total sum (P1):", totalProduct)
}
