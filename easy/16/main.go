package main

import "fmt"

// https://leetcode.com/problems/determine-the-winner-of-a-bowling-game/

func main() {
	a := []int{1, 5, 2, 8, 2}
	b := []int{1, 8, 9, 9, 0}
	fmt.Println(isWinner(a, b))
}

func isWinner(player1 []int, player2 []int) int {
	sum1 := findSum(player1)
	sum2 := findSum(player2)
	if sum1 > sum2 {
		return 1
	}
	if sum1 < sum2 {
		return 2
	}
	return 0
}

func findSum(x []int) int {
	var sum int
	var doublePoint int
	for i := 0; i < len(x); i++ {
		if doublePoint > 0 {
			sum += x[i]
			doublePoint--
		}
		if x[i] == 10 {
			doublePoint = 2
		}
		sum += x[i]
	}
	return sum
}
