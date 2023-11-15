package main

import "fmt"

// https://leetcode.com/problems/add-binary/

func main() {
	a := "1010"
	b := "1111"
	fmt.Println(addBinary(a, b))
}

func addBinary(a string, b string) string {
	var result string
	var current byte = '0'
	a, b = addZeroString(a, b)

	for i := len(a) - 1; i >= 0; i-- {
		result = string((a[i]^b[i])^current) + result
		if a[i]+b[i] == 97 {
			continue
		}
		current = a[i] & b[i]
	}

	if current != '0' {
		result = "1" + result
	}
	return string(result)
}

func addZeroString(a, b string) (string, string) {

	for len(a) > len(b) {

		b = "0" + b
	}
	for len(b) > len(a) {
		a = "0" + a
	}

	return a, b

}
