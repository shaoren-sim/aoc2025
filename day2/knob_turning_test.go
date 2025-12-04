package main

import (
	"fmt"
	"reflect"
	"testing"
)

type ResultItem struct {
	Val    int
	Clicks int
	Total  int
}

func getTargets() []ResultItem {
	return []ResultItem{
		{Val: 82, Clicks: 1, Total: 1},
		{Val: 52, Clicks: 0, Total: 1},
		{Val: 0, Clicks: 1, Total: 2},
		{Val: 95, Clicks: 0, Total: 2},
		{Val: 55, Clicks: 1, Total: 3},
		{Val: 0, Clicks: 1, Total: 4},
		{Val: 99, Clicks: 0, Total: 4},
		{Val: 0, Clicks: 1, Total: 5},
		{Val: 14, Clicks: 0, Total: 5},
		{Val: 32, Clicks: 1, Total: 6},
	}
}

func TestTurningExpectations(t *testing.T) {
	knobVal := 50
	totalClicks := 0

	lines, err := parseInputFile("test_inputs/p1_example.txt")
	if err != nil {
		panic(err)
	}

	targets := getTargets()

	for i, line := range lines {
		fmt.Println("----------")
		fmt.Println("Iteration", i, line)
		target := targets[i]

		direction, steps := getRotateDetails(line)
		// zeroHits, newKnobVal := turnKnobP2(knobVal, direction, steps)
		// knobVal = newKnobVal
		// totalClicks = totalClicks + zeroHits

		zeroHits, _ := turnKnob(knobVal, direction, steps)
		clicksFromOvershoots, newKnobVal := turnKnobP2(knobVal, direction, steps)

		knobVal = newKnobVal
		totalClicks = totalClicks + zeroHits + clicksFromOvershoots

		actual := ResultItem{Val: knobVal, Clicks: zeroHits + clicksFromOvershoots, Total: totalClicks}
		if !reflect.DeepEqual(target, actual) {
			t.Errorf("Iteration %d failed: \nExpected: %+v \nActual:   %+v", i, target, actual)
		} else {
			fmt.Printf("Iteration %d passed. Result: %+v\n", i, actual)
		}
	}
	if totalClicks != 6 {
		t.Errorf("Final count failed: \nExpected: 6 \nActual: %+v", totalClicks)
	}
}
