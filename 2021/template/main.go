package main

import (
	"advent/pkg"
	"advent/pkg/execute"
	"fmt"
	"strings"
)

func run(input string) (interface{}, interface{}) {
	//list := parse(input)

	part1, part2 := 0, 0

	return part1, part2
}

// func parse(s string) {
// 	lines := strings.Split(s, "\n")
// 	for _, line := range lines {
// 		pkg.MustScanf(line, "")
// 	}
// 	return
// }

func main() {
	year, day := pkg.GetAoVDate()
	fmt.Printf("Advent of Code %s Day %s\n", year, day)
	execute.Run(run, nil, puzzle, true)
}
