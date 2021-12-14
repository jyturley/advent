package main

import (
	// "fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPart2N(t *testing.T) {
	template, rules := parse(testpuzzle)
	Part2N(template, rules, 3)
	assert.Equal(t, true, true)
}

func TestStepCount(t *testing.T) {
	template, rules := parse(testpuzzle)
	t1 := stepN(template, 1, rules)
	t2 := stepN(template, 2, rules)
	pairCounts := stepCounts(getPairCountDict(t1), rules)
	assert.Truef(t, len(getPairCountDict(t2)) == len(pairCounts), "%s -> %s\ngot:  %v\nwant: %v\n", t1, t2, pairCounts, getPairCountDict(t2))
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
