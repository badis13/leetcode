package main

import "fmt"

// https://leetcode.com/problems/jewels-and-stones/

func main() {
	a := "stowbj"
	b := "WJvb"
	fmt.Println(numJewelsInStones(a, b))
}

func numJewelsInStones(jewels string, stones string) int {
	var result int
	jewelsSet := make(map[byte]bool)
	for _, value := range jewels {
		jewelsSet[byte(value)] = true
	}
	for i := 0; i < len(stones); i++ {
		if jewelsSet[stones[i]] {
			result++
		}
	}

	return result
}
