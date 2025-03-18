package main

import (
	"fmt"
)

//https://leetcode.com/problems/excel-sheet-column-number/description/

func main() {

	fmt.Println(titleToNumber("ZY"))
}

func titleToNumber(columnTitle string) int {

	if len(columnTitle) > 7 {
		return -1
	}

	return powSumLen(columnTitle)

}

func powSumLen(str string) int {
	alph := map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
		"D": 4,
		"E": 5,
		"F": 6,
		"G": 7,
		"H": 8,
		"I": 9,
		"J": 10,
		"K": 11,
		"L": 12,
		"M": 13,
		"N": 14,
		"O": 15,
		"P": 16,
		"Q": 17,
		"R": 18,
		"S": 19,
		"T": 20,
		"U": 21,
		"V": 22,
		"W": 23,
		"X": 24,
		"Y": 25,
		"Z": 26,
	}

	if len(str) == 1 {
		return alph[string(str[0])]
	}

	result := alph[string(str[0])]

	for i := 0; i < len(str)-1; i++ {
		result += alph[string(str[i])]
	}

	result += alph[string(str[len(str)-1])]

	return result
}
