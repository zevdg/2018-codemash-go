# Interfaces, defer, and panic!

In this lab, we'll be using the [os](https://golang.org/pkg/os/) and [bufio](https://golang.org/pkg/bufio/) packages to read our wordbank from a file on disk.  Since reading a file can fail, we will also be learning the most basic error handling strategy.

1. Add a new flag "wordbank" that accepts a filepath as a string.
1. Modify the `NewWordBank` function to take an [*os.File](https://golang.org/pkg/os/#File) as input
1. Create a [bufio.Scanner](https://golang.org/pkg/bufio/#Scanner) by passing the `*os.File` into the [bufio.NewScanner](https://golang.org/pkg/bufio/#NewScanner) constructor.  
*Note: the `bufio.NewScanner` function expects an [io.Reader](https://golang.org/pkg/io/#Reader) as its input.  Fortunately, `*os.File` implements the `io.Reader` interface.*
1. Loop over the lines in the file and use the [builtin append function](https://golang.org/pkg/builtin/#append) to add them to your wordbank.  
*Hint: see [this example](https://golang.org/pkg/bufio/#example_Scanner_lines) from the scanner documentation for inspiration.*
1. If `scanner.Err()` is not `nil`, call [panic](https://golang.org/pkg/builtin/#panic) on the error value.
1. Move the words from the old hardcoded slice into a file on disk.
1. Modify the `generateIpsum` function to [os.Open](https://golang.org/pkg/os/#Open) the file for reading
1. If the error value returned by `os.Open` is not `nil`, call `panic` on the error value
1. `defer` the [close method](https://golang.org/pkg/os/#File.Close) on the file object immediately after handling the error.  
*Hint: see [this example from effective go](https://golang.org/doc/effective_go.html#defer)*
1. pass the *os.File instance into the `NewWordbank` function.
