package main

import (
	"advent/pkg"
	"advent/pkg/execute"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Coord struct {
	x, y int
}

func run(input string) (part1 interface{}, part2 interface{}) {
	// part 1
	grid := parse(input)
	part1 = Part1(grid)
	part2 = Part2(grid)
	return part1, part2
}

func Part1(grid [][]int) int {
	_, lowVals := findLowPoints(grid)
	return pkg.Sum(lowVals...) + len(lowVals)
}

func Part2(grid [][]int) int {
	lowCoords, _ := findLowPoints(grid)
	basinSizes := getBasins(lowCoords, grid)
	sort.Sort(sort.Reverse(sort.IntSlice(basinSizes)))
	fmt.Println(basinSizes)
	if len(basinSizes) < 3 {
		return -1
	}

	return basinSizes[0] * basinSizes[1] * basinSizes[2]
}

func getBasins(lowCoords []Coord, grid [][]int) []int {
	out := make([]int, 0)
	for _, lc := range lowCoords {
		out = append(out, getBasin(lc, grid))
	}
	return out
}

func getBasin(lc Coord, grid [][]int) int {
	// fmt.Println("getBasin", lc)
	basinCoords := make([]Coord, 0)
	seen := map[Coord]bool{}
	q := make([]Coord, 0)
	q = append(q, lc)
	for len(q) > 0 {
		c := q[0]
		q = q[1:]
		if seen, ok := seen[c]; ok && seen {
			continue
		}

		seen[c] = true
		if grid[c.y][c.x] == 9 {
			continue
		}
		basinCoords = append(basinCoords, c)

		// add to queue
		for _, neighbor := range getAdjCoords(c.x, c.y, grid) {
			if seen, ok := seen[neighbor]; ok && seen {
				continue
			}
			q = append(q, neighbor)
		}

	}

	// fmt.Println(basinCoords)
	return len(basinCoords)
}

func findLowPoints(grid [][]int) ([]Coord, []int) {
	coords := make([]Coord, 0)
	values := make([]int, 0)
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			currentHeight := grid[y][x]
			neighbors := getAdjCoords(x, y, grid)
			smallest := true
			for _, c := range neighbors {
				if grid[c.y][c.x] <= currentHeight {
					smallest = false
				}
			}

			if smallest {
				coords = append(coords, Coord{x, y})
				values = append(values, currentHeight)
			}
		}
	}
	fmt.Println(coords)
	fmt.Println(values)
	return coords, values
}

func getAdjCoords(x, y int, grid [][]int) []Coord {
	out := make([]Coord, 0)
	// top
	if 0 <= y-1 {
		out = append(out, Coord{x, y - 1})
	}

	// bottom
	if y+1 < len(grid) {
		out = append(out, Coord{x, y + 1})
	}

	// left
	if 0 <= x-1 {
		out = append(out, Coord{x - 1, y})
	}

	// right
	if x+1 < len(grid[y]) {
		out = append(out, Coord{x + 1, y})
	}

	return out
}

func parse(s string) [][]int {
	out := [][]int{}
	lines := strings.Split(s, "\n")
	for _, line := range lines {
		heights := []int{}
		s := strings.Split(line, "")
		for i := 0; i < len(line); i++ {
			if h, err := strconv.Atoi(s[i]); err == nil {
				heights = append(heights, h)
			} else {
				fmt.Println(err)
			}
			// heights[i] = h
		}
		out = append(out, heights)
	}
	return out
}

func main() {
	year, day := pkg.GetAoVDate()
	fmt.Printf("Advent of Code %s Day %s\n", year, day)
	execute.Run(run, tests, puzzle, true)
}
