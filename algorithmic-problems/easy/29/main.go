package main

import "fmt"

// https://leetcode.com/problems/sqrtx/

func main() {
	k := 9
	fmt.Println(mySqrt(k))
}

func mySqrt(x int) int {
	left := 0
	right := x
	findSqrt := 0
	for current := x; ; current = findSqrt {
		findSqrt = (left + right) / 2
		if findSqrt < 1 {
			return x
		}
		if findSqrt == current || findSqrt*findSqrt == x {
			break
		}
		if findSqrt*findSqrt < x {
			left = findSqrt
			continue
		}
		right = findSqrt

	}
	return findSqrt

}
