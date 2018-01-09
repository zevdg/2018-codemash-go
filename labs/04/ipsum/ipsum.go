package main

import (
	"bufio"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"
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
	fmt.Println(generateIpsum(wordCountFlag, sentenceLengthFlag))
}

func generateIpsum(wordCount, sentenceLength int) string {
	f, err := os.Open(wordbankPath)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	wb := NewWordBank(f)
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

func NewWordBank(f *os.File) *WordBank {
	wb := &WordBank{}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		wb.bank = append(wb.bank, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return wb
}

func (wb *WordBank) getWord() string {
	index := rand.Intn(len(wb.bank))
	return wb.bank[index]
}
