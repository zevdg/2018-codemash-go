package main

import (
	"fmt"
	"math/rand"
	"time"
)

var wordbank = []string{}

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println(generateIpsum(100, 6))
}

func generateIpsum(wordCount, sentenceLength int) string {

}

func generateSentence(wordCount int) string {

}

func getWord() string {

}
