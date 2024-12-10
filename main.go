package main

import (
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"time"
	"unicode"
)

// TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>
func say(id int, phrase string) {
	for _, word := range strings.Fields(phrase) {
		fmt.Printf("Worker #%d says %s...\n", id, word)
		dur := time.Duration(rand.Intn(100)) * time.Millisecond
		time.Sleep(dur)
	}
}

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
	var wg sync.WaitGroup
	syncStats := new(sync.Map)
	words := strings.Fields(phrase)
	wg.Add(len(words))
	for _, word := range words {
		go func() {
			defer wg.Done()
			count := countDigits(word)
			syncStats.Store(word, count)
		}()
	}
	wg.Wait()
	return nil
}
func main() {
	messages := make(chan string)
	go func() {
		messages <- "Ping"
	}()
	msg := <-messages
	fmt.Println(msg)
}