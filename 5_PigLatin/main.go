package main

import (
	"fmt"
	"regexp"
	"strings"
)

func split(s string) []string {
	words := regexp.MustCompile(`\s+`).Split(s, -1)
	for i, word := range words {
		words[i] = string(strings.ToLower(word))
	}
	return words
}

func toPigLatin(s string) string {
	b := []byte(s)
	b[0], b[len(s)-1] = b[len(s)-1], b[0]
	b = append(b, 'a', 'y')
	return string(b)
}

func main() {
	phrase := `the cat is in the barn`
	words := split(phrase)
	for i, word := range words {
		words[i] = toPigLatin(word)
	}
	fmt.Println(strings.Join(words, " "))
}
