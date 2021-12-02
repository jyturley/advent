package main

import (
	"advent/pkg"
	"advent/pkg/execute"
	"fmt"
)

func run(input string) (interface{}, interface{}) {
	depths := pkg.ParseIntList(input, "\n")

	// part 1
	increaseCount := 0
	for i := 1; i < len(depths); i++ {
		if depths[i-1] < depths[i] {
			increaseCount++
		}
	}

	// part 2
	incSlidingWinCount := 0
	slidingWinValues := make([]int, len(depths))
	for i := 2; i < len(depths); i++ {
		slidingWinSum := depths[i-2] + depths[i-1] + depths[i]
		slidingWinValues[i] = slidingWinSum
	}

	for i := 3; i < len(slidingWinValues); i++ {
		if slidingWinValues[i-1] < slidingWinValues[i] {
			incSlidingWinCount++
		}
	}

	return increaseCount, incSlidingWinCount
}

func main() {
	year, day := pkg.GetAoVDate()
	fmt.Printf("Advent of Code %s Day %s\n", year, day)
	execute.Run(run, tests, puzzle, true)
}
