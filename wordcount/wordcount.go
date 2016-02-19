// Implement WordCount. It should return a map of the counts of each “word” in the string s. The wc.Test function runs a test suite against the provided function and prints success or failure.

//You might find strings.Fields helpful.

package main

import (
	"fmt"
	"strings"
	"unicode"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	// First create a function that selects the rune to be used in splitting the string into words. This will enable us to use the strings.FieldsFunc method for use in more complex strings
	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c) && !unicode.IsPunct(c)
	}

	words := strings.FieldsFunc(s, f)
	wordcount := make(map[string]int)
	for _, v := range words {
		if _, ok := wordcount[v]; ok {
			wordcount[v] += 1
		} else {
			wordcount[v] = 1
		}
	}
	return wordcount
}

func main() {
	wc.Test(WordCount)
	//WordCount("a man a plan a canal panama")
	wordcount := WordCount("Implement WordCount. It should return a map of the counts of each word in the string s. The wc.Test function runs a test suite against the provided function and prints success or failure. You might find strings.Fields helpful.")
	fmt.Printf("%v", wordcount)
}
