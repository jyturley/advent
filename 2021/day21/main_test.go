package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBoard(t *testing.T) {
	fmt.Println("testing")
	b := NewBoard(1, 3)
	fmt.Println(b.data)
	assert.Equal(t, 1, b.P1())
	assert.Equal(t, 3, b.P2())

	b = NewBoard(7, 8)
	assert.Equal(t, 2, b.AdvanceP1(5), b.data)
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
