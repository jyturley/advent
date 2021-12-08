package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAnalyze(t *testing.T) {
	entries := parse(minipuzzle)
	numToPattern := map[int]string{}
	assert.Equal(t, 5353, Analyze(entries[0], numToPattern))
	assert.Equal(t, "cagedb", numToPattern[0], "test for 0")
	assert.Equal(t, "ab", numToPattern[1], "test for 1")
	assert.Equal(t, "gcdfa", numToPattern[2], "test for 2")
	assert.Equal(t, "fbcad", numToPattern[3], "test for 3")
	assert.Equal(t, "eafb", numToPattern[4], "test for 4")
	assert.Equal(t, "cdfbe", numToPattern[5], "test for 5")
	assert.Equal(t, "cdfgeb", numToPattern[6], "test for 6")
	assert.Equal(t, "dab", numToPattern[7], "test for 7")
	assert.Equal(t, "cefabd", numToPattern[9], "test for 9")
	assert.Equal(t, "acedgfb", numToPattern[8], "test for 8")

	entries = parse(testpuzzle)
	assert.Equalf(t, 8394, Analyze(entries[0], numToPattern), "%v | %v", entries[0].SignalPatterns, entries[0].OutputValues)
	assert.Equalf(t, 9781, Analyze(entries[1], numToPattern), "%v | %v", entries[1].SignalPatterns, entries[1].OutputValues)
	assert.Equalf(t, 1197, Analyze(entries[2], numToPattern), "%v | %v", entries[2].SignalPatterns, entries[2].OutputValues)
	assert.Equalf(t, 9361, Analyze(entries[3], numToPattern), "%v | %v", entries[3].SignalPatterns, entries[3].OutputValues)
	assert.Equalf(t, 4873, Analyze(entries[4], numToPattern), "%v | %v", entries[4].SignalPatterns, entries[4].OutputValues)
	assert.Equalf(t, 8418, Analyze(entries[5], numToPattern), "%v | %v", entries[5].SignalPatterns, entries[5].OutputValues)
	assert.Equalf(t, 4548, Analyze(entries[6], numToPattern), "%v | %v", entries[6].SignalPatterns, entries[6].OutputValues)
	assert.Equalf(t, 1625, Analyze(entries[7], numToPattern), "%v | %v", entries[7].SignalPatterns, entries[7].OutputValues)
	assert.Equalf(t, 8717, Analyze(entries[8], numToPattern), "%v | %v", entries[8].SignalPatterns, entries[8].OutputValues)
	assert.Equalf(t, 4315, Analyze(entries[9], numToPattern), "%v | %v", entries[9].SignalPatterns, entries[9].OutputValues)
}

func TestContains(t *testing.T) {
	assert.True(t, contains("abcde", "abc"))
	assert.True(t, contains("abcde", "bac"))
	assert.True(t, contains("bcdea", "bac"))
	assert.False(t, contains("abde", "abc"))
}

func TestDiff(t *testing.T) {
	assert.Equal(t, "de", diff("bcdea", "bac"))

}
