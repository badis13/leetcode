package main

import "fmt"

// https://leetcode.com/problems/roman-to-integer/

func main() {

	b := "XL"
	fmt.Println(romanToInt(b))
}

const (
	i = 'I'
	v = 'V'
	x = 'X'
	l = 'L'
	c = 'C'
	d = 'D'
	m = 'M'
)

func romanToInt(s string) (result int) {
	for j := len(s) - 1; j >= 0; j-- {
		if s[j] == i {
			if result > 2 {
				result--
				continue
			}
			result++
			continue
		}

		if s[j] == v {
			result += 5
			continue
		}

		if s[j] == x {
			if result >= 50 {
				result -= 10
				continue
			}
			result += 10
			continue
		}

		if s[j] == l {
			result += 50
			continue
		}

		if s[j] == c {
			if result >= 500 {
				result -= 100
				continue
			}
			result += 100
			continue
		}

		if s[j] == d {
			result += 500
			continue
		}

		if s[j] == m {
			result += 1000
			continue
		}
	}

	return result
}
