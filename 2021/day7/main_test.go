package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindAlignment(t *testing.T) {
	// positions := parse(testpuzzle)
	assert.Equal(t, true, true)
}

func TestSummation(t *testing.T) {
	assert.Equal(t, 66, summation(11))
}

func TestCrabWalk(t *testing.T) {
	tests := []struct {
		inFrom, inTo int
		want         int
	}{
		{16, 5, 66},
		{1, 5, 10},
		{14, 5, 45},
	}

	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			assert.Equal(t, tt.want, crabWalkCost(tt.inFrom, tt.inTo))
		})
	}
}
