package main

import (
	"advent/pkg"
	"advent/pkg/execute"
	"fmt"
	"strings"
)

func run(input string) (part1 interface{}, part2 interface{}) {
	// part 1
	template, rules := parse(input)
	part1 = Part1(template, rules)
	part2 = Part2(template, rules)
	return part1, part2
}

func Part1(template string, rules map[string]string) int {
	template = stepN(template, 10, rules)
	counts := map[string]int{}
	for _, c := range template {
		counts[string(c)] += 1
	}

	mostCommon := string(template[0])
	leastCommon := string(template[0])
	for k, v := range counts {
		if counts[mostCommon] < v {
			mostCommon = k
		}

		if counts[leastCommon] > v {
			leastCommon = k
		}
	}
	return counts[mostCommon] - counts[leastCommon]
}

func Part2(template string, rules map[string]string) int {
	return Part2N(template, rules, 40)
}

func Part2N(template string, rules map[string]string, n int) int {
	pairCounts := getPairCountDict(template)

	for i := 0; i < n; i++ {
		pairCounts = stepCounts(pairCounts, rules)
	}

	charCounts := map[string]int{}
	for pair, count := range pairCounts {
		char2 := string(pair[1])
		charCounts[char2] += count
	}

	charCounts[string(template[0])] += 1

	mostCommon := string(template[0])
	leastCommon := string(template[0])
	for k, v := range charCounts {
		if charCounts[mostCommon] < v {
			mostCommon = k
		}

		if charCounts[leastCommon] > v {
			leastCommon = k
		}
	}
	return charCounts[mostCommon] - charCounts[leastCommon]
}

func getPairCountDict(template string) map[string]int {
	pairCounts := map[string]int{}
	for i := 2; i <= len(template); i++ {
		pair := template[i-2 : i]
		pairCounts[pair] += 1
	}
	return pairCounts
}

func getCountDict(template string) map[string]int {
	counts := map[string]int{}
	for _, c := range template {
		counts[string(c)] += 1
	}
	return counts
}

func stepCounts(pairCounts map[string]int, rules map[string]string) (pairCount map[string]int) {
	newPairCounts := map[string]int{}
	for pair, count := range pairCounts {
		newChar, ok := rules[pair]
		if !ok {
			panic(pair)
		}
		s1 := string(pair[0]) + newChar
		s2 := newChar + string(pair[1])
		newPairCounts[s1] += count
		newPairCounts[s2] += count
	}

	return newPairCounts
}

func stepN(template string, n int, rules map[string]string) string {
	for i := 0; i < n; i++ {
		template = step(template, rules)
	}
	return template
}

func step(template string, rules map[string]string) string {
	out := ""
	out += string(template[0:2][0])
	for i := 2; i <= len(template); i++ {
		pair := template[i-2 : i]
		c := rules[pair]
		out += c + string(pair[1])
	}
	return out
}

func parse(s string) (string, map[string]string) {
	var template string
	lines := strings.Split(s, "\n")
	rules := map[string]string{}
	for i, line := range lines {
		if i == 0 {
			template = line
		} else if strings.Contains(line, "->") {
			var from, to string
			pkg.MustScanf(line, "%s -> %s", &from, &to)
			rules[from] = to
		}
	}
	return template, rules
}

func main() {
	year, day := pkg.GetAoVDate()
	fmt.Printf("Advent of Code %s Day %s\n", year, day)
	execute.Run(run, tests, puzzle, true)
}
