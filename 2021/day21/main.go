package main

import (
	"advent/pkg"
	"advent/pkg/execute"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const (
	winscore = 1000
)

type Board struct {
	data   [10]int
	p1, p2 int
}

func NewBoard(p1Start, p2Start int) *Board {
	b := &Board{
		p1: p1Start - 1,
		p2: p2Start - 1,
	}
	b.data[p1Start-1] = 1
	b.data[p2Start-1] = 2

	return b
}

func (b *Board) AdvanceP1(num int) int {
	startPos := b.p1
	size := len(b.data)
	netAdvance := num % size
	for i := 0; i < netAdvance; i++ {
		b.p1 += 1
		if b.p1 == size {
			b.p1 = 0
		}
	}
	b.data[b.p1] = 1
	b.data[startPos] = 0 //clear previous pos
	return b.P1()
}

func (b *Board) AdvanceP2(num int) int {
	startPos := b.p2
	size := len(b.data)
	netAdvance := num % size
	for i := 0; i < netAdvance; i++ {
		b.p2 += 1
		if b.p2 == size {
			b.p2 = 0
		}
	}
	b.data[b.p2] = 1
	b.data[startPos] = 0 //clear previous pos
	return b.P2()
}

func (b *Board) P1() int {
	return b.p1 + 1
}

func (b *Board) P2() int {
	return b.p2 + 1
}

type Dice interface {
	Roll() int
	RollCount() int
}

type DeterministicDice struct {
	currentVal int
	count      int
}

func (dd *DeterministicDice) Roll() int {
	dd.currentVal += 1
	if dd.currentVal == 101 {
		dd.currentVal = 1
	}
	dd.count++
	return dd.currentVal
}
func (dd *DeterministicDice) RollCount() int {
	return dd.count
}

type NormalDice struct {
	count int
}

func (nd *NormalDice) Roll() int {
	nd.count++
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(100) + 1
}
func (nd *NormalDice) RollCount() int {
	return nd.count
}

func run(input string) (part1 interface{}, part2 interface{}) {
	// part 1
	pos1, pos2 := parse(input)
	part1 = Part1(pos1, pos2)
	return part1, part2
}

func Part1(pos1, pos2 int) int {
	out := 0
	dice := DeterministicDice{}
	out = playGame(NewBoard(pos1, pos2), &dice)
	return out
}

func playGame(board *Board, d Dice) int {
	var score1, score2 int
	for score1 < winscore && score2 < winscore {
		// player 1
		n := d.Roll()
		n1 := d.Roll()
		n2 := d.Roll()
		endingSpace := board.AdvanceP1(n + n1 + n2)
		score1 += endingSpace
		// fmt.Printf("Player 1 rolls %d+%d+%d and moves to space %d for a total score of %d.\n", n, n1, n2, endingSpace, score1)
		if score1 >= winscore {
			break
		}

		// player 2
		n = d.Roll()
		n1 = d.Roll()
		n2 = d.Roll()
		endingSpace = board.AdvanceP2(n + n1 + n2)
		score2 += endingSpace
		// fmt.Printf("Player 2 rolls %d+%d+%d and moves to space %d for a total score of %d.\n", n, n1, n2, endingSpace, score2)
	}

	var out int
	fmt.Printf("RollCount %d\n", d.RollCount())
	if score1 > score2 {
		fmt.Printf("Player1 wins with score of %d, over %d\n", score1, score2)
		out = score2 * d.RollCount()
	} else {
		fmt.Printf("Player2 wins with score of %d, over %d\n", score2, score1)
		out = score1 * d.RollCount()
	}
	return out
}

func Part2() int {
	out := 0
	return out
}

func parse(s string) (p1, p2 int) {
	var pos1, pos2 int
	lines := strings.Split(s, "\n")
	pkg.MustScanf(lines[0], "Player 1 starting position: %d", &pos1)
	pkg.MustScanf(lines[1], "Player 2 starting position: %d", &pos2)
	return pos1, pos2
}

func main() {
	year, day := pkg.GetAoVDate()
	fmt.Printf("Advent of Code %s Day %s\n", year, day)
	execute.Run(run, tests, puzzle, true)
}
