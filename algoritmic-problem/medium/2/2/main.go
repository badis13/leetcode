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
	frequencies := make(map[string]int)
	for i := 0; i < len(words); i++ {
		frequencies[words[i]]++
	}

	current := make([]WordFrequent, 0, len(frequencies))
	for word, frequency := range frequencies {
		current = append(current, WordFrequent{word: word, frequent: frequency})
	}

	sort.Slice(current, func(i, j int) bool {
		return (current[i].frequent > current[j].frequent) || (current[i].frequent == current[j].frequent && current[i].word < current[j].word)
	})

	result := make([]string, k)
	for i := 0; i < k; i++ {
		result[i] = current[i].word
	}

	return result
}
