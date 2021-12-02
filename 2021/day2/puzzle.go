package main

import (
	"advent/pkg/execute"
)

var tests = execute.TestCases{
	{
		testpuzzle,
		`150`,
		``,
	},
}

var testpuzzle = `forward 5
down 5
forward 8
up 3
down 8
forward 2`

var puzzle = ``
