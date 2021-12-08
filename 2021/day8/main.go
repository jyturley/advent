package main

import (
	"advent/pkg"
	"advent/pkg/execute"
	"fmt"
	"strings"
)

const (
	unknown = -8
)

type Entry struct {
	SignalPatterns [10]string
	OutputValues   [4]string
}

func run(input string) (part1 interface{}, part2 interface{}) {
	// part 1
	entries := parse(input)
	part1 = Part1(entries)

	// part 2
	part2 = Part2(entries)
	return part1, part2
}

func Part2(entries []Entry) int {
	total := 0
	for _, entry := range entries {
		total += Analyze(entry)
	}
	return total
}

func Part1(entries []Entry) int {
	one, four, seven, eight := 0, 0, 0, 0
	for _, entry := range entries {
		for _, val := range entry.OutputValues {
			switch len(val) {
			case 2:
				one++
			case 3:
				seven++
			case 4:
				four++
			case 7:
				eight++
			}
		}
	}
	return one + four + seven + eight
}

func Analyze(entry Entry) int {

}

func parse(s string) []Entry {
	entries := make([]Entry, 0)
	lines := strings.Split(s, "\n")
	for _, line := range lines {
		s := strings.Split(line, " ")
		signals := [10]string{}
		output := [4]string{}
		copy(signals[:], s[:10])
		copy(output[:], s[11:])
		entries = append(entries, Entry{signals, output})
	}
	return entries
}

func main() {
	year, day := pkg.GetAoVDate()
	fmt.Printf("Advent of Code %s Day %s\n", year, day)
	execute.Run(run, tests, puzzle, true)
}
