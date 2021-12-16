package main

import (
	"advent/pkg"
	"advent/pkg/execute"
	"encoding/hex"
	"fmt"
)

/*
0 = 0000
1 = 0001
2 = 0010
3 = 0011
4 = 0100
5 = 0101
6 = 0110
7 = 0111
8 = 1000
9 = 1001
A = 1010
B = 1011
C = 1100
D = 1101
E = 1110
F = 1111
*/

func run(input string) (part1 interface{}, part2 interface{}) {
	// part 1
	list := parse(input)
	part1 = Part1(list)
	return part1, part2
}

func Part1(binary string) int {
	out := 0
	return out
}

func Part2() int {
	out := 0
	return out
}

func parse(s string) string {
	out := ""
	decoded, err := hex.DecodeString(s)
	if err != nil {
		panic(err)
	}
	out = convertToBinaryStr(decoded)
	return out
}

func convertToBinaryStr(decoded []byte) string {
	out := ""
	for _, b := range decoded {
		out = fmt.Sprintf("%s%b", out, b)
	}
	return out
}

func main() {
	year, day := pkg.GetAoVDate()
	fmt.Printf("Advent of Code %s Day %s\n", year, day)
	execute.Run(run, tests, puzzle, true)
}
