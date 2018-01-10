package main

import (
	"flag"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/zevdg/2018-codemash-go/labs/09/wordbank"
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

	svc, err := NewService(wordbankPath)
	if err != nil {
		fmt.Println("Error initializing webservice:", err)
		return
	}
	http.HandleFunc("/ipsum", svc.ipsumHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error serving webservice:", err)
	}
}

type Service struct {
	wb *wordbank.WordBank
}

func NewService(filePath string) (*Service, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	wb, err := wordbank.New(f)
	if err != nil {
		return nil, fmt.Errorf("couldn't create the wordbank: %s", err)
	}
	return &Service{wb}, nil
}

func (svc *Service) ipsumHandler(w http.ResponseWriter, r *http.Request) {
	ipsum, err := svc.generateIpsum(wordCountFlag, sentenceLengthFlag)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("Error:", err)
		return
	}
	fmt.Fprint(w, ipsum)
}

func (svc *Service) generateIpsum(wordCount, sentenceLength int) (string, error) {
	c := make(chan string)
	sentenceCount := 0
	for wordsLeft := wordCount; wordsLeft > 0; wordsLeft -= sentenceLength {
		sentenceCount++
		numWords := sentenceLength
		if wordsLeft < numWords {
			numWords = wordsLeft
		}
		go func() {
			c <- svc.generateSentence(numWords)
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

func (svc *Service) generateSentence(wordCount int) string {
	ipsum := ""
	for i := 0; i < wordCount; i++ {
		if ipsum != "" {
			ipsum += " "
		}
		ipsum += svc.wb.GetWord()
	}
	return ipsum + "."
}
