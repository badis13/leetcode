package main

import (
	"container/heap"
	"fmt"
)

// https://leetcode.com/problems/relative-ranks/

func main() {

	b := []int{0, 1, 2, 4, 5, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17}
	fmt.Println(findRelativeRanks(b))
}

type ScoreRanks struct {
	score    int
	rank     string
	priority int
}

type Heap []*int

func findRelativeRanks(score []int) []string {
	var result []string

	scorePriority := make(map[int]int, len(score))
	for i := 0; i < len(score); i++ {
		scorePriority[i] = score[i]
	}

	ranksHeap := make(Heap, 0, len(score))

	for value, priority := range scorePriority {
		var curRank string
		if priority == 1 {
			curRank = "Gold Medal"
		}
		heap.Push(&ranksHeap, ScoreRanks{
			score:    value,
			priority: priority,
			rank:     curRank,
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
