// Task:  Word Frequency Count
// Write a Go function that takes a string as input and returns a dictionary containing the frequency of each word in the string. Treat words in a case-insensitive manner and ignore punctuation marks.
// [Optional]: Write test for your function

// Task : Palindrome Check
// Write a Go function that takes a string as input and checks whether it is a palindrome or not. A palindrome is a word, phrase, number, or other sequence of characters that reads the same forward and backward (ignoring spaces, punctuation, and capitalization).
// [Optional]: Write test for your function
package main

import (
	"fmt"
	"strings"
	"unicode"
)

// Function to count word frequencies
func wordFrequencyCount(input string) map[string]int {
	normalized := strings.ToLower(input)
	normalized = strings.Map(func(r rune) rune {
		if unicode.IsLetter(r) || unicode.IsSpace(r) {
			return r
		}
		return -1 // Remove punctuation
	}, normalized)

	// Split the string into words
	words := strings.Fields(normalized)

	// Count frequencies
	frequency := make(map[string]int)
	for _, word := range words {
		frequency[word]++
	}

	return frequency
}

func isPalindrome(input string) bool {
	normalized := strings.ToLower(input)
	normalized = strings.Map(func(c rune) rune {
		if unicode.IsLetter(c) {
			return c
		}
		return -1
	}, normalized)
	i, j := 0, len(normalized)-1
	for i < j {
		if normalized[i] != normalized[j] {
			return false
		}
	}
	return true
}
func main2() {
	// Example usage
	text := "Hello, world! Hello Go. Go is fun."
	frequencies := wordFrequencyCount(text)
	fmt.Println("Word Frequencies:", frequencies)
	fmt.Println("Is text Palindrom: ", isPalindrome(text))
}
