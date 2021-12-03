package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBinToInt(t *testing.T) {
	assert.Equal(t, 22, binToInt([]int{1, 0, 1, 1, 0}))
	assert.Equal(t, 9, binToInt([]int{0, 1, 0, 0, 1}))
}
