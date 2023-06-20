package main

import "fmt"

// https://leetcode.com/problems/add-digits/

func main() {
	a := 1236
	fmt.Println(addDigits(a))
}

func addDigits(num int) int {
	for num >= 10 {
		sum := 0
		for ; num > 0; num /= 10 {
			sum += num % 10
		}

		num = sum
	}

	return num
}
