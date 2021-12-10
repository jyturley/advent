package main

import (
	"advent/pkg"
	"advent/pkg/execute"
	"fmt"
	"sort"
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
	closing = map[string]string{
		"(": ")",
		"[": "]",
		"{": "}",
		"<": ">",
	}

	autoScore = map[rune]int{
		')': 1,
		']': 2,
		'}': 3,
		'>': 4,
	}
)

func run(input string) (part1 interface{}, part2 interface{}) {
	// part 1
	lines := parse(input)
	part1 = Part1(lines)

	// part 2
	part2 = Part2(lines)
	return part1, part2
}

func Part2(lines []string) int {
	scores := make([]int, 0)
	for _, l := range lines {
		if checkLine(l) == 0 {
			scores = append(scores, getScore(repair(l)))
		}
	}
	sort.Ints(scores)
	return scores[len(scores)/2]
}

func Part1(lines []string) int {
	errorScore := 0
	for _, l := range lines {
		errorScore += checkLine(l)
	}
	return errorScore
}

func repair(l string) string {
	out := ""
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
				panic("calling repair on bad string")
			}
			pop := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if opening[c] != pop {
				panic("calling repair on bad string")
			}
		}
	}

	for len(stack) > 0 {
		pop := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		out += closing[pop]
	}

	return out
}

func getScore(l string) int {
	total := 0
	for _, c := range l {
		// fmt.Printf("%d x 5 + %d = ", total, autoScore[c])
		total = total*5 + autoScore[c]
		// fmt.Println(total)
	}
	return total
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
