package main

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestBinToInt(t *testing.T) {
	assert.Equal(t, 22, binToInt([]int{1, 0, 1, 1, 0}))
	assert.Equal(t, 9, binToInt([]int{0, 1, 0, 0, 1}))
	assert.Equal(t, 23, binToInt([]int{1, 0, 1, 1, 1}))
	assert.Equal(t, 10, binToInt([]int{0, 1, 0, 1, 0}))
}

func TestGetRating(t *testing.T) {
	binNums := strings.Split(testpuzzle, "\n")
	most, least := getMostLeastCommonDigitsPart2(binNums)
	assert.Equal(t, 23, getRating(binNums, most))
	assert.Equal(t, 10, getRating(binNums, least))
}
