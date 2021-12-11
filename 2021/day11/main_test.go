package main

import (
	// "fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStepNew(t *testing.T) {
	tests := []struct {
		in         string
		want       string
		numSteps   int
		numFlashes int
	}{
		{
			`0000000000
0000000000
0000000000
0001111100
0001999100
0001919100
0001999100
0001111100
0000000000
0000000000`,
			`1111111111
1111111111
1111111111
1113454311
1114000411
1115000511
1114000411
1113454311
1111111111
1111111111`, 1,
			0},
		{
			testpuzzle,
			`6594254334
3856965822
6375667284
7252447257
7468496589
5278635756
3287952832
7993992245
5957959665
6394862637`, 1,
			0},
		{testpuzzle,
			`8807476555
5089087054
8597889608
8485769600
8700908800
6600088989
6800005943
0000007456
9000000876
8700006848`, 2, 0},
		{testpuzzle, `0481112976
0031112009
0041112504
0081111406
0099111306
0093511233
0442361130
5532252350
0532250600
0032240000`, 10, 204},
	}

	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			start := parse(tt.in)
			want := parse(tt.want)
			flashCount := stepN(tt.numSteps, &start)
			assert.Truef(t, deepEqual(want, start), "Correct From To:\n%s\nYour From To:\n%s\n", SideBySide(parse(tt.in), parse(tt.want)), SideBySide(parse(tt.in), start))
			if tt.numFlashes > 0 {
				assert.Equal(t, tt.numFlashes, flashCount)
			}
		})
	}
}

func deepEqual(g1, g2 Grid) bool {
	for y, _ := range g1 {
		for x, _ := range g1 {
			if g1[y][x] != g2[y][x] {
				return false
			}
		}
	}
	return true
}
