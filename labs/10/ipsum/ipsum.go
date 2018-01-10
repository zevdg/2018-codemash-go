package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/zevdg/2018-codemash-go/labs/10/wordbank"
)

var wordbankPath string

func init() {
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
	req := IpsumReq{}
	reqDecoder := json.NewDecoder(r.Body)
	if err := reqDecoder.Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("couldn't decode request:", err)
		return
	}

	ipsum, err := svc.generateIpsum(req.Words, req.SentenceLength)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("couldn't generate ipsum:", err)
		return
	}
	respEncoder := json.NewEncoder(w)
	if err := respEncoder.Encode(IpsumResp{ipsum}); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("couldn't encode respose:", err)
		return
	}
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

type IpsumReq struct {
	Words          int
	SentenceLength int
}

type IpsumResp struct {
	Ipsum string
}
