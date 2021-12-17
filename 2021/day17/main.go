package main

import (
	"advent/pkg"
	"advent/pkg/execute"
	"fmt"
	"strconv"
	"strings"
)

var (
	startPos = Coord{0, 0}
	badValue = -999999
)

type Coord struct {
	x, y int
}

type Velocity struct {
	x, y int
}

type TargetArea struct {
	x1, x2 int
	y1, y2 int
}

func run(input string) (part1 interface{}, part2 interface{}) {
	// part 1
	ta := parse(input)
	part1 = Part1(ta)
	part2 = Part2(ta)
	return part1, part2
}

func Part2(ta TargetArea) int {
	goodVelocities := make([]Velocity, 0)
	for x := 0; x < 1000; x++ {
		for y := -500; y < 500; y++ {
			v := Velocity{x, y}
			height := highestPointInTrajectory(v, ta)
			if height != badValue {
				goodVelocities = append(goodVelocities, v)
			}
		}
	}
	return len(goodVelocities)
}

func Part1(ta TargetArea) int {
	maxHeight := -100000
	for x := 0; x < 500; x++ {
		for y := 0; y < 500; y++ {
			v := Velocity{x, y}
			maxHeight = pkg.Max(maxHeight, highestPointInTrajectory(v, ta))
		}
	}
	return maxHeight
}

func highestPointInTrajectory(initial Velocity, ta TargetArea) int {
	pos := startPos
	velocity := initial
	maxY := -1000000
	for !positionInTargetArea(pos, ta) {
		pos, velocity = nextStep(pos, velocity)
		maxY = pkg.Max(maxY, pos.y)
		if pastTargetArea(pos, ta) {
			return badValue
		}
	}

	return maxY
}

func pastTargetArea(pos Coord, ta TargetArea) bool {
	if pos.y < ta.y1 {
		return true
	}

	return false
}

func positionInTargetArea(pos Coord, ta TargetArea) bool {
	if pos.x < ta.x1 || pos.x > ta.x2 {
		return false
	}

	if pos.y < ta.y1 || pos.y > ta.y2 {
		return false
	}

	return true
}

func nextStep(cc Coord, cv Velocity) (Coord, Velocity) {
	newPos := cc
	newVelocity := cv

	newPos.x += cv.x
	newPos.y += cv.y

	if cv.x > 0 {
		newVelocity.x -= 1
	} else if cv.x < 0 {
		newPos.x += 1
	}

	newVelocity.y -= 1
	return newPos, newVelocity
}

func parse(s string) TargetArea {
	var sx, sy string
	pkg.MustScanf(s, "target area: %s %s", &sx, &sy)
	x := strings.Split(sx[2:len(sx)-1], "..")
	y := strings.Split(sy[2:], "..")
	x0, _ := strconv.Atoi(x[0])
	x1, _ := strconv.Atoi(x[1])
	y0, _ := strconv.Atoi(y[0])
	y1, _ := strconv.Atoi(y[1])
	return TargetArea{pkg.Min(x0, x1), pkg.Max(x0, x1), pkg.Min(y0, y1), pkg.Max(y0, y1)}
}

func main() {
	year, day := pkg.GetAoVDate()
	fmt.Printf("Advent of Code %s Day %s\n", year, day)
	execute.Run(run, tests, puzzle, true)
}
