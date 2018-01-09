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

func main() {
	flag.Parse()
	rand.Seed(time.Now().UnixNano())
	fmt.Println(generateIpsum(wordCountFlag, sentenceLengthFlag))
}

func generateIpsum(wordCount, sentenceLength int) string {
	wb := NewWordBank()
	ipsum := ""
	for wordsLeft := wordCount; wordsLeft > 0; wordsLeft -= sentenceLength {
		numWords := sentenceLength
		if wordsLeft < numWords {
			numWords = wordsLeft
		}
		if ipsum != "" {
			ipsum += "  "
		}
		ipsum += generateSentence(wb, numWords)
	}
	return ipsum
}

func generateSentence(wb *WordBank, wordCount int) string {
	ipsum := ""
	for i := 0; i < wordCount; i++ {
		if ipsum != "" {
			ipsum += " "
		}
		ipsum += wb.getWord()
	}
	return ipsum + "."
}

type WordBank struct {
	bank []string
}

func NewWordBank() *WordBank {
	wb := &WordBank{
		bank: []string{
			"foo",
			"bar",
			"baz",
		},
	}
	return wb
}

func (wb *WordBank) getWord() string {
	index := rand.Intn(len(wb.bank))
	return wb.bank[index]
}
