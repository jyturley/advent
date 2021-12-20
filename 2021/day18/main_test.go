package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStrToPair(t *testing.T) {
	fmt.Println("testing")
	// positions := parse(testpuzzle)
	testPair := Pair{}
	assert.Equal(t, testPair, StrToPair("[[1,9],[8,5]]"))
	assert.Equal(t, testPair, StrToPair("[1,[8,5]]"))
	assert.Equal(t, testPair, StrToPair("[1,5]"))
	assert.Equal(t, testPair, StrToPair("[[[[1,3],[5,3]],[[1,3],[8,7]]],[[[4,9],[6,9]],[[8,2],[7,3]]]]"))
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
