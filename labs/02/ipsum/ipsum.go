package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

var wordCountFlag, sentenceLengthFlag int

func init() {
	flag.IntVar(&wordCountFlag, "words", 100, "number of words to generate")
	flag.IntVar(&sentenceLengthFlag, "sentence-length", 6, "the length sentences should be")
}

var wordbank = []string{
	"foo",
	"bar",
	"baz",
}

func main() {
	flag.Parse()
	rand.Seed(time.Now().UnixNano())
	fmt.Println(generateIpsum(wordCountFlag, sentenceLengthFlag))
}

func generateIpsum(wordCount, sentenceLength int) string {
	ipsum := ""
	for wordsLeft := wordCount; wordsLeft > 0; wordsLeft -= sentenceLength {
		numWords := sentenceLength
		if wordsLeft < numWords {
			numWords = wordsLeft
		}
		if ipsum != "" {
			ipsum += "  "
		}
		ipsum += generateSentence(numWords)
	}
	return ipsum
}

func generateSentence(wordCount int) string {
	ipsum := ""
	for i := 0; i < wordCount; i++ {
		if ipsum != "" {
			ipsum += " "
		}
		ipsum += getWord()
	}
	return ipsum + "."
}

func getWord() string {
	index := rand.Intn(len(wordbank))
	return wordbank[index]
}
