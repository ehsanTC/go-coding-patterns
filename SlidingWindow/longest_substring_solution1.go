package main

import (
	"fmt"
)

func main() {
	var m string
	var k int

	fmt.Println("Enter the string: ")
	fmt.Scanf("%s", &m)

	fmt.Println("Enter the K: ")
	fmt.Scan(&k)

	var start, end, w int = 0, 0, 0
	var resultStart, resultEnd, longest int

	for end <= len(m) {
		substring := m[start:end]
		w = countUnique(substring)

		if w <= k {
			longest = len(substring)
			resultStart, resultEnd = start, end
			end++
		} else {
			start++
		}

		if w > k && end == len(m) {
			break
		}
	}

	fmt.Printf("The longest string is %s with len=%d", m[resultStart:resultEnd], longest)
}

func countUnique(str string) int {
	m := make(map[rune]bool)
	for _, r := range str {
		m[r] = true
	}
	return len(m)
}
