package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAnalyze(t *testing.T) {
	entries := parse(minipizzle)
	assert.Equal(t, 5353, Analyze(entries[0]))
}
