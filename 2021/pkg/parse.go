package pkg

import (
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

func GetAoVDate() (year, day string) {
	dir, err := os.Getwd()
	Check(err)
	split := strings.Split(dir, "/")
	if len(split) == 9 {
		return split[7], split[8]
	}

	return "2021", "Unknown"
}

func Check(err error) {
	if err != nil {
		panic(err)
	}
}
