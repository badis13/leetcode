package main

import (
	"fmt"
)

// heap solution https://leetcode.com/problems/top-k-frequent-words
func main() {
	words := []string{"plpaboutit", "jnoqzdute", "sfvkdqf", "mjc", "nkpllqzjzp", "foqqenbey", "ssnanizsav", "nkpllqzjzp", "sfvkdqf", "isnjmy", "pnqsz", "hhqpvvt", "fvvdtpnzx", "jkqonvenhx", "cyxwlef", "hhqpvvt", "fvvdtpnzx", "plpaboutit", "sfvkdqf", "mjc", "fvvdtpnzx", "bwumsj", "foqqenbey", "isnjmy", "nkpllqzjzp", "hhqpvvt", "foqqenbey", "fvvdtpnzx", "bwumsj", "hhqpvvt", "fvvdtpnzx", "jkqonvenhx", "jnoqzdute", "foqqenbey", "jnoqzdute", "foqqenbey", "hhqpvvt", "ssnanizsav", "mjc", "foqqenbey", "bwumsj", "ssnanizsav", "fvvdtpnzx", "nkpllqzjzp", "jkqonvenhx", "hhqpvvt", "mjc", "isnjmy", "bwumsj", "pnqsz", "hhqpvvt", "nkpllqzjzp", "jnoqzdute", "pnqsz", "nkpllqzjzp", "jnoqzdute", "foqqenbey", "nkpllqzjzp", "hhqpvvt", "fvvdtpnzx", "plpaboutit", "jnoqzdute", "sfvkdqf", "fvvdtpnzx", "jkqonvenhx", "jnoqzdute", "nkpllqzjzp", "jnoqzdute", "fvvdtpnzx", "jkqonvenhx", "hhqpvvt", "isnjmy", "jkqonvenhx", "ssnanizsav", "jnoqzdute", "jkqonvenhx", "fvvdtpnzx", "hhqpvvt", "bwumsj", "nkpllqzjzp", "bwumsj", "jkqonvenhx", "jnoqzdute", "pnqsz", "foqqenbey", "sfvkdqf", "sfvkdqf"}
	k := 1
	fmt.Println(topKFrequent(words, k))
}

type WordFrequent struct {
	word     string
	frequent int
}

func topKFrequent(words []string, k int) []string {
	frequencies := make(map[string]int)
	for i := 0; i < len(words); i++ {
		frequencies[words[i]]++
	}

	current := make([]WordFrequent, 0, len(frequencies))
	for word, frequency := range frequencies {
		current = append(current, WordFrequent{word: word, frequent: frequency})
	}

	current = newHeap(current)

	result := make([]string, k)
	for i := 0; i < k; i++ {
		result[i] = current[i].word
	}

	return result
}

func newHeap(arr []WordFrequent) []WordFrequent {
	var leftChild, rightChild, i, j int

	for largestChild := 0; (i < len(arr)) && (j < len(arr)); {
		leftChild = 2*i + 1
		rightChild = 2*i + 2
		largestChild = i

		if (leftChild < len(arr)) && (arr[leftChild].frequent > arr[largestChild].frequent) {
			temp := arr[largestChild]
			arr[largestChild] = arr[leftChild]
			arr[leftChild] = temp
			continue
		}

		if (rightChild < len(arr)) && (arr[rightChild].frequent > arr[largestChild].frequent) {
			temp := arr[largestChild]
			arr[largestChild] = arr[rightChild]
			arr[rightChild] = temp
			continue

		}

		if largestChild == i {
			i++
			continue
		}

	}

	return arr

}
