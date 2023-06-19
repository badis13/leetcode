package main

import (
	"fmt"
)

// https://leetcode.com/problems/merge-strings-alternately/
func main() {

	fmt.Println(mergeAlternately("hello", "world"))
}

func mergeAlternately(word1 string, word2 string) string {
	strMax := ""
	strMin := ""
	ac := make([]byte, 0)
	if len(word1) >= len(word2) {
		strMax = word1
		strMin = word2
	} else {
		strMax = word2
		strMin = word1
	}
	i := 0
	for i < len(strMin) {
		ac = append(ac, word1[i])
		ac = append(ac, word2[i])
		i++
	}
	for i < len(strMax) {
		ac = append(ac, strMax[i])
		i++
	}

	return string(ac)
}
