package main

import "fmt"

//https://leetcode.com/problems/sqrtx/

func main() {
	fmt.Println(mySqrt(8))
}

func mySqrt(x int) int {
	if x < 1 {
		return x
	}

	if x < 4 {
		return 1
	}
	left := 0
	right := x
	currentRoot := 0
	return recurseBinarySearch(left, right, currentRoot, x)
}

func recurseBinarySearch(l, r, cur, x int) int {
	result := (l + r) / 2
	if result*result == x || result == cur {
		return result
	}
	if result*result < x {
		return recurseBinarySearch(result, r, result, x)
	}
	if result*result > x {
		return recurseBinarySearch(l, result, result, x)
	}
	return result

}
