package main

import (
	"advent/pkg"
	"advent/pkg/execute"
	"fmt"
	"strings"
	"unicode"
)

// type cave struct {
// 	name        string
// 	connections []*cave
// }

func run(input string) (part1 interface{}, part2 interface{}) {
	// part 1
	cave := parse(input)
	// fmt.Println(cave)
	part1 = Part1(cave)
	part2 = Part2(cave)
	return part1, part2
}

func Part1(caves map[string][]string) int {
	ans := make([]string, 0)
	dfs([]string{"start"}, &ans, caves)
	return len(ans)
}

func Part2(caves map[string][]string) int {
	ans := make([]string, 0)
	dfs2([]string{"start"}, &ans, "", caves)
	return len(ans)
}

func dfs(start []string, ans *[]string, caves map[string][]string) {
	if len(start) < 1 {
		s := fmt.Sprintf("bad args - start:%v ans:%v\n", start, ans)
		panic(s)
	}

	startCave := start[len(start)-1]
	if startCave == "end" {
		s := stringify(start)
		*ans = append(*ans, s)
		return
	}

	seen := false
	// if startCave is lowercase, ensure it hasn't been seen before
	if !unicode.IsUpper(rune(startCave[0])) {
		for i := 0; i < len(start)-1; i++ {
			if startCave == start[i] {
				seen = true
				break
			}
		}
	}

	if seen {
		return
	}

	for _, connectionCave := range caves[startCave] {
		start = append(start, connectionCave)
		dfs(start, ans, caves)
		start = start[:len(start)-1]
	}

	return
}

func dfs2(start []string, ans *[]string, doubleSmallLetter string, caves map[string][]string) {
	if len(start) < 1 {
		s := fmt.Sprintf("bad args - start:%v ans:%v\n", start, ans)
		panic(s)
	}

	startCave := start[len(start)-1]
	if startCave == "end" {
		s := stringify(start)
		// fmt.Println(s)
		*ans = append(*ans, s)
		return
	}

	if startCave == "start" && len(start) > 2 {
		return
	}

	seenCount := 0
	lowercase := !unicode.IsUpper(rune(startCave[0]))

	// if startCave is lowercase, ensure it hasn't been seen before
	if lowercase {
		for i := 0; i < len(start)-1; i++ {
			if startCave == start[i] {
				seenCount++
				break
			}
		}
	}

	if doubleSmallLetter == "" && seenCount == 1 {
		doubleSmallLetter = startCave
	} else if seenCount > 0 {
		return
	}

	for _, connectionCave := range caves[startCave] {
		start = append(start, connectionCave)
		dfs2(start, ans, doubleSmallLetter, caves)
		start = start[:len(start)-1]
	}

	return
}

func stringify(caves []string) string {
	out := ""
	delim := ","
	for _, c := range caves {
		out += c + delim
	}
	return out[:len(out)-1]
}

func parse(s string) map[string][]string {
	caves := map[string][]string{}
	lines := strings.Split(s, "\n")
	for _, line := range lines {
		var from, to string
		split := strings.Split(line, "-")
		from, to = split[0], split[1]
		if _, ok := caves[from]; !ok {
			caves[from] = make([]string, 0)
		}
		caves[from] = append(caves[from], to)

		if _, ok := caves[to]; !ok {
			caves[to] = make([]string, 0)
		}
		caves[to] = append(caves[to], from)
	}
	return caves
}

func main() {
	year, day := pkg.GetAoVDate()
	fmt.Printf("Advent of Code %s Day %s\n", year, day)
	execute.Run(run, tests, puzzle, true)
}
