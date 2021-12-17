package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHighestPointInTrajectory(t *testing.T) {
	fmt.Println("testing")
	ta := parse(testpuzzle)
	assert.Equal(t, 45, highestPointInTrajectory(Velocity{6, 9}, ta))
	assert.Equal(t, 3, highestPointInTrajectory(Velocity{7, 2}, ta))
	assert.Equal(t, -1, highestPointInTrajectory(Velocity{17, -4}, ta))
}

func TestPart2(t *testing.T) {
	ta1 := parse(testpuzzle)
	ta2 := parse(puzzle)
	assert.Equal(t, 112, Part2(ta1))
	fmt.Println()
	assert.Equal(t, 112, Part2(ta2))
}

func TestPositionInTargetArea(t *testing.T) {
	ta := parse(testpuzzle)
	tests := []struct {
		inCoord      Coord
		inTargetArea TargetArea
		expect       bool
	}{
		{Coord{6, 9}, ta, false},
		{Coord{20, 9}, ta, false},
		{Coord{20, -10}, ta, true},
		{Coord{20, -8}, ta, true},
		{Coord{20, -1}, ta, false},
	}

	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			assert.Equal(t, tt.expect, positionInTargetArea(tt.inCoord, tt.inTargetArea), tt.inTargetArea)
		})
	}
}
