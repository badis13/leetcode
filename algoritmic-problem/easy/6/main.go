package main

import "fmt"

// https://leetcode.com/problems/number-of-1-bits/

func main() {
	var b uint32 = 00000000000000000000000000001011
	fmt.Println(hammingWeight(b))
}

func hammingWeight(num uint32) int {
	result := 0
	for num > 0 {
		result += int(num % 2)
		num = num >> 1

	}
	return result

}
