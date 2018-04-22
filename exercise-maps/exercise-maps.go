package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	word_map := make(map[string]int)
	words := strings.Fields(s)
	for _, count := range words {
		word_map[count]++	
	}
	return word_map
}

func main() {
	wc.Test(WordCount)
}