package main

import "fmt"

func main() {
	a := [][]int{{1, 1, 0, 0, 0}, {1, 1, 1, 1, 0}, {1, 0, 0, 0, 0}, {1, 1, 0, 0, 0}, {1, 1, 1, 1, 1}}
	k := 3
	fmt.Println(kWeakestRows(a, k))
}

func kWeakestRows(mat [][]int, k int) []int {

}
