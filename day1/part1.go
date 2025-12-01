package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func parseInputFile(inputPath string) (lines []string, err error) {
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
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func getRotateDetails(input string) (direction string, steps int) {
	direction = input[:1]
	steps, err := strconv.Atoi(input[1:])
	if err != nil {
		panic(err)
	}

	return direction, steps
}

func turnKnob(currentVal int, direction string, steps int) (zeroHits int, finalVal int) {
	switch direction {
	case "L":
		finalVal = currentVal - steps
		for finalVal < 0 {
			finalVal = finalVal + 100
		}
	case "R":
		finalVal = currentVal + steps
		for finalVal > 99 {
			finalVal = finalVal - 100
		}
	}
	if finalVal == 0 {
		return 1, finalVal
	} else {
		return 0, finalVal
	}
}

func MainPart1() {
	knobVal := 50
	totalZeroHits := 0

	// lines, err := parseInputFile("test_inputs/p1_example.txt")
	lines, err := parseInputFile("input.txt")
	if err != nil {
		panic(err)
	}

	// Loop through lines and turn the knob accordingly.
	for _, line := range lines {
		direction, steps := getRotateDetails(line)
		// fmt.Println(direction, steps)
		// fmt.Println("Current knob val", knobVal)
		zeroHits, newKnobVal := turnKnob(knobVal, direction, steps)
		knobVal = newKnobVal
		totalZeroHits = totalZeroHits + zeroHits
		fmt.Println(totalZeroHits, knobVal)
		fmt.Println("----------")
	}

	fmt.Println("Final zero count:", totalZeroHits)

}
