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

type pair struct {
	word  string
	count int
}

func countDigitsInWords(next func() string) counter {
	pending := make(chan string)
	counted := make(chan pair)
	// sends words to be counted
	go func() {
		for {
			word := next()
			pending <- word
			if word == "" {
				break
			}
		}
	}()
	// counts digits in words
	go func() {
		for {
			word := <-pending
			count := countDigits(word)
			counted <- pair{word, count}
			if word == "" {
				break
			}
		}
	}()
	stats := counter{}
	for {
		p := <-counted
		if p.word == "" {
			break
		}
		stats[p.word] = p.count
	}
	return stats
}
func wordGenerator(phrase string) func() string {
	words := strings.Fields(phrase)
	idx := 0
	return func() string {
		if idx >= len(words) {
			return ""
		}
		word := words[idx]
		idx++
		return word
	}
}
func main() {
	phrase := "on1 two22 th333"
	next := wordGenerator(phrase)
	fmt.Println(countDigitsInWords(next))
}