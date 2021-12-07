package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFishAfterDays(t *testing.T) {
	fish := parse(testpuzzle)
	assert.Equal(t, 26, FishAfterDays(18, fish))

	fish = parse(testpuzzle)
	assert.Equal(t, 26, FishAfterManyDays(18, fish))
}
