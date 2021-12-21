package main

import (
	"advent/pkg/execute"
)

var tests = execute.TestCases{
	{
		testpuzzle,
		`739785`,
		``,
	},
}

var testpuzzle = `Player 1 starting position: 4
Player 2 starting position: 8`
var puzzle = `Player 1 starting position: 8
Player 2 starting position: 9`
