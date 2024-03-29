package main

import (
	"advent/pkg"
	"advent/pkg/execute"
	"fmt"
	"strconv"
	"strings"
)

const (
	dot  = "."
	mark = "#"
)

type Grid [][]string

type Fold struct {
	axis  string
	value int
}

type Coord struct {
	x, y int
}

func run(input string) (part1 interface{}, part2 interface{}) {
	// part 1
	grid, folds := parse(input)
	part1 = Part1(grid, folds)
	part2 = Part2(grid, folds) // answer is printed out, not returned
	return part1, part2
}

func Part2(grid Grid, folds []Fold) int {
	for _, fold := range folds {
		grid, _ = foldGrid(grid, fold)
	}
	fmt.Println(grid.String())
	return 0
}

func Part1(grid Grid, folds []Fold) int {
	_, markCount := foldGrid(grid, folds[0])
	return markCount
}

func foldGrid(grid Grid, fold Fold) (Grid, int) {
	var newGrid Grid

	if fold.axis == "y" {
		newGrid = make(Grid, fold.value)
		for y, _ := range grid {
			if y < fold.value {
				newGrid[y] = make([]string, len(grid[y]))
				for x, _ := range grid[y] {
					newGrid[y][x] = grid[y][x]
				}
			} else if y > fold.value {
				newY := fold.value - pkg.Abs((y - fold.value))
				for x, _ := range grid[y] {
					if grid[newY][x] == mark {
						newGrid[newY][x] = mark
					} else {
						newGrid[newY][x] = grid[y][x]
					}
				}
			}
		}
	} else {
		newGrid = make(Grid, len(grid))
		for y, _ := range grid {
			newGrid[y] = make([]string, fold.value)
			for x, _ := range grid[y] {
				if x < fold.value {
					newGrid[y][x] = grid[y][x]
				} else if x > fold.value {
					newX := fold.value - pkg.Abs((x - fold.value))
					if grid[y][newX] == mark {
						newGrid[y][newX] = mark
					} else {
						newGrid[y][newX] = grid[y][x]
					}
				}
			}
		}
	}

	markCount := 0
	for y, _ := range newGrid {
		for x, _ := range newGrid[y] {
			if newGrid[y][x] == mark {
				markCount++
			}
		}
	}

	return newGrid, markCount
}

func parse(s string) (Grid, []Fold) {
	lines := strings.Split(s, "\n")
	folds := make([]Fold, 0)
	coords := make([]Coord, 0)
	maxX, maxY := 0, 0
	for _, line := range lines {
		if strings.HasPrefix(line, "fold along") {
			s := strings.Split(line, "=")
			value, _ := strconv.Atoi(s[1])
			folds = append(folds, Fold{string(s[0][11]), value})
		} else if strings.Contains(line, ",") {
			s := strings.Split(line, ",")
			x, _ := strconv.Atoi(s[0])
			y, _ := strconv.Atoi(s[1])
			coords = append(coords, Coord{x, y})
			maxX = pkg.Max(maxX, x)
			maxY = pkg.Max(maxY, y)
		}
	}

	grid := make(Grid, maxY+1)
	for y, _ := range grid {
		grid[y] = make([]string, maxX+1)
		for x, _ := range grid[y] {
			grid[y][x] = dot
		}
	}

	for _, c := range coords {
		grid[c.y][c.x] = mark
	}

	return grid, folds
}

func (g *Grid) String() string {
	s := ""
	delim := ""
	for y, _ := range *g {
		// s += (*g)[y]
		for x, _ := range (*g)[y] {
			s += (*g)[y][x]
			s += delim
		}
		s += "\n"
	}
	return s
}

func main() {
	year, day := pkg.GetAoVDate()
	fmt.Printf("Advent of Code %s Day %s\n", year, day)
	execute.Run(run, tests, puzzle, true)
}
