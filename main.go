package main

import (
	"fmt"
	"strings"
	"unicode"
)

func countDigits(str string) int {
	count := 0
	for _, char := range str {
		if unicode.IsDigit(char) {
			count++
		}
	}
	return count
}

type counter map[string]int

type pair struct {
	word  string
	count int
}

func countDigitsInWords(next func() string) counter {

	pending := make(chan string)
	go submitWords(next, pending)

	counted := make(chan pair)
	go countWords(pending, counted)

	return fillStats(counted)
}

// submitWords sends words to be counted.
func submitWords(next func() string, out chan string) {
	// sends words to be counted
	for {
		word := next()
		out <- word
		if word == "" {
			break
		}
	}
}

// countWords counts digits in words.
func countWords(in chan string, out chan pair) {
	for {
		word := <-in
		count := countDigits(word)
		out <- pair{word, count}
		if word == "" {
			break
		}
	}
}

// fillStats prepares the final statistics.
func fillStats(in chan pair) counter {
	stats := counter{}
	for {
		p := <-in
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