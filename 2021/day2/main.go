package main

import (
	"advent/pkg"
	"advent/pkg/execute"
	"fmt"
	"strings"
)

// returns part1 and part2
func run(input string) (interface{}, interface{}) {
	moves, amount := parse(input)
	fmt.Println(moves)
	fmt.Println(amount)
	part1, part2 := 0, 0

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
	execute.Run(run, tests, puzzle, true)
}
