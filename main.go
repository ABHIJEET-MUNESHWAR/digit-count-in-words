package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>
func say(phrase string) {
	for _, word := range strings.Fields(phrase) {
		fmt.Printf("Abhijeet says %s...\n", word)
		dur := time.Duration(rand.Intn(100)) * time.Millisecond
		time.Sleep(dur)
	}
}

func main() {
	say("Golang is awesome")
}