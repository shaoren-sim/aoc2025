package main

import "fmt"

func MainPart2() {
	positions, err := parseInputFile("test_inputs/p1_example.txt")
	// verticalLines, err := parseInputFileP2("input.txt")
	if err != nil {
		panic(err)
	}
	for _, position := range positions {
		fmt.Println(position)
	}
}
