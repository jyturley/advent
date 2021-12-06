package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCheck(t *testing.T) {
	nums, boards := parse(testpuzzle)
	winner := boards[2]
	for i, n := range nums {
		if i > 11 {
			continue
		}
		assert.Falsef(t, winner.Check(0, 3), "bad check at num %d", n)
		Mark(&winner, n)
	}
	assert.True(t, winner.Check(0, 3))
}

func TestCalculateScore(t *testing.T) {
	nums, boards := parse(testpuzzle)
	winner := boards[2]
	for i, n := range nums {
		if i > 11 {
			continue
		}
		Mark(&winner, n)
	}
	assert.Equal(t, 4512, winner.calculateScore(24))
}
