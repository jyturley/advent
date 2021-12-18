package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParse(t *testing.T) {
	fmt.Println("testing")
	s := parse(miniExample)
	assert.Equal(t, "110100101111111000101000", s, miniExample)
}

func TestPart1(t *testing.T) {
	input := "38006F45291200"
	s := parse(input)
	assert.Equal(t, "110100101111111000101000", Part1(s), input)
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
