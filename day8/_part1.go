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

func findCircuits(connections [][2]int) (circuits [][]int) {
	circuits = append(circuits, []int{connections[0][0], connections[0][1]})
	for _, connection := range connections[1:] {
		fmt.Println("Checking circuits for", connection, circuits)
		// if i == 11 {
		// 	return circuits
		// }
		maxCircuitLength := 0
		for _, circuit := range circuits {
			maxCircuitLength = max(maxCircuitLength, len(circuit))
		}
		foundMatch := false
		for i := range maxCircuitLength {
			fmt.Println(i, "/", maxCircuitLength)
			if foundMatch {
				break
			}
			for indCircuit, circuit := range circuits {
				if i >= len(circuit) {
					break
				}
				if connection[0] == circuit[i] || connection[1] == circuit[i] {
					fmt.Println("Found match of", connection, "in circuit", circuit)
					foundMatch = true
					if !slices.Contains(circuit, connection[0]) {
						circuit = append(circuit, connection[0])
					}
					if !slices.Contains(circuit, connection[1]) {
						circuit = append(circuit, connection[1])
					}
					circuits[indCircuit] = circuit
					break
				}
			}
		}
		if !foundMatch {
			circuits = append(circuits, []int{connection[0], connection[1]})
		}
	}

	return circuits
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
	connections := make([][2]int, 0)
	for {
		minIndY, minIndX := findMinimumIndex(distanceMat)
		if minIndY < 0 || minIndX < 0 {
			break
		}
		fmt.Println("Found connection between", positions[minIndY], "and", positions[minIndX])
		// for i := range distanceMat[minIndY] {
		// 	distanceMat[minIndY][i] = 0
		// }
		// connections = append(connections, [2]int{minIndY, minIndX})
		// Simplification since order does not matter.
		// We ensure that the connections are lower-values-first
		if minIndY < minIndX {
			connections = append(connections, [2]int{minIndY, minIndX})
		} else {
			connections = append(connections, [2]int{minIndX, minIndY})
		}
		distanceMat[minIndY][minIndX] = 0
	}
	fmt.Println(connections)
	circuits := findCircuits(connections)
	fmt.Println(circuits)

	fmt.Println("Circuit length check")
	totalProduct := 1
	for i, circuit := range circuits {
		if i == 3 {
			break
		}
		totalProduct = totalProduct * len(circuit)
		fmt.Println(len(circuit))
	}

	fmt.Println("Total sum (P1):", totalProduct)
}
