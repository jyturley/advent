package pkg

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ParseIntList(s, sep string) []int {
	lines := strings.Split(s, sep)
	list := make([]int, len(lines))
	for i, line := range lines {
		nb, err := strconv.Atoi(line)
		Check(err)
		list[i] = nb
	}
	return list
}

func ParseIntGrid(s string) [][]int {
	out := [][]int{}
	lines := strings.Split(s, "\n")
	for _, line := range lines {
		row := make([]int, 0)
		s := strings.Split(line, "")
		for i := 0; i < len(line); i++ {
			if h, err := strconv.Atoi(s[i]); err == nil {
				row = append(row, h)
			} else {
				fmt.Println(err)
			}
		}
		out = append(out, row)
	}
	return out
}

func GetAoVDate() (year, day string) {
	dir, err := os.Getwd()
	Check(err)
	split := strings.Split(dir, "/")
	if len(split) == 9 && strings.HasPrefix(split[8], "day") {
		return split[7], split[8][3:]
	} else if len(split) == 8 && strings.HasPrefix(split[7], "day") {
		return split[6], split[7][3:]
	}

	return "2021", "Unknown"
}

func MustScanf(line, format string, a ...interface{}) {
	n, err := fmt.Sscanf(line, format, a...)
	Check(err)
	if n != len(a) {
		panic(fmt.Errorf("%d args expected in scanf, got %d", len(a), n))
	}
}

func Check(err error) {
	if err != nil {
		panic(err)
	}
}
