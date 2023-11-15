package main

import (
	"container/heap"
	"fmt"
)

// heap solution https://leetcode.com/problems/top-k-frequent-words
func main() {
	words := []string{"i", "love", "leetcode", "i", "love", "coding"}
	k := 2
	fmt.Println(topKFrequent(words, k))
}

type WordFrequent struct {
	word     string
	frequent int
	index    int
}

type Heap []*WordFrequent

func topKFrequent(words []string, k int) []string {
	var result []string
	frequencies := make(map[string]int, len(words))
	for i := 0; i < len(words); i++ {
		frequencies[words[i]]++
	}

	wfi := make(Heap, 0, len(frequencies))

	i := 0
	for value, priority := range frequencies {
		heap.Push(&wfi, WordFrequent{
			word:     value,
			frequent: priority,
			index:    i,
		})
		i++
	}

	for i := 0; i < k; i++ {
		item := heap.Pop(&wfi).(*WordFrequent)
		result = append(result, item.word)
	}

	return result
}

func (h Heap) Less(i, j int) bool {
	if h[i].frequent == h[j].frequent {
		return h[i].word < h[j].word
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
	n := len(*h)
	item := x.(WordFrequent)
	item.index = n
	*h = append(*h, &item)
}

func (h *Heap) Pop() any {
	old := *h
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*h = old[0 : n-1]
	return item
}
