package main

import (
	"advent/pkg"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDiagonal(t *testing.T) {
	board := initBoard(10)
	segment := Segment{Coord{0, 0}, Coord{8, 8}}
	print(board)
	assert.NoError(t, addDiagonalSegment(segment, board))
	print(board)
	segment = Segment{Coord{0, 9}, Coord{2, 7}}
	assert.NoError(t, addDiagonalSegment(segment, board))
	print(board)
	segment = Segment{Coord{8, 0}, Coord{0, 8}}
	assert.NoError(t, addDiagonalSegment(segment, board))
	print(board)

	tests := []struct {
		Expect   bool
		inCoord1 Coord
		inCoord2 Coord
	}{
		{Expect: true, inCoord1: Coord{0, 0}, inCoord2: Coord{8, 8}},
		{Expect: true, inCoord1: Coord{1, 1}, inCoord2: Coord{9, 9}},
		{Expect: true, inCoord1: Coord{5, 5}, inCoord2: Coord{0, 0}},
		{Expect: true, inCoord1: Coord{6, 6}, inCoord2: Coord{1, 1}},
		{Expect: true, inCoord1: Coord{0, 9}, inCoord2: Coord{2, 7}},
	}

	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			// isDiagonal := (tt.inCoord1.x - tt.inCoord2.x) == (tt.inCoord1.y - tt.inCoord2.y)
			slope := pkg.Abs((tt.inCoord1.x - tt.inCoord2.x) / (tt.inCoord1.y - tt.inCoord2.y))
			isDiagonal := (slope == 1)
			assert.Equalf(t, tt.Expect, isDiagonal, "incorrect slope %v", slope)
		})
	}

}
