package main

import (
	"advent/pkg"
	"advent/pkg/execute"
	"fmt"
	// "strings"
)

type Grid [][]int
type Coord struct {
	x, y int
}

var (
	badcoord = Coord{-1, -1}
)

func run(input string) (part1 interface{}, part2 interface{}) {
	// part 1
	grid := pkg.ParseIntGrid(input)
	part1 = Part1(grid)
	// if len(input) > 1000 {
	// 	return 0, 0
	// }
	return part1, part2
}

func Part1(grid Grid) int {
	out := 0
	q := make([]Coord, 0)
	dist := map[Coord]int{}
	prev := map[Coord]Coord{}
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid); x++ {
			c := Coord{x, y}
			dist[c] = 100000
			// prev[c] = badcoord
			q = append(q, c)
		}
	}

	seen := map[Coord]bool{}

	start, end := Coord{0, 0}, Coord{len(grid) - 1, len(grid[0]) - 1}
	dist[start] = 0

	for len(q) > 0 {
		current, idx := minDist(dist, q, grid)
		// fmt.Println(len(q), current)
		q = remove(q, idx)
		if seen[current] {
			continue
		}

		seen[current] = true

		for _, neighbor := range getNeighbors(current, grid) {
			if !inQ(neighbor, q) {
				continue
			}
			distance := dist[current] + grid[neighbor.y][neighbor.x]
			// fmt.Println("distance:", distance)
			if distance < dist[neighbor] {
				dist[neighbor] = distance
				prev[neighbor] = current
			}
		}
	}

	fmt.Println("done")
	// if len(grid) < 100 {
	// 	Print(dist)
	// 	fmt.Println(prev)
	// }

	sequence := make([]Coord, 0)
	node := end
	sequence = append(sequence, node)
	if _, ok := prev[end]; ok {
		ok := true
		for ok {
			node, _ = prev[node]
			sequence = append(sequence, node)
			node = prev[node]
			_, ok = prev[node]
		}
	}
	fmt.Println(sequence)

	for i := len(sequence) - 1; i >= 0; i-- {
		x, y := sequence[i].x, sequence[i].y
		fmt.Printf("%d, %d = %d\n", x, y, grid[y][x])
		out += grid[y][x]
	}

	return out - grid[0][0]
}

func inQ(neighbor Coord, q []Coord) bool {
	for _, c := range q {
		if c == neighbor {
			return true
		}
	}

	return false
}

func Print(dist map[Coord]int) {
	for k, v := range dist {
		fmt.Printf("dist: %d, %d = %d\n", k.x, k.y, v)
	}
}

func getNeighbors(current Coord, grid Grid) []Coord {
	lenY, lenX := len(grid), len(grid[0])
	out := make([]Coord, 0)
	x, y := current.x, current.y
	if 0 <= y-1 {
		out = append(out, Coord{x, y - 1})
	}
	if x+1 < lenX {
		out = append(out, Coord{x + 1, y})
	}
	if y+1 < lenY {
		out = append(out, Coord{x, y + 1})
	}
	if 0 <= x-1 {
		out = append(out, Coord{x - 1, y})
	}
	return out
}

func remove(q []Coord, idx int) []Coord {
	return append(q[:idx], q[idx+1:]...)
}

func minDist(dist map[Coord]int, q []Coord, grid Grid) (Coord, int) {
	var minCoord = q[0]
	var minCoordIdx = 0
	for i := 1; i < len(q); i++ {
		right := q[i]
		if dist[right] < dist[minCoord] {
			minCoord = right
			minCoordIdx = i
		}
	}

	return minCoord, minCoordIdx
}

func main() {
	year, day := pkg.GetAoVDate()
	fmt.Printf("Advent of Code %s Day %s\n", year, day)
	execute.Run(run, tests, puzzle, true)
}
