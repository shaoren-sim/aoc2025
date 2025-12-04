package main

import "fmt"

func MainPart2() {
	// parts, err := parseInputFile("test_inputs/p1_example.txt")
	parts, err := parseInputFile("input.txt")
	if err != nil {
		panic(err)
	}

	multidim := getMultidimensionalRepresentation(parts)

	validCount := 0
	roundCount := 0
	for {
		roundCount += 1
		coordsMarkedForRemoval := make([][]int, 0)
		for rowInd, row := range multidim {
			for colInd, rowVal := range row {
				if rowVal == "@" {
					adjacentHits := countAdjacentHits(multidim, rowInd, colInd)
					// fmt.Println("Found @ at", rowInd, colInd, "with", adjacentHits, "adjacent @s")
					if adjacentHits < 4 {
						// fmt.Println("Found valid at", rowInd, colInd)
						coordsMarkedForRemoval = append(coordsMarkedForRemoval, []int{rowInd, colInd})
					}
				}
			}
		}
		validCount += len(coordsMarkedForRemoval)
		// fmt.Println("Total removals on round", roundCount, ":", len(coordsMarkedForRemoval))
		// Remove accordingly by marking
		for _, coordSlice := range coordsMarkedForRemoval {
			multidim[coordSlice[0]][coordSlice[1]] = "x"
		}
		// for _, row := range multidim {
		// 	fmt.Println(row)
		// }
		if len(coordsMarkedForRemoval) == 0 {
			break
		}
	}

	// fmt.Println("final state")
	// for _, row := range multidim {
	// 	fmt.Println(row)
	// }
	// fmt.Println("==endviz")
	fmt.Println("Total counts (P2):", validCount)
}
