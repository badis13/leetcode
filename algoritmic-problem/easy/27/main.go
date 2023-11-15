package main

import "fmt"

// https://leetcode.com/problems/find-the-index-of-the-first-occurrence-in-a-string/

func main() {
	a := "sadsoul"
	b := "sad"
	fmt.Println(strStr(a, b))
}

func strStr(haystack string, needle string) int {
	if len(haystack) < len(needle) {
		return -1
	}

	return result(haystack, needle, findKnuthMorrisPrattArr(needle))
}

func result(haystack string, needle string, arrKnuthMorrisPratt []int) int {
	indexNeedle := 0
	for indexHaystack := 0; indexHaystack < len(haystack); indexHaystack++ {
		if needle[indexNeedle] == haystack[indexHaystack] {
			indexNeedle++
			if indexNeedle == len(needle) {
				return indexHaystack - (len(needle) - 1)
			}
			continue
		}
		if indexNeedle != 0 {
			indexNeedle = arrKnuthMorrisPratt[indexNeedle-1]
			indexHaystack--
			continue
		}
	}
	return -1
}

func findKnuthMorrisPrattArr(needle string) []int {
	arrKnuthMorrisPratt := make([]int, len(needle))
	i := 0
	j := 1
	for j < len(needle) {

		if needle[i] != needle[j] {
			if i != 0 {
				i = arrKnuthMorrisPratt[i-1]
				continue
			}
			arrKnuthMorrisPratt[j] = 0
			j++
			continue
		}

		if needle[i] == needle[j] {
			arrKnuthMorrisPratt[j] = i + 1
			i++
			j++
			continue
		}
	}
	return arrKnuthMorrisPratt
}
