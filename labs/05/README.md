# Error handing continued - DON'T PANIC!

In this lab, we'll learn how to handle our errors more gracefully by propogating them up the stack and adding context when it would be helpful.

1. Change the `NewWordBank` function to return both a `*WordBank` and an [error](https://golang.org/pkg/builtin/#error).
1. Replace the `panic` in the `NewWordBank` function to return the zero type for `*WordBank` (`nil`) and the error value.
1. Add `, nil` to all other returns in that function.
1. Change the `generateIpsum` function to also return an `error` type alongside it's current return type.
1. Return the error from `os.Open` when it is not `nil`  
*Note: the zero value for string is the empty string*
1. For the error returned from the constructor, use [fmt.Errorf](https://golang.org/pkg/fmt/#Errorf) to prepend some additional information to the error before returning it.
1. Add `, nil` to all other returns in that function.
1. In the main function, handle the error returned by generateIpsum by printing "Error: \<errorVal\>" and then returning.

*Bonus: Find an error that can be caused by malicious or nonsensical input and handle it as well.  New errors can be created with either `fmt.Errorf` as above, or somply [errors.New](https://golang.org/pkg/errors/#New) when a format string isn't needed.*