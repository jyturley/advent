package main

import (
	"advent/pkg"
	"advent/pkg/execute"
	"fmt"
	"strings"
)

func run(input string) (part1 interface{}, part2 interface{}) {
	moves, amount := parse(input)
	deltaX, deltaDepth := 0, 0
	for i := 0; i < len(moves); i++ {
		switch moves[i] {
		case "forward":
			deltaX += amount[i]
		case "down":
			deltaDepth += amount[i]
		case "up":
			deltaDepth -= amount[i]
		default:
			fmt.Println("unknown case: ", moves[i])
		}
	}
	part1 = deltaX * deltaDepth

	aim := 0
	deltaX, deltaDepth = 0, 0
	for i := 0; i < len(moves); i++ {
		switch moves[i] {
		case "forward":
			deltaX += amount[i]
			deltaDepth += aim * amount[i]
		case "down":
			aim += amount[i]
		case "up":
			aim -= amount[i]
		default:
			fmt.Println("unknown case: ", moves[i])
		}
	}
	part2 = deltaX * deltaDepth
	return part1, part2
}

func parse(s string) ([]string, []int) {
	lines := strings.Split(s, "\n")
	moves := make([]string, len(lines))
	amount := make([]int, len(lines))
	for i, line := range lines {
		pkg.MustScanf(line, "%s %d", &moves[i], &amount[i])
	}
	return moves, amount
}

func main() {
	year, day := pkg.GetAoVDate()
	fmt.Printf("Advent of Code %s Day %s\n", year, day)
	execute.Run(run, tests, puzzle, true)
}
