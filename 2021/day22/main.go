package main

import (
	"advent/pkg"
	"advent/pkg/execute"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Cuboid struct {
	x1, x2, y1, y2, z1, z2 int
}

type Step struct {
	on         bool
	directions Cuboid
}

func NewCuboid(x1, x2, y1, y2, z1, z2 int) Cuboid {
	return Cuboid{
		x1: pkg.Min(x1, x2),
		x2: pkg.Max(x1, x2),
		y1: pkg.Min(y1, y2),
		y2: pkg.Max(y1, y2),
		z1: pkg.Min(z1, z2),
		z2: pkg.Max(z1, z2),
	}
}

func run(input string) (part1 interface{}, part2 interface{}) {
	// part 1
	steps, maxmin := parse(input)
	part1 = Part1(steps, maxmin)
	return part1, part2
}

func Part1(steps []Step, maxmin [6]int) int {
	size := 1000
	grid := make([][][]bool, size)
	for x := 0; x < len(grid); x++ {
		grid[x] = make([][]bool, size)
		for y := 0; y < len(grid[x]); y++ {
			grid[x][y] = make([]bool, size)
		}
	}
	// grid := make([][][]bool, pkg.Abs(maxmin[0]-maxmin[1]))
	// for x := 0; x < len(grid); x++ {
	// 	grid[x] = make([][]bool, pkg.Abs(maxmin[2]-maxmin[3]))
	// 	for y := 0; y < len(grid[x]); y++ {
	// 		grid[x][y] = make([]bool, pkg.Abs(maxmin[4]-maxmin[5]))
	// 	}
	// }
	fmt.Println(maxmin)
	fmt.Println(len(grid), len(grid[0]), len(grid[0][0]))
	onCount := 0
	for i, step := range steps {
		onCount = follow(step, &grid, maxmin)
		fmt.Printf("After step %d, on count: %d\n", i+1, onCount)
	}

	return onCount
}

func Part2() int {
	out := 0
	return out
}

func follow(step Step, grid *[][][]bool, maxmin [6]int) int {
	onCount := 0
	adjustX := 0
	if maxmin[1] < 0 {
		adjustX = -maxmin[1]
	}
	adjustY := 0
	if maxmin[3] < 0 {
		adjustY = -maxmin[3]
	}
	adjustZ := 0
	if maxmin[5] < 0 {
		adjustZ = -maxmin[4]
	}

	for x := step.directions.x1; x <= step.directions.x2; x++ {
		for y := step.directions.y1; y <= step.directions.y2; y++ {
			for z := step.directions.z1; z <= step.directions.z2; z++ {
				(*grid)[x+adjustX][y+adjustY][z+adjustZ] = step.on
			}
		}
	}

	for x := 0; x < len(*grid); x++ {
		for y := 0; y < len((*grid)[0]); y++ {
			for z := 0; z < len((*grid)[0][0]); z++ {
				if (*grid)[x][y][z] {
					onCount++
				}
			}
		}
	}
	return onCount
}

func parse(s string) ([]Step, [6]int) {
	steps := make([]Step, 0)
	maxmin := [6]int{}
	lines := strings.Split(s, "\n")
	re := regexp.MustCompile(`(on|off) x=(-?[0-9]+)\.\.(-?[0-9]+),y=(-?[0-9]+)\.\.(-?[0-9]+),z=(-?[0-9]+)\.\.(-?[0-9]+)`)
	for _, line := range lines {
		m := re.FindStringSubmatch(line)
		on := m[1] == "on"
		x1, _ := strconv.Atoi(m[2])
		x2, _ := strconv.Atoi(m[3])
		y1, _ := strconv.Atoi(m[4])
		y2, _ := strconv.Atoi(m[5])
		z1, _ := strconv.Atoi(m[6])
		z2, _ := strconv.Atoi(m[7])
		step := Step{on, NewCuboid(x1, x2, y1, y2, z1, z2)}
		steps = append(steps, step)
		maxmin[0] = pkg.Max(x1, x2, maxmin[0])
		maxmin[1] = pkg.Min(x1, x2, maxmin[1])
		maxmin[2] = pkg.Max(y1, y2, maxmin[2])
		maxmin[3] = pkg.Min(y1, y2, maxmin[3])
		maxmin[4] = pkg.Max(z1, z2, maxmin[4])
		maxmin[5] = pkg.Min(z1, z2, maxmin[5])
	}
	return steps, maxmin
}

func main() {
	year, day := pkg.GetAoVDate()
	fmt.Printf("Advent of Code %s Day %s\n", year, day)
	execute.Run(run, tests, puzzle, true)
}
