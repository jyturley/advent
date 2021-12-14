package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestLowPoints(t *testing.T) {
	grid := parse(testpuzzle)
	coords, vals := findLowPoints(grid)
	require.Equalf(t, 4, len(vals), "bad length\n%v\n%v\n", vals, coords)
	require.Equalf(t, 4, len(coords), "bad length %v", coords)
	assert.Equal(t, 1, vals[0])
	assert.Equal(t, 1, coords[0].x, "bad coords")
	assert.Equal(t, 0, coords[0].y, "bad coords")
	assert.Equal(t, 0, vals[1])
	assert.Equal(t, 5, vals[2])
	assert.Equal(t, 5, vals[3])
}

func TestGetAdjCoords(t *testing.T) {
	grid := parse(testpuzzle)
	assert.Equal(t, 3, len(getAdjCoords(0, 1, grid)))
}
func TestGetBasin(t *testing.T) {
	grid := parse(testpuzzle)
	assert.Equal(t, 3, getBasin(Coord{0, 0}, grid))
	assert.Equal(t, 9, getBasin(Coord{9, 0}, grid))
	assert.Equal(t, 14, getBasin(Coord{2, 2}, grid))
}
