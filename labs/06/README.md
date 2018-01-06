# Packages

In this lab, we will move our wordbank into a separate package and then import and use it in our main package.

1. Create a folder called `wordbank` next to your `ipsum` folder, and create a `wordbank.go` file starting with the line `package wordbank`.
1. Move the `WordBank` type, methods, and constructor from `ipsum.go` into `wordbank.go`.
1. In `ipsum.go`, prepend `wordbank.` to call of the `NewWordBank` constructor and save the file.  
*Notice that the tooling found your `wordbank` package and added the approprate imports just as it did with the standard library packages.  Goimports will work for any package in your GOPATH.*
1. The squigly red underline that just appeared under the call to `getWord` is telling you that you are trying to call an unexported method from a separate package.  Export the `getWord` function by captializing the first letter.
1. The program is technically complete now, but [Go naming conventions](https://blog.golang.org/package-names) recomend to avoid "stutter" in package contents like `wordbank.NewWordBank()` when possible.  Renaming the `NewWordBank` function to just `New` will make the calling code read as `wordbank.New()`.  This avoids the stutter and is an idiomatic contructor name for packages that are built around a single type.