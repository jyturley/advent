package main

import (
	"advent/pkg"
	"advent/pkg/execute"
	"fmt"
	"strconv"
	"strings"
)

func run(input string) (part1 interface{}, part2 interface{}) {
	// part 1
	binNums := strings.Split(input, "\n")
	mostCommon, leastCommon := getMostLeastCommonDigits(binNums)
	gamma := binToInt(mostCommon)
	epsilon := binToInt(leastCommon)
	part1 = gamma * epsilon

	// part 2
	o2rating := getO2Rating(binNums)
	co2scrubber := getCO2Rating(binNums)
	part2 = o2rating * co2scrubber
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
	return most, least
}

func getO2Rating(binNums []string) int {
	return getRating(binNums, true)
}
func getCO2Rating(binNums []string) int {
	return getRating(binNums, false)
}

func getRating(binNums []string, o2 bool) int {
	digitsPerNumber := len(binNums[0])
	filteredNums := make([]string, len(binNums))
	copy(filteredNums, binNums)

	digitIdx := 0
	for len(filteredNums) > 1 && digitIdx < digitsPerNumber {
		zeros := make([]string, 0)
		ones := make([]string, 0)
		for _, num := range filteredNums {
			currDigit := string(num[digitIdx])
			if currDigit == "0" {
				zeros = append(zeros, num)
			} else {
				ones = append(ones, num)
			}
		}
		if o2 {
			if len(zeros) <= len(ones) {
				filteredNums = ones
			} else {
				filteredNums = zeros
			}
		} else {
			if len(zeros) <= len(ones) {
				filteredNums = zeros
			} else {
				filteredNums = ones
			}
		}

		digitIdx++
	}
	out, err := strconv.ParseInt(filteredNums[0], 2, 64)
	if err != nil {
		fmt.Println(filteredNums[0], err)
	}
	return int(out)
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
	execute.Run(run, tests, puzzle, true)
}
