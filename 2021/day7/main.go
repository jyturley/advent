package main

import (
	"advent/pkg"
	"advent/pkg/execute"
	"fmt"
	"strconv"
	"strings"
)

func run(input string) (part1 interface{}, part2 interface{}) {
	// part 1
	positions := parse(input)
	part1 = FindAlignment(positions)

	// part 1
	positions = parse(input)
	part2 = FindAlignmentP2(positions)

	return part1, part2
}

func FindAlignmentP2(positions []int) int {
	minFuel := 9990000000
	min, max := pkg.Min(positions...), pkg.Max(positions...)
	size := max - min + 1
	vals := make([]int, size)
	for _, p := range positions {
		vals[p-min] += 1
	}

	for pos := min; pos < max; pos++ {
		cost := 0
		for i := 0; i < size; i++ {
			cost += vals[i] * crabWalkCost(pos, i)
		}
		minFuel = pkg.Min(minFuel, cost)
	}

	return minFuel
}

func crabWalkCost(from, to int) int {
	return summation(pkg.Abs((from - to)))
}

func summation(n int) int {
	out := 0
	for i := 1; i <= n; i++ {
		out += i
	}
	return out
}

func FindAlignment(positions []int) int {
	minFuel := 10000000
	min, max := pkg.Min(positions...), pkg.Max(positions...)
	size := max - min + 1
	vals := make([]int, size)
	for _, p := range positions {
		vals[p-min] += 1
	}

	for pos := min; pos < max; pos++ {
		cost := 0
		for i := 0; i < size; i++ {
			cost += vals[i] * pkg.Abs((pos - i))
		}
		minFuel = pkg.Min(minFuel, cost)
	}

	return minFuel
}

func parse(s string) []int {
	out := make([]int, 0)
	split := strings.Split(s, ",")
	for _, s := range split {
		i, _ := strconv.Atoi(s)
		out = append(out, i)
	}
	return out
}

func main() {
	year, day := pkg.GetAoVDate()
	fmt.Printf("Advent of Code %s Day %s\n", year, day)
	execute.Run(run, tests, puzzle, true)
}
