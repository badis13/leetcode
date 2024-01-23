package main

import "fmt"

func main() {
	m := [][]int{{10, 1}, {10, 2}, {3, 10}, {10, 4}, {5, 10}, {10, 6}, {10, 7}, {8, 10}, {10, 9}, {10, 11}, {12, 10}, {10, 13}, {14, 10}}
	fmt.Println(findCenter(m))
}

func findCenter(edges [][]int) (result int) {

	if len(edges) == 0 {
		return -1
	}

	a := edges[0][0]
	b := edges[0][1]

	for i := 0; i < len(edges); i++ {
		if a == edges[i][0] || a == edges[i][1] {
			result = a
			continue
		}

		result = b

	}
	return result
}
