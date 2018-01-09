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
	wb, err := NewWordBank(f)
	if err != nil {
		return "", fmt.Errorf("couldn't create the wordbank: %s", err)
	}
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
	return ipsum, nil
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

func NewWordBank(f *os.File) (*WordBank, error) {
	wb := &WordBank{}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		wb.bank = append(wb.bank, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return wb, nil
}

func (wb *WordBank) getWord() string {
	index := rand.Intn(len(wb.bank))
	return wb.bank[index]
}
