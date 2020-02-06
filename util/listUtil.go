package util

import (
	"regexp"
	"sort"
	"strings"
)

// ParseUtil ...
type ParseUtil interface {
}

// ListUtil ...
type ListUtil struct {
	Text   string
	Unique bool
}

// Words ...
func (l *ListUtil) Words() []string {
	r := regexp.MustCompile(`[^\w\s]`)
	text := r.ReplaceAllString(l.Text, "")
	text = strings.ToLower(text)
	wordSlice := strings.Split(text, " ")

	if l.Unique {
		wordSlice = unique(wordSlice)
	}

	sort.Slice(wordSlice, func(i, j int) bool {
		return wordSlice[i] < wordSlice[j]
	})

	return wordSlice
}

// Chars ...
func (l *ListUtil) Chars() []string {
	r := regexp.MustCompile(`[^\w]`)
	text := r.ReplaceAllString(l.Text, "")
	text = strings.ToLower(text)
	charSlice := strings.Split(text, "")

	if l.Unique {
		charSlice = unique(charSlice)
	}

	return charSlice
}

// Lines ...
func (l *ListUtil) Lines() []string {
	lineSlice := strings.Split(l.Text, "\n")

	if l.Unique {
		lineSlice = unique(lineSlice)
	}

	return lineSlice
}

func unique(slice []string) []string {
	uniqueMap := make(map[string]int)
	uniqueValues := make([]string, 0)

	for _, v := range slice {
		if _, ok := uniqueMap[v]; !ok {
			uniqueMap[v] = 1
			uniqueValues = append(uniqueValues, v)
		}
	}

	return uniqueValues
}
