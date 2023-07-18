package main

import (
	"fmt"
	"sort"
)

// https://leetcode.com/problems/top-k-frequent-words
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
	current := make([]WordFrequent, len(words))
	result := make([]string, k)
	currentIndex := 0
	currentFrequent := 0
	for i := 0; i < len(words); i++ {
		currentIndex, currentFrequent = findIndex(words[i], current)
		current[currentIndex].frequent = currentFrequent
		current[currentIndex].word = words[i]
	}

	sort.Slice(current, func(i, j int) bool {
		return (current[i].frequent > current[j].frequent) || (current[i].frequent == current[j].frequent && current[i].word < current[j].word)
	})

	for i := 0; i < k; i++ {
		result[i] = current[i].word
	}

	return result
}

func findIndex(word string, current []WordFrequent) (int, int) {
	for i := 0; i < len(current); i++ {
		if current[i].word == word {
			return i, current[i].frequent + 1
		}
		if current[i].word == "" {
			return i, 1
		}
	}
	return len(current), 1
}
