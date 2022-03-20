package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

type repeatingWord struct {
	Word        string
	RepeatCount int
}

func Top10(unsortedString string) []string {
	if len(unsortedString) == 0 {
		return make([]string, 0)
	}

	clearSlice := strings.Fields(unsortedString)
	if len(clearSlice) == 0 {
		return make([]string, 0)
	}

	mapString := countAmountWord(clearSlice)

	unsortedSlice := make([]repeatingWord, len(mapString))
	unsortedIndex := 0
	for word, repeatCount := range mapString {
		unsortedSlice[unsortedIndex] = repeatingWord{word, repeatCount}
		unsortedIndex++
	}

	sort.Slice(unsortedSlice, func(i, j int) bool {
		if unsortedSlice[i].RepeatCount == unsortedSlice[j].RepeatCount {
			return unsortedSlice[i].Word < unsortedSlice[j].Word
		}

		return unsortedSlice[i].RepeatCount > unsortedSlice[j].RepeatCount
	})

	currentLen := len(unsortedSlice)
	if currentLen > 10 {
		currentLen = 10
	}

	sortedSlice := make([]string, currentLen)
	for i, repeatingWord := range unsortedSlice[0:currentLen] {
		sortedSlice[i] = repeatingWord.Word
	}

	return sortedSlice
}

func countAmountWord(sliceStings []string) map[string]int {
	mapString := map[string]int{}
	for _, line := range sliceStings {
		mapString[line]++
	}

	return mapString
}
