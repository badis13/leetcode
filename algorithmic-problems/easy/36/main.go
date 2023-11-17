package main

import (
	"container/heap"
	"fmt"
)

//https://leetcode.com/problems/the-k-weakest-rows-in-a-matrix/

func main() {
	a := [][]int{{1, 1, 1, 1, 1, 1}, {1, 1, 1, 0, 0, 0}, {1, 1, 0, 0, 0, 0}}
	k := 2
	fmt.Println(kWeakestRows(a, k))
}

type WinRanks struct {
	soldiers int
	priority int
}

type Heap []*WinRanks

func kWeakestRows(mat [][]int, k int) []int {
	result := make([]int, 0, len(mat))
	winHeap := make(Heap, 0, len(mat))
	for i := 0; i < len(mat); i++ {
		middle := len(mat[i]) / 2
		for middle < len(mat[i]) {
			if mat[i][middle] != 0 && middle+1 < len(mat[i]) && mat[i][middle+1] != 1 {
				heap.Push(&winHeap, WinRanks{
					soldiers: middle + 1,
					priority: i,
				})
				break
			}

			if middle-1 >= 0 && mat[i][middle-1] != 0 && mat[i][middle] != 1 {
				heap.Push(&winHeap, WinRanks{
					soldiers: middle + 1,
					priority: i,
				})
				break
			}

			if middle < 1 {
				heap.Push(&winHeap, WinRanks{
					soldiers: middle,
					priority: i,
				})
				break
			}

			if mat[i][middle] != 1 {
				middle--
				middle /= 2
				continue
			}
			if middle+1 < len(mat[i]) && mat[i][middle+1] != 0 {
				middle = middle/2 + middle
				continue
			}

			if middle == len(mat[i])-1 {
				if mat[i][len(mat[i])-1] != 0 {
					heap.Push(&winHeap, WinRanks{
						soldiers: middle + 1,
						priority: i,
					})
					break
				}
				heap.Push(&winHeap, WinRanks{
					soldiers: middle + 1,
					priority: i,
				})
				break
			}

			if middle == 0 {
				if mat[i][middle] != 0 {
					heap.Push(&winHeap, WinRanks{
						soldiers: middle + 1,
						priority: i,
					})
					break

				}
				heap.Push(&winHeap, WinRanks{
					soldiers: middle,
					priority: i,
				})
				break
			}
			middle++

		}
	}

	for i := 0; i < k; i++ {
		item := heap.Pop(&winHeap).(*WinRanks)
		result = append(result, item.priority)
	}

	return result

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
