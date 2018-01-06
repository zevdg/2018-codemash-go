# Goroutines and channels

In this lab, we will speed up our program by splitting up the work into multiple goroutines so that we can take advantage of multiple processor cores.

1. `go build` your current program and then rename the resulting binary to `ipsum-single-threaded`
1. In the `generateIpsum` function, after initializing the `WordBank`, make a channel of strings.
1. Inside the for loop, create an anonymous function with no input that calls `generateSentence` and pushes the result into the channel.  
*Note: functional buffs will recognize this as a closure*
1. Run this function in a goroutine by putting the `go` keyword directly before the function definition and `()` right after the function's closing curly brace.  

        // for example
        go func(){
            doStuff()
        }()

    *Note: Javascript buffs will recognize this (minus the go keyword) as an [Immediately Invoked Function Expression or IIFE](https://en.wikipedia.org/wiki/Immediately-invoked_function_expression)*
1. After this for loop, create another for loop that reads each sentence from the channel and builds them into the paragraph to be returned.
1. `go build` this version of the executable and compare the execution times.
    
        time ./ipsum-single-threaded
        time ./ipsum

    *Note: see [this stack overflow](https://superuser.com/questions/228056/windows-equivalent-to-unix-time-command) for equivalent non-unix commands.

