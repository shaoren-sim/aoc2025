package main

import "fmt"

func turnKnobP2(currentVal int, direction string, steps int) (clicks int, finalVal int) {
	switch direction {
	case "L":
		if currentVal == 0 {
			clicks = clicks - 1
		}
		finalVal = currentVal - steps
		for finalVal < 0 {
			finalVal = finalVal + 100
			clicks = clicks + 1
		}
	case "R":
		finalVal = currentVal + steps
		for finalVal > 99 {
			finalVal = finalVal - 100
			clicks = clicks + 1
		}
		if finalVal == 0 {
			clicks = clicks - 1
		}
	}
	return clicks, finalVal
}

func MainPart2() {
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
		zeroHits, _ := turnKnob(knobVal, direction, steps)
		clicksFromOvershoots, newKnobVal := turnKnobP2(knobVal, direction, steps)

		knobVal = newKnobVal
		totalZeroHits = totalZeroHits + zeroHits + clicksFromOvershoots
	}

	fmt.Println("Final click count:", totalZeroHits)
}
