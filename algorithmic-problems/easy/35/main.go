package main

import (
	"container/heap"
	"fmt"
	"strconv"
)

// https://leetcode.com/problems/relative-ranks/

func main() {

	b := []int{2, 4, 17, 7, 8, 9, 15, 11, 12, 13, 14, 10, 16, 5}
	fmt.Println(findRelativeRanks(b))
}

type ScoreRanks struct {
	score int
	index int
}

type Heap []*ScoreRanks

func findRelativeRanks(score []int) []string {
	result := make([]string, len(score))

	ranksHeap := make(Heap, 0, len(score))
	for index, value := range score {
		heap.Push(&ranksHeap, ScoreRanks{
			score: value,
			index: index,
		})

	}

	for i := 0; i < len(score); i++ {
		item := heap.Pop(&ranksHeap).(*ScoreRanks)

		switch i {
		case 0:
			result[item.index] = "Gold Medal"
		case 1:
			result[item.index] = "Silver Medal"
		case 2:
			result[item.index] = "Bronze Medal"
		default:
			result[item.index] = strconv.Itoa(i + 1)
		}

	}

	return result

}

func (h Heap) Less(i, j int) bool {
	return h[i].score > h[j].score
}

func (h Heap) Len() int {
	return len(h)
}

func (h Heap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *Heap) Push(x any) {
	item := x.(ScoreRanks)
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
