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
		{"[(])", scores["]"]},
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

func TestRepair(t *testing.T) {
	tests := []struct {
		in   string
		want string
	}{
		{"[({(<(())[]>[[{[]{<()<>>", "}}]])})]"},
		{"[(()[<>])]({[<{<<[]>>(", ")}>]})"},
		{"(((({<>}<{<{<>}{[]{[]{}", "}}>}>))))"},
		{"{<[[]]>}<{[{[{[]{()[[[]", "]]}}]}]}>"},
		{"<{([{{}}[<[[[<>{}]]]>[]]", "])}>"},
	}

	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			assert.Equal(t, tt.want, repair(tt.in), tt.in)
		})
	}
}

func TestGetScore(t *testing.T) {
	tests := []struct {
		in   string
		want int
	}{
		{"}}]])})]", 288957},
		{")}>]})", 5566},
		{"}}>}>))))", 1480781},
		{"]]}}]}]}>", 995444},
		{"])}>", 294},
	}

	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			assert.Equal(t, tt.want, getScore(tt.in), tt.in)
		})
	}
}
