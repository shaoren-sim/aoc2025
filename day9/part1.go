package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func parseInputFile(inputPath string) (positions [][2]int, err error) {
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
		position := [2]int{}
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

func computeArea(a [2]int, b [2]int) int {
	if len(a) != len(b) {
		panic("a and b need to have same length for distance calculation")
	}
	xA := a[0]
	yA := a[1]
	xB := b[0]
	yB := b[1]
	xLen := xA - xB
	if xB > xA {
		xLen = xLen * -1
	}
	xLen = xLen + 1
	yLen := yA - yB
	if yB > yA {
		yLen = yLen * -1
	}
	yLen = yLen + 1
	return xLen * yLen
}

func findLargestIndex(areaMat [][]int) (int, int) {
	currentMaxVal := math.MinInt
	maxIndY := -1
	maxIndX := -1
	for indY, areaVec := range areaMat {
		for indX, valX := range areaVec {
			if indY <= indX {
				continue
			}
			if valX == 0 {
				continue
			}
			if valX > currentMaxVal {
				maxIndY = indY
				maxIndX = indX
				currentMaxVal = valX
			}
		}
	}
	return maxIndY, maxIndX
}

func MainPart1() {
	// positions, err := parseInputFile("test_inputs/p1_example.txt")
	positions, err := parseInputFile("input.txt")
	if err != nil {
		panic(err)
	}

	areaMat := make([][]int, len(positions))
	for indA, positionA := range positions {
		areaRow := make([]int, len(positions))
		for indB, positionB := range positions {
			if indA <= indB {
				areaRow[indB] = 0
			} else {
				areaRow[indB] = computeArea(positionA, positionB)
			}
		}
		areaMat[indA] = areaRow
	}
	fmt.Println(areaMat)
	maxIndY, maxIndX := findLargestIndex(areaMat)
	largestArea := areaMat[maxIndY][maxIndX]
	fmt.Println("Largest area:", largestArea)
}
