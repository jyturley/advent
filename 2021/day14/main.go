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
	fmt.Println(template)
	fmt.Println(rules)
	part1 = Part1(template, rules)
	return part1, part2
}

func Part1(template string, rules map[string]string) int {
	n := 10
	for i := 0; i < n; i++ {
		template = step(template, rules)
	}
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
