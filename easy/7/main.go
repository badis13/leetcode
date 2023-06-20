package main

import "fmt"

// https://leetcode.com/problems/calculate-delayed-arrival-time/

func main() {
	a, b := 12, 13
	fmt.Println(findDelayedArrivalTime(a, b))
}

func findDelayedArrivalTime(arrivalTime int, delayedTime int) int {
	return (arrivalTime + delayedTime) % 24
}
