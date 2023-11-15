package main

import "fmt"

// https://leetcode.com/problems/length-of-last-word/

func main() {
	s := "   fly me   to   the moon  "
	fmt.Println(lengthOfLastWord(s))
}

func lengthOfLastWord(s string) int {
	right := len(s) - 1
	for ; right >= 0; right-- {
		if s[right] != ' ' {
			break
		}
	}
	left := len(s[:right])
	for ; left >= 0; left-- {
		if s[left] == ' ' {
			break
		}
	}
	return right - left
}
