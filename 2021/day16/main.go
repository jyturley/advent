package main

import (
	"advent/pkg"
	"advent/pkg/execute"
	"encoding/hex"
	"fmt"
	"strconv"
)

/*
0 = 0000
1 = 0001
2 = 0010
3 = 0011
4 = 0100
5 = 0101
6 = 0110
7 = 0111
8 = 1000
9 = 1001
A = 1010
B = 1011
C = 1100
D = 1101
E = 1110
F = 1111
*/

const (
	literalValueTypeID = 4
	aLengthTypeID      = 0
	bLengthTypeID      = 1
)

func run(input string) (part1 interface{}, part2 interface{}) {
	// part 1
	list := parse(input)
	part1 = Part1(list)
	return part1, part2
}

func Part1(binary string) int {
	fmt.Println(binary)
	out := 0
	version, typeID := binary[0:3], binary[3:6]
	fmt.Println(version, typeID)
	decVersion := binstrToInt(version)
	decTypeID := binstrToInt(typeID)
	fmt.Printf("version %s -> %d\n", version, decVersion)
	fmt.Printf("typeID %s -> %d\n", version, decTypeID)

	if decTypeID == literalValueTypeID {
		out = getLiteralValue(binary)
		fmt.Printf("literal value: %d\n", out)
		return out
	}

	ltid := string(binary[7])
	lengthTypeID := binstrToInt(ltid)
	fmt.Printf("lengthTypeID %d\n", lengthTypeID)
	if lengthTypeID == aLengthTypeID {
		totalLenInBits := binstrToInt(binary[7:22])
		fmt.Println(binary[7:22])
		fmt.Printf("total length in bits: %d\n", totalLenInBits)
		subpackets := binary[22 : 22+totalLenInBits]
		fmt.Println(subpackets)
		out = parseSubpackets(subpackets)
		fmt.Printf("total subpacket: %d\n", out)
	}

	return out
}

func parseSubpackets(subpacket string) int {
	out := 0
	version, typeID, rest := binstrToInt(binary[0:3]), binstrToInt(binary[3:6])
	if typeID == literalValueTypeID {
		lv := getLiteralValue(subpacket)
		fmt.Println("literal value returned", lv)
		out += lv
	}
	return out
}

func getLiteralValue(whole string) int {
	for len(whole)%4 != 0 {
		fmt.Println("not multiple of 4")
		whole = "0" + whole
	}
	rest := whole[6:]
	g1, g2, g3 := rest[0:5], rest[5:10], rest[10:15]
	literal := g1[1:] + g2[1:] + g3[1:]
	fmt.Println(literal)
	decLiteral := binstrToInt(literal)
	return int(decLiteral)
}

func Part2() int {
	out := 0
	return out
}

func parse(s string) string {
	out := ""
	decoded, err := hex.DecodeString(s)
	if err != nil {
		panic(err)
	}
	out = convertToBinaryStr(decoded)
	return out
}

func binstrToInt(s string) int {
	out, _ := strconv.ParseInt(s, 2, 64)
	return int(out)
}

func convertToBinaryStr(decoded []byte) string {
	out := ""
	fmt.Println(len(decoded))
	for _, b := range decoded {
		out = fmt.Sprintf("%s%08b", out, b)
	}
	return out
}

func main() {
	year, day := pkg.GetAoVDate()
	fmt.Printf("Advent of Code %s Day %s\n", year, day)
	execute.Run(run, tests, puzzle, true)
}
