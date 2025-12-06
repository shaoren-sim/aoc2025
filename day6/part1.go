package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func parseInputFile(inputPath string) (questionLines []string, operations string, err error) {
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
			questionLines = append(questionLines, part)
		}
	}

	// Last line is the operations

	if err := scanner.Err(); err != nil {
		return nil, "", err
	}

	return questionLines[:len(questionLines)-1], questionLines[len(questionLines)-1], nil
}

func splitParts(line string) (parts []string) {
	return regexp.MustCompile(`[^\s]+`).FindAllString(line, -1)
}

func MainPart1() {
	// questionLines, operationLine, err := parseInputFile("test_inputs/p1_example.txt")
	questionLines, operationLine, err := parseInputFile("input.txt")
	if err != nil {
		panic(err)
	}

	questions := make([][]int, 0)

	for _, questionLine := range questionLines {
		line := splitParts(questionLine)
		lineInts := make([]int, 0)
		for _, valStr := range line {
			val, err := strconv.Atoi(valStr)
			if err != nil {
				panic(err)
			}
			lineInts = append(lineInts, val)
		}
		questions = append(questions, lineInts)
	}
	operations := splitParts(operationLine)

	totalSum := 0
	for x, opStr := range operations {
		answer := questions[0][x]
		usedInOps := make([]int, 0)
		for _, questionLine := range questions[1:] {
			switch opStr {
			case "*":
				answer = answer * questionLine[x]
			case "+":
				answer = answer + questionLine[x]
			}
			usedInOps = append(usedInOps, questionLine[x])
		}
		// fmt.Println(answer, "for", opStr, usedInOps)
		totalSum = totalSum + answer
	}

	fmt.Println("Total sum (P1):", totalSum)
}
