package main

import "fmt"

// https://leetcode.com/problems/plus-one/

func main() {
	k := []int{1, 2, 3, 4, 5, 9}
	fmt.Println(plusOne(k))
}

func plusOne(digits []int) []int {

	if digits[len(digits)-1] != 9 {
		digits[len(digits)-1]++
		return digits
	}

	return increment(reverse(digits))

}

func reverse(digits []int) []int {
	left, right := 0, len(digits)-1

	for left < right {
		current := digits[left]
		digits[left] = digits[right]
		digits[right] = current
		left++
		right--
	}
	return digits
}

func increment(digits []int) []int {
	for i, j := 0, 0; i < len(digits); i++ {
		if digits[i] != 9 {
			digits[j]++
			return reverse(digits)
		}
		digits[i] = 0
		j++
	}
	digits = append(digits, 1)
	return reverse(digits)
}
