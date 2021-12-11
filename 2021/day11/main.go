package main

import (
	"advent/pkg"
	"advent/pkg/execute"
	"fmt"
	"strconv"
	"strings"
)

type Grid [10][10]int
type Coord struct {
	x, y int
}

func run(input string) (part1 interface{}, part2 interface{}) {
	// part 1
	grid := parse(input)
	part1 = Part1(grid)
	return part1, part2
}

func Part1(grid Grid) int {
	return stepN(100, &grid)
}

func stepN(n int, grid *Grid) int {
	numFlashes := 0
	for i := 0; i < n; i++ {
		numFlashes += step(grid)
	}
	return numFlashes
}

func step(grid *Grid) int {
	q := make([]Coord, 0)
	// +1 every celll
	for y, _ := range grid {
		for x, _ := range grid {
			(*grid)[y][x] += 1
			if (*grid)[y][x] == 10 {
				q = append(q, Coord{x, y})
			}
		}
	}

	// flash
	seen := map[Coord]bool{}
	for len(q) > 0 {
		current := q[0]
		q = q[1:]
		// fmt.Printf("val: %d at (%d, %d)\n", (*grid)[current.y][current.x], current.x, current.y)
		if _, ok := seen[current]; ok {
			// fmt.Println("skipped")
			continue
		}
		seen[current] = true
		for _, neighbor := range getNeighbors(current, grid) {
			x, y := neighbor.x, neighbor.y
			(*grid)[y][x] += 1
			if (*grid)[y][x] == 10 {
				q = append(q, Coord{x, y})
			}
		}
	}

	numFlashes := 0
	// set back flashes to zero
	for y, _ := range grid {
		for x, _ := range grid {
			if (*grid)[y][x] > 9 {
				numFlashes++
				(*grid)[y][x] = 0
			}
		}
	}
	return numFlashes
}

func getNeighbors(current Coord, grid *Grid) []Coord {
	out := make([]Coord, 0)
	x, y := current.x, current.y
	if 0 <= x-1 && 0 <= y-1 {
		out = append(out, Coord{x - 1, y - 1})
	}
	if 0 <= y-1 {
		out = append(out, Coord{x, y - 1})
	}
	if x+1 < 10 && 0 <= y-1 {
		out = append(out, Coord{x + 1, y - 1})
	}
	if x+1 < 10 {
		out = append(out, Coord{x + 1, y})
	}
	if x+1 < 10 && y+1 < 10 {
		out = append(out, Coord{x + 1, y + 1})
	}
	if y+1 < 10 {
		out = append(out, Coord{x, y + 1})
	}
	if 0 <= x-1 && y+1 < 10 {
		out = append(out, Coord{x - 1, y + 1})
	}
	if 0 <= x-1 {
		out = append(out, Coord{x - 1, y})
	}

	return out
}

func SideBySide(g1, g2 Grid) string {
	s := ""
	delim := ","
	for y, _ := range g1 {
		row := ""
		for x, _ := range g1 {
			n := strconv.Itoa(g1[y][x])
			row += n
			row += delim
			if len(n) == 1 {
				row += " "
			}
		}
		lenNormalRow := 20 + 10*len(delim)
		if len(row) > lenNormalRow {
			diff := len(row) - lenNormalRow
			for i := 0; i < 5-diff; i++ {
				row += " "
			}
		} else {
			row += "   "
		}
		for x, _ := range g2 {
			n := strconv.Itoa(g2[y][x])
			row += n
			row += delim
			if len(n) == 1 {
				row += " "
			}
		}
		s += row + "\n"
	}
	return s
}

func (g *Grid) String() string {
	s := ""
	delim := ","
	for y, _ := range g {
		for x, _ := range g {
			s += strconv.Itoa(g[y][x])
			s += delim
		}
		s += "\n"
	}
	return s
}

func parse(s string) Grid {
	lines := strings.Split(s, "\n")
	out := Grid{}
	for i, line := range lines {
		row := [10]int{}
		for ii, c := range strings.Split(line, "") {
			n, _ := strconv.Atoi(c)
			row[ii] = n
		}
		out[i] = row
	}

	return out
}

func main() {
	year, day := pkg.GetAoVDate()
	fmt.Printf("Advent of Code %s Day %s\n", year, day)
	execute.Run(run, tests, puzzle, true)
}
