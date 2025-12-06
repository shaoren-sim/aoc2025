package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func parseInputFileP2(inputPath string) ([][]string, error) {
	verticalLines := initializeVerticalLines(inputPath)
	file, err := os.Open(inputPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	lineInd := 0
	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}
		// If all values in this line are "", mark it, it's a break
		for i, char := range scanner.Text() {
			charStr := string(char)
			verticalLines[i][lineInd] = charStr
		}
		lineInd = lineInd + 1
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return verticalLines, nil
}

func initializeVerticalLines(inputPath string) [][]string {
	file, err := os.Open(inputPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	numLines := 0
	lineLength := 0
	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}
		if len(scanner.Text()) > lineLength {
			lineLength = len(scanner.Text())
		}
		numLines = numLines + 1
	}

	verticalLines := make([][]string, 0)
	for i := 0; i < lineLength; i++ {
		verticalLines = append(verticalLines, make([]string, numLines))
	}
	return verticalLines
}

func allBlanks(a []string) bool {
	for i := 1; i < len(a); i++ {
		if a[i] != " " {
			return false
		}
		if a[i] != a[0] {
			return false
		}
	}
	return true
}

func doWeirdSum(stringsToOpOn []string, op string) int {
	initialValue, err := strconv.Atoi(stringsToOpOn[0])
	if err != nil {
		panic(err)
	}
	currentSum := initialValue
	for _, valStr := range stringsToOpOn[1:] {
		val, err := strconv.Atoi(valStr)
		if err != nil {
			panic(err)
		}
		switch op {
		case "*":
			currentSum = currentSum * val
		case "+":
			currentSum = currentSum + val
		}
	}
	return currentSum
}

func doWeirdMath(verticalLines [][]string) int {
	totalSum := 0
	valStrCache := make([]string, 0)
	op := ""
	for _, verticalLine := range verticalLines {
		if allBlanks(verticalLine) {
			// Break line reached, compute values from cache
			currentSum := doWeirdSum(valStrCache, op)
			// fmt.Println("Current Sum:", currentSum)
			totalSum = totalSum + currentSum
			// Reset value cache
			valStrCache = make([]string, 0)
			op = ""
			continue
		}
		if verticalLine[len(verticalLine)-1] != " " {
			op = verticalLine[len(verticalLine)-1]
		}
		valStrCacheForLine := ""
		for _, charStr := range verticalLine[:len(verticalLine)-1] {
			if charStr != " " {
				valStrCacheForLine = valStrCacheForLine + charStr
			}
		}
		valStrCache = append(valStrCache, valStrCacheForLine)
	}

	// Last line needs to be summed too.
	currentSum := doWeirdSum(valStrCache, op)
	// fmt.Println("Current Sum:", currentSum)
	totalSum = totalSum + currentSum
	return totalSum
}

func MainPart2() {
	// verticalLines, err := parseInputFileP2("test_inputs/p1_example.txt")
	verticalLines, err := parseInputFileP2("input.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("Total sum (P2):", doWeirdMath(verticalLines))
}
