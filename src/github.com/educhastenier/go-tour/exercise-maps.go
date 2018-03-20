package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	result := map[string]int{}
	for _, w := range words {
		result[w] = result[w] + 1
	}
	return result
}

func main() {
	wc.Test(WordCount)
}
