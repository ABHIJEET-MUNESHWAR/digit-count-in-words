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
	pending := submitWords(next)
	counted := countWords(pending)
	return fillStats(counted)
}

// submitWords sends words to be counted.
func submitWords(next func() string) chan string {
	// sends words to be counted
	out := make(chan string)
	go func() {
		for {
			word := next()
			out <- word
			if word == "" {
				break
			}
		}
	}()
	return out
}

// countWords counts digits in words.
func countWords(in chan string) chan pair {
	out := make(chan pair)
	go func() {
		for {
			word := <-in
			count := countDigits(word)
			out <- pair{word, count}
			if word == "" {
				break
			}
		}
	}()
	return out
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