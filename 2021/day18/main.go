package main

import (
	"advent/pkg"
	"advent/pkg/execute"
	"fmt"
	// "strings"
	"strconv"
)

/*
Addition:
- [1,2] + [[3,4],5] becomes [[1,2],[[3,4],5]]

Reduce:
- If any pair is nested inside four pairs, the leftmost such pair explodes.
- If any regular number is 10 or greater, the leftmost such regular number splits.

Explode:
- the pair's left value is added to the first regular number to the left of the exploding pair (if any),
- the pair's right value is added to the first regular number to the right of the exploding pair (if any)
- [[[[[9,8],1],2],3],4] becomes
  [[[[0,9],2],3],4]
  (the 9 has no regular number to its left, so it is not added to any regular number)
- [[6,[5,[4,[3,2]]]],1] becomes
  [[6,[5,[7,0]]],3]

*/

type Pair struct {
	left, right *Pair
	value       int
	isLeaf      bool
}

func run(input string) (part1 interface{}, part2 interface{}) {
	// part 1
	// list := parse(input)
	// part1 = Part1(list)
	return part1, part2
}

func Part1() int {
	out := 0
	return out
}

func Part2() int {
	out := 0
	return out
}

func (p Pair) String() string {
	if p.isLeaf {
		return strconv.Itoa(p.value)
	}

	return fmt.Sprintf("[%s,%s]", p.left.String(), p.right.String())
}

func StrToPair(s string) Pair {
	if s[0] != '[' || s[len(s)-1] != ']' {
		panic(s)
	}

	out := Pair{}
	level, middleCommaIdx := 0, -1
	for i, c := range s {
		switch c {
		case '[':
			level++
		case ']':
			level--
		case ',':
			if level == 1 {
				middleCommaIdx = i
			}
		}
	}
	left, right := s[1:middleCommaIdx], s[middleCommaIdx+1:len(s)-1]
	var leftPair, rightPair Pair
	if left[0] == '[' {
		leftPair = StrToPair(left)
	} else {
		v, err := strconv.Atoi(left[1 : len(left)-1])
		if err != nil {
			panic(err)
		}

		leftPair = Pair{value: v, isLeaf: true}
	}

	if right[0] == '[' {
		rightPair = StrToPair(right)
	} else {
		v, err := strconv.Atoi(right[1 : len(right)-1])
		if err != nil {
			panic(err)
		}

		leftPair = Pair{value: v, isLeaf: true}
	}

	// fmt.Printf("{%s,%s}\n", left, right)
	return Pair{left: leftPair, right: rightPair, isLeaf: false}
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
