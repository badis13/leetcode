package main

import "fmt"

// https://leetcode.com/problems/lexicographically-smallest-palindrome/description/

func main() {

	b := "egcfe"
	fmt.Println(makeSmallestPalindrome(b))
}

func makeSmallestPalindrome(s string) string {

	cur := []byte(s)
	l := 0
	r := len(s) - 1

	for l < r {
		if cur[l] > cur[r] {
			cur[l] = cur[r]
			l++
			r--
			continue
		}
		cur[r] = cur[l]
		l++
		r--
	}

	return string(cur)
}
