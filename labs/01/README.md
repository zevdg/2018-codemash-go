# Basic Syntax

In this lab, we'll be creating a very basic lorem ipsum generating cli tool.

1. Copy this `ipsum` folder to `<gopath>/src/github.com/<your-github-username>/ipsum`.
1. Populate the wordbank with some hardcoded strings.
1. Implement the `getWord` function to return a random word from your wordbank (see https://golang.org/pkg/math/rand/#Rand.Intn).  
1. Implement the `generateSentence` function to create a sentence with the length given by it's input by calling `getWord` in a loop.
1. Implement the `generateIpsum` function to return a paragraph of that many words with sentences of that length by calling `generateSentence` in a loop with the appropriate parameters. 
*Note: The last sentence may be shorter than sentenceLength.*

*Notice that the tooling automatically found and imported the time and rand packages for you.  If this did not occur, make sure that your editor is configured to run goimports or goreturns on save.*