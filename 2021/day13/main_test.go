package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFoldGrid(t *testing.T) {
	grid, folds := parse(testpuzzle)
	fmt.Println(grid.String())
	fmt.Println()
	grid, markCount := foldGrid(grid, folds[0])
	assert.Equalf(t, 17, markCount, "%s %d\n", grid.String(), markCount)
	// grid, markCount = foldGrid(grid, folds[1])
	// assert.Equalf(t, 17, markCount, "%s %d\n", grid.String(), markCount)
}

// func Test(t *testing.T) {
// 	tests := []struct {
// 		in       int
// 		want     int
// 		outError bool
// 	}{
// 		{0, 0, false},
// 	}

// 	for _, tt := range tests {
// 		t.Run("test", func(t *testing.T) {
// 			assert.Equal(t, tt.want, somefunc(tt.in))
// 		})
// 	}
// }
