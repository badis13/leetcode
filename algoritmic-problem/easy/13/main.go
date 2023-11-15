package main

import "fmt"

// https://leetcode.com/problems/longest-common-prefix/

func main() {

	b := []string{"flur", "flug", "flag", "flig"}
	fmt.Println(longestCommonPrefix(b))
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	var lcpLength int
	for letterIndex := 0; ; letterIndex++ {
		for strIndex := 0; strIndex < len(strs)-1; strIndex++ {
			if letterIndex >= len(strs[strIndex]) || letterIndex >= len(strs[strIndex+1]) {
				return strs[0][:lcpLength]
			}
			if strs[strIndex][letterIndex] == strs[strIndex+1][letterIndex] {
				continue
			}
			return strs[0][:lcpLength]
		}
		lcpLength++
		continue
	}
}
