package main

import "fmt"

// https://leetcode.com/problems/remove-trailing-zeros-from-a-string/

func main() {
	nums := "1272105100"
	fmt.Println(removeTrailingZeros(nums))
}

func removeTrailingZeros(num string) string {
	var result string
	b := "0"
	for i := len(num) - 1; i >= 0; i-- {
		if num[i] != b[0] {
			result = num[0 : i+1]
			break
		}
		continue
	}
	return result
}
