package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	m := make(map[string]int)

	for _, word := range words {
		_, ok := m[word]

		if ok {
			m[word]++
		} else {
			m[word] = 1
		}
	}

	return m
}

func main() {
	wc.Test(WordCount)
}
