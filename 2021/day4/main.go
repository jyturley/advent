package main

import (
	"advent/pkg"
	"advent/pkg/execute"
	"fmt"
	"strconv"
	"strings"
)

const (
	marked = -9
)

type BingoBoard [5][5]int

func run(input string) (part1 interface{}, part2 interface{}) {
	// part 1
	draw, boards := parse(input)
	for _, num := range draw {
		if winner, _ := Announce(boards, num); winner != nil {
			part1 = winner.calculateScore(num)
			break
		}
	}

	// part 2
	draw, boards = parse(input) // reset board
	finalNum := -1
	for _, num := range draw {
		if winners := AnnouncePt2(boards, num); len(winners) > 0 {
			if len(boards) > 1 {
				boards = removeWinners(boards, winners)
			} else {
				finalNum = num
				break
			}
		}
	}
	part2 = boards[0].calculateScore(finalNum)
	return part1, part2
}

func remove(boards []BingoBoard, idx int) []BingoBoard {
	return append(boards[:idx], boards[idx+1:]...)
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

func removeWinners(boards []BingoBoard, winners []int) []BingoBoard {
	keepBoards := make([]BingoBoard, 0)
	for i, board := range boards {
		isIndexAWinner := false
		for _, winIdx := range winners {
			if i == winIdx {
				isIndexAWinner = true
				break
			}
		}

		if !isIndexAWinner {
			keepBoards = append(keepBoards, board)
		}

	}
	return keepBoards
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

func Announce(boards []BingoBoard, num int) (*BingoBoard, int) {
	for i, _ := range boards {
		if win := MarkAndCheck(&boards[i], num); win {
			for ii, _ := range boards {
				if ii == i {
					continue
				}
				Mark(&boards[ii], num)
			}
			return &boards[i], i
		}
	}
	return nil, 0
}

func AnnouncePt2(boards []BingoBoard, num int) []int {
	existWinners := false
	coords := map[int][]int{}
	for i, _ := range boards {
		r, c, win := Mark(&boards[i], num)
		if win {
			coords[i] = []int{r, c}
			existWinners = true
		}
	}

	if !existWinners {
		return nil
	}

	winners := make([]int, 0)

	for i, _ := range boards {
		if v, ok := coords[i]; ok {
			r, c := v[0], v[1]
			if win := boards[i].Check(r, c); win {
				winners = append(winners, i)
			}

		}
	}
	return winners
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
