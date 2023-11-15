package main

import "fmt"

// https://leetcode.com/problems/palindrome-number/

func main() {

	b := 252
	fmt.Println(isPalindrome(b))
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}
	var value2 int
	for value1 := x; value1 > 0; value1 /= 10 {
		value2 = value2*10 + value1%10
	}
	return value2 == x

}
