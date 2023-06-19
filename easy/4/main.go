package main

import "fmt"

// https://leetcode.com/problems/reverse-string/
func main() {
	b := []byte("hello")
	fmt.Println(string(b))
	reverseString(b)
	fmt.Println(string(b))
}

func reverseString(s []byte) {
	l := 0
	r := len(s) - 1
	for l < r {
		s[l], s[r] = s[r], s[l]
		r--
		l++
	}

}
