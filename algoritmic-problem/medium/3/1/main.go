package main

import (
	"container/heap"
	"fmt"
)

// heap solution https://leetcode.com/problems/top-k-frequent-elements/
func main() {
	nums := []int{1, 1, 1, 2, 2, 3}
	k := 2
	fmt.Println(topKFrequent(nums, k))
}

type NumFrequent struct {
	num      int
	frequent int
}

type Heap []*NumFrequent

func topKFrequent(nums []int, k int) []int {
	var result []int
	frequencies := make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		frequencies[nums[i]]++
	}

	wfi := make(Heap, 0, len(frequencies))

	for value, priority := range frequencies {
		heap.Push(&wfi, NumFrequent{
			num:      value,
			frequent: priority,
		})

	}

	for i := 0; i < k; i++ {
		item := heap.Pop(&wfi).(*NumFrequent)
		result = append(result, item.num)
	}

	return result
}

func (h Heap) Less(i, j int) bool {
	if h[i].frequent == h[j].frequent {
		return false
	}

	return h[i].frequent > h[j].frequent
}

func (h Heap) Len() int {
	return len(h)
}

func (h Heap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *Heap) Push(x any) {
	item := x.(NumFrequent)
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
