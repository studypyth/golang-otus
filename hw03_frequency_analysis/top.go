package hw03

import (
	"fmt"
	"strings"
)

func Top10(text string) []string {
	frequency := map[string]int{}
	words := strings.Fields(text)
	for _, word := range words {
		fmt.Println(word)
		frequency[word]++
	}
	sortedKV := rankByWordCount(frequency)
	result := []string{}
	for i, word := range sortedKV {
		if i > 9 {
			break
		}
		result = append(result, word.Key)
	}
	return result
}
