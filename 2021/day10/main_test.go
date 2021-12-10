package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCheckLine(t *testing.T) {
	tests := []struct {
		in   string
		want int
	}{
		{"[({(<(())[]>[[{[]{<()<>>", 0},
		{"[>", scores[">"]},
		{"{()()()>", scores[">"]},
		{"(((()))}", scores["}"]},
		{"<([]){()}[{}])", scores[")"]},
		{"{([(<{}[<>[]}>{[]{[(<()>", 1197}, // Expected ], but found } instead.
		{"[[<[([]))<([[{}[[()]]]", 3},      // Expected ], but found ) instead.
		{"[{[{({}]{}}([{[{{{}}([]", 57},    // Expected ), but found ] instead.
		{"[<(<(<(<{}))><([]([]()", 3},      // Expected >, but found ) instead.
		{"<{([([[(<>()){}]>(<<{{", 25137},  //Expected ], but found > instead.
	}

	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			assert.Equal(t, tt.want, checkLine(tt.in), tt.in)
		})
	}
}
