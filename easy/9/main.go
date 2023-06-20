package main

import "fmt"

// https://leetcode.com/problems/sum-multiples/

func main() {
	a := 10
	fmt.Println(sumOfMultiples(a))
}

func sumOfMultiples(n int) int {
	sum := 0
	for n > 0 {
		if n%3 == 0 || n%5 == 0 || n%7 == 0 {
			sum += n
		}
		n--
	}
	return sum
}
