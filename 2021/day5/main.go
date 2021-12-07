package main

import (
	"advent/pkg"
	"advent/pkg/execute"
	"fmt"
	"strconv"
	"strings"
)

var (
	dot  = 0
	line = 1
)

type Coord struct {
	x, y int
}

type Segment struct {
	E1, E2 Coord
}

func run(input string) (part1 interface{}, part2 interface{}) {
	// init board
	smallN, bigN := 10, 1000
	segments := parse(input)
	n := bigN
	if len(segments) < 15 {
		n = smallN
	}

	// part 1
	board := initBoard(n)
	AddDots(segments, board)
	print(board)
	part1 = DangerAreas(board)

	// part 2
	board = initBoard(n)
	AddDotsP2(segments, board)
	print(board)
	part2 = DangerAreas(board)

	return part1, part2
}

func DangerAreas(board [][]int) int {
	sum := 0
	for r := 0; r < len(board); r++ {
		for c := 0; c < len(board); c++ {
			if board[r][c] > 1 {
				sum += 1
			}
		}
	}
	return sum
}

func initBoard(size int) [][]int {
	board := make([][]int, size)
	for i := 0; i < size; i++ {
		board[i] = make([]int, size)
		for ii := 0; ii < size; ii++ {
			board[i][ii] = dot
		}
	}
	return board
}

func parse(s string) []Segment {
	lines := strings.Split(s, "\n")
	segments := make([]Segment, 0)
	for _, line := range lines {
		var x1, x2, y1, y2 int
		pkg.MustScanf(line, "%d,%d -> %d,%d", &x1, &y1, &x2, &y2)
		s := Segment{Coord{x1, y1}, Coord{x2, y2}}
		segments = append(segments, s)
	}
	return segments
}

func print(board [][]int) {
	if len(board) > 20 {
		fmt.Println("board too large: ", len(board))
		return
	}

	s := ""
	delim := " "
	for _, row := range board {
		printRow := ""
		for _, item := range row {
			if item == dot {
				s += "."
			} else {
				s += strconv.Itoa(item)
			}
			s += delim
		}
		s += printRow + "\n"
	}
	fmt.Println(s)
}

func addVerticalSegment(segment Segment, board [][]int) error {
	e1, e2 := segment.E1, segment.E2
	var left, right Coord
	if e1.x < e2.x {
		left, right = e1, e2
	} else {
		right, left = e1, e2
	}

	for i := left.x; i <= right.x; i++ {
		board[e1.y][i] += 1
	}

	return nil
}

func addHorizontalSegment(segment Segment, board [][]int) error {
	e1, e2 := segment.E1, segment.E2
	var top, bottom Coord
	if e1.y < e2.y {
		bottom, top = e1, e2
	} else {
		top, bottom = e1, e2
	}

	for i := bottom.y; i <= top.y; i++ {
		board[i][e1.x] += 1
	}

	return nil
}

// 0,0 -> 8,8
// 8,0 -> 0,8  0,8 -> 8,0
func addDiagonalSegment(segment Segment, board [][]int) error {
	e1, e2 := segment.E1, segment.E2
	var left, right Coord
	if e1.x < e2.x {
		left, right = e1, e2
	} else {
		right, left = e1, e2
	}
	up := left.y < right.y
	for i := left.x; i <= right.x; i++ {
		diff := i - left.x
		if up {
			board[left.y+diff][i] += 1
		} else {
			board[left.y-diff][i] += 1
		}
	}

	return nil
}

func AddDotsP2(segments []Segment, board [][]int) {
	for i, _ := range segments {
		horizontal := segments[i].E1.x == segments[i].E2.x
		vertical := segments[i].E1.y == segments[i].E2.y
		nominator, denominator := (segments[i].E1.x - segments[i].E2.x), (segments[i].E1.y - segments[i].E2.y)
		diagonal := false
		if denominator != 0 {
			diagonal = (pkg.Abs(nominator/denominator) == 1)
		}
		if horizontal {
			addHorizontalSegment(segments[i], board)
		} else if vertical {
			addVerticalSegment(segments[i], board)
		} else if diagonal {
			err := addDiagonalSegment(segments[i], board)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			// fmt.Printf("skipping %v since not horiz/veritical\n", segments[i])
		}
	}
}

func AddDots(segments []Segment, board [][]int) {
	for i, _ := range segments {
		horizontal := segments[i].E1.x == segments[i].E2.x
		vertical := segments[i].E1.y == segments[i].E2.y
		if horizontal {
			err := addHorizontalSegment(segments[i], board)
			if err != nil {
				fmt.Println(err)
			}
		} else if vertical {
			err := addVerticalSegment(segments[i], board)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			// fmt.Printf("skipping %v since not horiz/veritical\n", segments[i])
		}
	}
}

func main() {
	year, day := pkg.GetAoVDate()
	fmt.Printf("Advent of Code %s Day %s\n", year, day)
	execute.Run(run, tests, puzzle, true)
}
