package main

import (
	"advent/pkg"
	"advent/pkg/execute"
	"fmt"
	"strconv"
	"strings"
)

const (
	marked = -999
)

type BingoBoard [5][5]int

func run(input string) (part1 interface{}, part2 interface{}) {
	// part 1
	score := 0
	draw, boards := parse(input)
	for _, num := range draw {
		if winner := Announce(boards, num); winner != nil {
			score = winner.calculateScore(num)
			break
		}
	}

	part1 = score

	return part1, part2
}

func parse(s string) ([]int, []BingoBoard) {
	lines := strings.Split(s, "\n")
	draws := make([]int, 0)
	boards := make([]BingoBoard, 0)
	for ln, line := range lines {
		if ln == 0 {
			for _, s := range strings.Split(line, ",") {
				i, _ := strconv.Atoi(s)
				draws = append(draws, i)
			}
		} else if (ln-2)%6 == 0 { // start of each bingo board
			bb := BingoBoard{}
			for i := 0; i < 5; i++ {
				bb.AddNums(lines[ln+i], i)
			}

			boards = append(boards, bb)
		}
	}
	return draws, boards
}

func Mark(bingo *BingoBoard, num int) (int, int, bool) {
	bb := *bingo
	for r, _ := range bb {
		for c, _ := range bb {
			if bb[r][c] == num {
				(*bingo)[r][c] = marked
				return r, c, true
			}
		}
	}
	return 0, 0, false
}

func MarkAndCheck(bb *BingoBoard, num int) bool {
	r, c, found := Mark(bb, num)
	if found {
		return bb.Check(r, c)
	}

	return false
}

func Announce(boards []BingoBoard, num int) *BingoBoard {
	for i, _ := range boards {
		if win := MarkAndCheck(&boards[i], num); win {
			return &boards[i]
		}
	}
	return nil
}

func (bb *BingoBoard) Check(r, c int) bool {
	fullRowMarked, fullColumnMarked := true, true

	for i := 0; i < 5; i++ {
		if bb[r][i] != marked {
			fullRowMarked = false
			break
		}
	}

	for i := 0; i < 5; i++ {
		if bb[i][c] != marked {
			fullColumnMarked = false
			break
		}
	}

	return fullRowMarked || fullColumnMarked
}

func (bb *BingoBoard) Print() {
	for _, row := range bb {
		fmt.Println(row)
	}
}

func (bb *BingoBoard) AddNums(line string, bbRow int) {
	column := 0
	for _, s := range strings.Split(line, " ") {
		if s == "" || s == " " {
			continue
		}
		i, err := strconv.Atoi(s)
		if err != nil {
			fmt.Println(err)
		}
		bb[bbRow][column] = i
		column++
	}
}

func (bb *BingoBoard) calculateScore(winningNum int) int {
	sum := 0
	for r, _ := range bb {
		for c, _ := range bb {
			i := bb[r][c]
			if i != marked {
				sum += i
			}
		}
	}
	return sum * winningNum
}

func main() {
	year, day := pkg.GetAoVDate()
	fmt.Printf("Advent of Code %s Day %s\n", year, day)
	execute.Run(run, tests, puzzle, true)
}
