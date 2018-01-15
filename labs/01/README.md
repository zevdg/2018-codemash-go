# Basic Syntax

In this lab, we'll be creating a very basic lorem ipsum generating cli tool.

1. Create a folder at the path `<gopath>/src/github.com/<your-github-username>/ipsum`
1. Create an `ipsum.go` file in that folder with the following contents.

    ```go
    package main

    import (
        "fmt"
        "math/rand"
        "time"
    )

    var wordbank = []string{

    }

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
    ```
1. Populate the wordbank with some hardcoded strings.
1. Implement the `getWord` function to return a random word from your wordbank (see https://golang.org/pkg/math/rand/#Rand.Intn).  
*Note:`[0,n)` is mathspeak for "from 0 to n-1"*
1. Implement the `generateSentence` function to create a sentence with the length given by it's input by calling `getWord` in a loop.
1. Implement the `generateIpsum` function to return a paragraph of that many words with sentences of that length by calling `generateSentence` in a loop with the appropriate parameters.  
*Note: The last sentence may be shorter than sentenceLength.*

*Notice that the tooling automatically adds and removes import statements for you.  If this did not occur, make sure that your editor is configured to run goimports or goreturns on save.*
