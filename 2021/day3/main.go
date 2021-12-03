package main

import (
	"advent/pkg"
	"advent/pkg/execute"
	"fmt"
	"strconv"
	"strings"
)

func run(input string) (part1 interface{}, part2 interface{}) {
	//part 1
	binNums := strings.Split(input, "\n")
	mostCommon, leastCommon := getMostLeastCommonDigits(binNums)
	gamma := binToInt(mostCommon)
	epsilon := binToInt(leastCommon)
	part1 = gamma * epsilon

	return part1, part2
}

func getMostLeastCommonDigits(binNums []string) (most, least []int) {
	digitsPerNumber := len(binNums[0])
	majority := len(binNums) / 2
	most = make([]int, digitsPerNumber)
	least = make([]int, digitsPerNumber)
	for j := 0; j < digitsPerNumber; j++ {
		zero, one := 0, 0
		for i := 0; i < len(binNums); i++ {
			if zero >= majority {
				most[j] = 0
				least[j] = 1
				break
			}
			if one >= majority {
				most[j] = 1
				least[j] = 0
				break
			}
			if binNums[i][j] == '0' {
				zero++
			} else {
				one++
			}
		}
	}
	fmt.Println(most)
	fmt.Println(least)
	return most, least
}

func binToInt(commonDigits []int) int {
	strNum := ""
	for _, d := range commonDigits {
		a := strconv.Itoa(d)
		strNum += a
	}

	out, err := strconv.ParseInt(strNum, 2, 64)
	if err != nil {
		fmt.Println(strNum, err)
	}

	return int(out)
}

func main() {
	year, day := pkg.GetAoVDate()
	fmt.Printf("Advent of Code %s Day %s\n", year, day)
	execute.Run(run, tests, testpuzzle, true)
}
