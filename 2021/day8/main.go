package main

import (
	"advent/pkg"
	"advent/pkg/execute"
	"fmt"
	"sort"
	"strings"
)

/*
  0:      1:      2:      3:      4:
 aaaa    ....    aaaa    aaaa    ....
b    c  .    c  .    c  .    c  b    c
b    c  .    c  .    c  .    c  b    c
 ....    ....    dddd    dddd    dddd
e    f  .    f  e    .  .    f  .    f
e    f  .    f  e    .  .    f  .    f
 gggg    ....    gggg    gggg    ....

  5:      6:      7:      8:      9:
 aaaa    aaaa    aaaa    aaaa    aaaa
b    .  b    .  .    c  b    c  b    c
b    .  b    .  .    c  b    c  b    c
 dddd    dddd    ....    dddd    dddd
.    f  e    f  .    f  e    f  .    f
.    f  e    f  .    f  e    f  .    f
 gggg    gggg    ....    gggg    gggg
*/

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
		numToPattern := map[int]string{}
		total += Analyze(entry, numToPattern)
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

func Analyze(entry Entry, numToPattern map[int]string) int {
	findOneSevenFourEight(entry, numToPattern)
	findThree(entry, numToPattern)
	findFiveTwo(entry, numToPattern)
	findZeroSixNine(entry, numToPattern)

	return getValue(entry.OutputValues, numToPattern)
}

func getValue(outputvalues [4]string, numToPattern map[int]string) int {
	out := [4]int{}
	for i, ov := range outputvalues {
		out[i] = match(ov, numToPattern)
	}

	ans := 1000*out[0] + 100*out[1] + 10*out[2] + out[3]
	fmt.Println(ans)
	return ans
}

func match(ov string, numToPattern map[int]string) int {
	sov := strings.Split(ov, "")
	sort.Strings(sov)
	sortedOV := strings.Join(sov, "")
	for k, v := range numToPattern {
		sv := strings.Split(v, "")
		sort.Strings(sv)
		sortedV := strings.Join(sv, "")
		if sortedOV == sortedV {
			return k
		}
	}

	return -10
}

func findOneSevenFourEight(entry Entry, numToPattern map[int]string) {
	for _, val := range entry.SignalPatterns {
		switch len(val) {
		case 2:
			numToPattern[1] = val
		case 3:
			numToPattern[7] = val
		case 4:
			numToPattern[4] = val
		case 7:
			numToPattern[8] = val
		}
	}
}

func findZeroSixNine(entry Entry, numToPattern map[int]string) {
	for _, val := range entry.SignalPatterns {
		if len(val) == 6 {
			if contains(val, numToPattern[4]) {
				numToPattern[9] = val
			} else {
				if contains(val, numToPattern[7]) {
					numToPattern[0] = val
				} else {
					numToPattern[6] = val
				}
			}
		}
	}
}

func findFiveTwo(entry Entry, numToPattern map[int]string) {
	for _, val := range entry.SignalPatterns {
		if len(val) == 5 && val != numToPattern[3] {
			d := diff(val, numToPattern[4])
			if len(d) == 2 {
				numToPattern[5] = val
			} else {
				numToPattern[2] = val
			}
		}
	}
}

func findThree(entry Entry, numToPattern map[int]string) {
	seven := numToPattern[7]
	for _, val := range entry.SignalPatterns {
		if len(val) == 5 && contains(val, seven) {
			numToPattern[3] = val
		}
	}
}

func diff(bigStr, smallStr string) string {
	// fmt.Println("diffing ", bigStr, "and", smallStr)
	out := ""
	seen := map[string]bool{}
	for _, sc := range strings.Split(smallStr, "") {
		seen[sc] = true
	}

	for _, c := range strings.Split(bigStr, "") {
		if _, ok := seen[c]; !ok {
			out += c
		}
	}

	return out
}

func contains(s, threeStr string) bool {
	seen := map[string]bool{}
	for _, sc := range strings.Split(s, "") {
		seen[sc] = true
	}

	for _, c := range strings.Split(threeStr, "") {
		if _, ok := seen[c]; !ok {
			return false
		}
	}

	return true
}

func missingLetter(bigger, smaller string) string {
	if len(bigger) <= len(smaller) {
		panic("bad argument")
	}

	if len(bigger) != len(smaller)+1 {
		panic("bad argument len")
	}

	seen := map[string]bool{}
	for _, sc := range strings.Split(smaller, "") {
		seen[sc] = true
	}

	for _, bc := range strings.Split(bigger, "") {
		if _, ok := seen[bc]; !ok {
			return bc
		}
	}

	return "x"
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
