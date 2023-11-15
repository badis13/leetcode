package main

import (
	"fmt"
)

// https://leetcode.com/problems/reverse-integer/
func main() {

	fmt.Println(reverse(123456))
}

const (
	min = -1 << 31
	max = 1<<31 - 1
)

func reverse(x int) int {
	d := 0
	isNegative := x < 0
	if isNegative {
		x = -x
	}
	for ; x > 0; x /= 10 {
		b := x % 10
		d = (d * 10) + b
	}
	if isNegative {
		d = -d
	}
	if d < min || d > max {
		return 0
	}
	return d
}
