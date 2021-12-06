package main

import (
	"advent/pkg"
	"advent/pkg/execute"
	"fmt"
	"strings"
)

func run(input string) (part1 interface{}, part2 interface{}) {
	// part 1
	//list := parse(input)
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
	execute.Run(run, tests, puzzle, true)
}
