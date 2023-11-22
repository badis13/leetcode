package main

import (
	"container/heap"
	"fmt"
)

//https://leetcode.com/problems/the-k-weakest-rows-in-a-matrix/

func main() {
	a := [][]int{{1, 0}, {0, 0}, {1, 0}}
	k := 2
	fmt.Println(kWeakestRows(a, k))
}

type WinRanks struct {
	soldiers int
	priority int
}

type Heap []*WinRanks

func kWeakestRows(mat [][]int, k int) []int {
	result := make([]int, k)
	winHeap := make(Heap, 0, len(mat))
	for i := 0; i < len(mat); i++ {
		heap.Push(&winHeap, WinRanks{
			soldiers: countSoldiers(mat[i]),
			priority: i,
		})
	}

	for i := 0; i < k; i++ {
		result[i] = heap.Pop(&winHeap).(*WinRanks).priority
	}

	return result

}

func countSoldiers(arr []int) int {
	left, right := 0, len(arr)-1

	if arr[left] == 0 {
		return 0
	}

	if arr[right] == 1 {
		return len(arr)
	}

	for middle := right / 2; ; middle = left + (right-left)/2 {
		switch {
		case arr[middle] != 0:
			left = middle + 1
		case arr[middle-1] != 1:
			right = middle - 1
		default:
			return middle

		}
	}
}

func (h Heap) Less(i, j int) bool {
	if h[i].soldiers == h[j].soldiers {
		return h[i].priority < h[j].priority
	}
	return h[i].soldiers < h[j].soldiers
}

func (h Heap) Len() int {
	return len(h)
}

func (h Heap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *Heap) Push(x any) {
	item := x.(WinRanks)
	*h = append(*h, &item)
}

func (h *Heap) Pop() any {
	old := *h
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	*h = old[0 : n-1]
	return item
}
