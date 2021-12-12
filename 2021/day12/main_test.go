package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStringify(t *testing.T) {
	// positions := parse(testpuzzle)
	assert.Equal(t, "start,A,b,A,c,A,end", stringify([]string{"start", "A", "b", "A", "c", "A", "end"}))
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
