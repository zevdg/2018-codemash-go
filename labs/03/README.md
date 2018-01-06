# Custom Types

In this lab, we'll be turning our simple wordbank into a user defined type.

1. Create a `WordBank` struct with a single field: a slice of strings.
1. Create a function `NewWordBank` with no input that returns a pointer to an initialized WordBank.
1. Change getWord to be a method on the `*WordBank` type.  Modify it to use the fields of `WordBank` instead of the package scoped string slice.
1. Modify the `generateSentence` function to take a `*WordBank` as an additional input and use it to generate the words
1. Modify the `generateIpsum` function to create an instance of this `WordBank` type by calling the constructor and pass it into `generateSentence`.
