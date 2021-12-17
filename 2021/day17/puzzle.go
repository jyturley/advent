package main

import (
	"advent/pkg/execute"
)

var tests = execute.TestCases{
	{
		testpuzzle,
		`45`,
		`112`,
	},
}

var testpuzzle = `target area: x=20..30, y=-10..-5`
var puzzle = `target area: x=179..201, y=-109..-63`
