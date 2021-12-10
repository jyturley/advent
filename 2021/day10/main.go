package main

import (
	"advent/pkg"
	"advent/pkg/execute"
	"fmt"
	"strings"
)

var (
	scores = map[string]int{
		")": 3,
		"]": 57,
		"}": 1197,
		">": 25137,
	}

	opening = map[rune]string{
		')': "(",
		']': "[",
		'}': "{",
		'>': "<",
	}
)

func run(input string) (part1 interface{}, part2 interface{}) {
	// part 1
	lines := parse(input)
	part1 = Part1(lines)
	return part1, part2
}

func Part1(lines []string) int {
	errorScore := 0
	for _, l := range lines {
		errorScore += checkLine(l)
	}
	return errorScore
}

func checkLine(l string) int {
	stack := make([]string, 0)
	for _, c := range l {
		switch c {
		case '[':
			fallthrough
		case '(':
			fallthrough
		case '{':
			fallthrough
		case '<':
			stack = append(stack, string(c))

		default:
			if len(stack) < 1 {
				return scores[string(c)]
			}
			pop := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if opening[c] != pop {
				return scores[string(c)]
			}
		}
	}
	return 0
}

func parse(s string) []string {
	return strings.Split(s, "\n")
}

func main() {
	year, day := pkg.GetAoVDate()
	fmt.Printf("Advent of Code %s Day %s\n", year, day)
	execute.Run(run, tests, puzzle, true)
}
