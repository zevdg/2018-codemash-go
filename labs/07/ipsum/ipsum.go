package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/zevdg/2018-codemash-go/labs/07/wordbank"
)

var wordCountFlag, sentenceLengthFlag int
var wordbankPath string

func init() {
	flag.IntVar(&wordCountFlag, "words", 100, "number of words to generate")
	flag.IntVar(&sentenceLengthFlag, "sentence-length", 6, "the length sentences should be")
	flag.StringVar(&wordbankPath, "wordbank", "words.txt", "path to the text file containing the wordbank")
}

func main() {
	flag.Parse()
	rand.Seed(time.Now().UnixNano())
	ipsum, err := generateIpsum(wordCountFlag, sentenceLengthFlag)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(ipsum)
}

func generateIpsum(wordCount, sentenceLength int) (string, error) {
	f, err := os.Open(wordbankPath)
	if err != nil {
		return "", err
	}
	defer f.Close()
	wb, err := wordbank.New(f)
	if err != nil {
		return "", fmt.Errorf("couldn't create the wordbank: %s", err)
	}

	c := make(chan string)
	sentenceCount := 0
	for wordsLeft := wordCount; wordsLeft > 0; wordsLeft -= sentenceLength {
		sentenceCount++
		numWords := sentenceLength
		if wordsLeft < numWords {
			numWords = wordsLeft
		}
		go func() {
			c <- generateSentence(wb, numWords)
		}()
	}
	ipsum := ""
	for i := 0; i < sentenceCount; i++ {
		if ipsum != "" {
			ipsum += " "
		}
		ipsum += <-c
	}
	return ipsum, nil
}

func generateSentence(wb *wordbank.WordBank, wordCount int) string {
	ipsum := ""
	for i := 0; i < wordCount; i++ {
		if ipsum != "" {
			ipsum += " "
		}
		ipsum += wb.GetWord()
	}
	return ipsum + "."
}
