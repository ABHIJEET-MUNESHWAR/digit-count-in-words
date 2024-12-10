package main

import (
	"fmt"
	"strings"
	"unicode"
)

// countDigits counts the number of digit characters in a given string.
//
// Parameters:
//   - str: The input string to be analyzed.
//
// Returns:
//
//	An integer representing the count of digit characters found in the input string.
func countDigits(str string) int {
	count := 0
	for _, char := range str {
		if unicode.IsDigit(char) {
			count++
		}
	}
	return count
}

// counter stores the number of digits in each word.
// The key is the word, and the value is the number of digits.
type counter map[string]int

func countDigitsInWords(phrase string) counter {
	words := strings.Fields(phrase)
	counted := make(chan int)
	go func() {
		for _, word := range words {
			counted <- countDigits(word)
		}
	}()
	stats := counter{}
	for _, word := range words {
		count := <-counted
		stats[word] = count
	}
	return stats
}
func main() {
	fmt.Println(countDigitsInWords("on1 two22 th333"))
}