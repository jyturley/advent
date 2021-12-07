package main

import (
	"advent/pkg"
	"advent/pkg/execute"
	"fmt"
	"strconv"
	"strings"
)

type Fish int

const (
	newFish       = Fish(8)
	regenatedFish = Fish(6)
)

func run(input string) (part1 interface{}, part2 interface{}) {
	// part 1
	fish := parse(input)
	part1 = FishAfterDays(80, fish)

	// part 2
	fish = parse(input)
	part2 = FishAfterManyDays(256, fish)
	return part1, part2
}

func FishAfterManyDays(numDays int, fish []Fish) int {
	count := [9]int{}
	for _, f := range fish {
		count[int(f)] += 1
	}

	days := 0
	for days < numDays {
		numNewFish := count[0]
		numRegeneratedFish := count[0]
		count[0], count[1], count[2], count[3], count[4], count[5], count[6], count[7], count[8] =
			count[1], count[2], count[3], count[4], count[5], count[6], count[7]+numRegeneratedFish, count[8], numNewFish
		days++
	}
	sum := count[0] + count[1] + count[2] + count[3] + count[4] + count[5] + count[6] + count[7] + count[8]
	fmt.Printf("After %d days, num fish: %d \n", days+1, sum)
	fmt.Println(count)
	return sum
}

func FishAfterDays(numDays int, fish []Fish) int {
	days := 0
	for days < numDays {
		n := len(fish)
		for i := 0; i < n; i++ {
			f := fish[i]
			if f == 0 {
				fish = append(fish, newFish)
				fish[i] = regenatedFish
				continue
			}
			fish[i] -= 1
		}
		// fmt.Printf("After %d days, num fish: %d  %s\n", days+1, len(fish), time.Now().Sub(loopStart))
		// fmt.Printf("After %d days, num fish: %d, %v\n", days+1, len(fish), fish)
		days++
	}
	return len(fish)
}

func parse(s string) []Fish {
	lines := strings.Split(s, "\n")
	out := make([]Fish, 0, 9903182618)
	for _, line := range lines {
		fish := strings.Split(line, ",")
		for _, f := range fish {
			i, _ := strconv.Atoi(f)
			out = append(out, Fish(i))
		}
	}
	return out
}

func main() {
	year, day := pkg.GetAoVDate()
	fmt.Printf("Advent of Code %s Day %s\n", year, day)
	execute.Run(run, tests, puzzle, true)
}
