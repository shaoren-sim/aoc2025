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

func addToCircuits(circuits [][]int, indFrom int, indTo int) ([][]int, int, bool) {
	for indCircuit, circuit := range circuits {
		fromIn := slices.Contains(circuit, indFrom)
		toIn := slices.Contains(circuit, indTo)
		if fromIn && toIn {
			return circuits, indCircuit, true
		}
		if fromIn || toIn {
			if !fromIn {
				circuit = append(circuit, indFrom)
			}
			if !toIn {
				circuit = append(circuit, indTo)
			}
			circuits[indCircuit] = circuit
			return circuits, indCircuit, false
		}
	}
	return circuits, -1, false
}
func compactCircuits(circuits [][]int) (newCircuits [][]int) {
	usedIndices := make(map[int]int)
	for _, circuit := range circuits {
		newCircuit := make([]int, 0)
		previousFound := false
		for _, val := range circuit {
			_, exists := usedIndices[val]
			if !exists {
				usedIndices[val] = len(newCircuits)
				newCircuit = append(newCircuit, val)
			} else {
				destinationIndex := usedIndices[val]
				for _, valToAdd := range circuit {
					if !slices.Contains(newCircuits[destinationIndex], valToAdd) {
						newCircuits[destinationIndex] = append(newCircuits[destinationIndex], valToAdd)
					}
				}
				for _, node := range circuit {
					usedIndices[node] = destinationIndex
				}
				previousFound = true
				break
			}
		}
		if !previousFound {
			newCircuits = append(newCircuits, newCircuit)
		}
	}

	return newCircuits
}

func MainPart1() {
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
	ind := 0
	for {
		fmt.Println(ind)
		// if ind == 10 {
		// 	break
		// }
		if ind == 1000 {
			break
		}
		ind += 1
		minIndY, minIndX := findMinimumIndex(distanceMat)
		if minIndY < 0 || minIndX < 0 {
			break
		}
		fmt.Println("Found connection between", minIndY, "and", minIndX)
		// Adding to circuits.
		circuits, matchedCircuitInd, _ = addToCircuits(circuits, minIndY, minIndX)
		distanceMat[minIndY][minIndX] = 0
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
		// for _, val := range circuits[matchedCircuitInd] {
		// 	// If a value already exists in a circuit, disable tracking of all linked values.
		// 	// i.e. the "nothing happens" case
		// 	distanceMat[minIndY][val] = 0
		// }
	}
	fmt.Println(circuits)
	// Compacting circuits
	for {
		prevLen := len(circuits)
		circuits = compactCircuits(circuits)
		if len(circuits) == prevLen {
			break
		}
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
