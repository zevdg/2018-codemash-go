# Intro to Go
From zero to a production quality webservice

Note: talk about self.  6 Years at Bloomberg, 

----

## Outline

```go
func Precompile(){
    Intro()
    //Labs
    for i := 0; i <= 10; i++ {
        ShowSlides(i)
        DoLab(i)
        uploadSolution(i)
        WalkthroughSolution(i)
        if i == 6 {
            StretchLegs()
        }
    }
    TakeQuestions()
}

```

https://github.com/zevdg/2018-codemash-go

----

## Labs 0-6
* Environment setup
* Basic Syntax
* Documentation and pointers
* Structs, Methods, and Constructors
* IO, defer, panic
* Error handing continued - DON'T PANIC!
* Packages

----

## Labs 7-10
* Concurrency, goroutines, and channels
* The net/http package
* Service scoped variables
* JSON Encoding

----

## Secret Bonus Labs

* Polymorphism without inheritance
* Testing with `go test`

----

## Course Calibration

Note: Poll the audience about language knowledge: Java/C#, C, C++, Python, Javascript.  Poll the audience about go experience: Raise your hands if you've written any go at all, lower your hand if you've 

---

# Install and/or test your prerequisites right now!

---

## The History of Go
Simple by design

Note: Robert Griesemer, Rob Pike, and Ken Thompson.  Avoid design by comittee.  Unanimous approval.

----

## Pain Points

* slow builds
* uncontrolled dependencies
* each programmer using a different subset of the language
* poor program understanding (code hard to read, poorly documented, and so on)
* difficulty of writing automatic tools
* cross-language builds

----

## Key Features 

* Fast compilation
* Low-Latency garbage collection
* Single deployable static binary
* Minimally cumbersome type system
* Paralellism in the core language
* Batteries included

----

## Ideology

Go is defined more by the features it excludes  
than by the features it has.  

(e.g. inheritance, exceptions, and generics)

----

## Ideology (cont.)
#### Consistancy
* Opinionated tooling
* Convention over configuration
* Encourage "one way to do it"

Note: gofmt, no-one's favorite and everyone's favorite.  Only one knob for the garbage collector

----

## Ideology (cont.)
#### Avoid "magic"
* Libraries over frameworks
* Explicit is better than implicit
* No preprocessor / Macro system

Note: no getter/setter properties, no operator overloading.

----

## Ideology (cont.)
#### Readability
* Easy to read > easy to write
* Keep error handling near the error
* Avoid excessive verbosity (or terseness)

---

# Lab 0
## Workspace Setup

----

## $GOPATH

$GOPATH/  <!-- .element: style="text-align: left;"> -->  
├─ pkg/  
├─ bin/  
└─ src/

----

## $GOPATH/src

src/  <!-- .element: style="text-align: left;"> -->  
├─ github.com/  
├─ user-or-org-name/  
│&nbsp;&nbsp;&nbsp;&nbsp;├─ repo1  
│&nbsp;&nbsp;&nbsp;&nbsp;└─ repo2  
├─ gitlab.com/  
└─ internal.mycorp.io/  

----

## Go command

* go get
* go run
* go build
* go install
* go test
* (and many more)

----

# Lab time!
https://github.com/zevdg/2018-codemash-go/tree/master/labs/00

---

# Lab 1
## Basic Syntax

----

## Variables

The `var` keyword

```go
var foo bool
```
```go
var minHeight, maxHeight int
```
```go
var (
    name string
    age int
)
```

----

## Variables (cont.)

The `:=` operator

```go
foo := false
```
```go
minHeight, maxHeight := 100, 800
```
```go
name, age := "John", 47
```
```go
retVal := someFunc()
```
Note: Like auto in C++.

----

## Basic Types

* bool
* string
* int, int8, int16, int32, int64
* uint, uint8, uint16, uint32, uint64, uintptr
* byte (alias for uint8)
* rune (alias for int32)
    * represents a Unicode code point
* float32 float64
* complex64 complex128

----

## Slices

```go
var foo []int //nil
```
```go
bar := make([]string, 0)
```
```go
baz := []string{
    "the",
    "final",
    "countdown",
}

// Read
firstWord := baz[0]
section := baz[1:3]

// Write
baz = append(baz, "da-na-na-nah")
baz = append(baz, "da-na-nuh-nuh-nuh!")
```

----

## Maps

```go
var foo map[int]bool // nil
```
```go
bar := make(map[int]bool)
```
```go
baz := map[string]string {
    "first" : "partridge in a pair tree",
    "second": "turtle doves",
    "third" : "french hens",
    "fourth": "calling birds",
    "fifth" : "gold rings",
}

// Read
day5 := baz["fifth"]
day6, ok := baz["sixth"]

//Write
baz["first"] = "Nintendo Switch"

```

----

## if

```go
if theAnswer == 42 {
    doStuff()
}
```

```go
if err := doSomething() ; err != nil {
    log.Fatal(err)
}
```

----

## Loops

for
```go
for i := 0; i < len(myList); i++ {
    ...
}
```

```go
for key, val := range mySliceOrMap {
    ...
}
```
while
```go
for iter.HasNext() {
    ...
}
```
```go
for {
    ...
} 
```

----

## Functions

```go
func simple() {
    doStuff()
}
```
```go
func inputExample(input1 int, input2, input3 string){
    doStuff()
}
```
```go
func outputExample() int {
    return 3
}
```
```go
func multipleOutputExample() (int, string) {
    return 1, "is the loneliest number"
}
```

----

# Lab time!
https://github.com/zevdg/2018-codemash-go/tree/master/labs/01

---

# Lab 2
## Documentation and Pointers

----

## Documentation

* standard lib only - https://golang.org/pkg/
* stdlib + public 3rd party - https://godoc.org
* local and/or private - [go doc](https://golang.org/cmd/go/#hdr-Show_documentation_for_package_or_symbol)

Note: generated from code comments

----

## Pass by Value vs Pass By Reference

```go
func main(){
    myInt := 3
    incrementVal(myInt)
    fmt.Println(myInt)
    incrementPtr(&myInt)
    fmt.Println(myInt)
}

func incrementVal(val int){
    val++
}

func incrementPtr(ptr *int){
    *ptr++
}
```

----

## Pointer Syntax

#### Type declarations

A `*` before a type name indicates that type is a pointer

```go
var foo *int
```
```go
func bar(ptr *string){}
```

----

## Pointer Syntax (cont.)

#### Operators
Reference: A `&` before a variable returns a pointer pointing to that variable.

```go
myVar := "harambe"
expectsPointer(&myVar)
```

Dereference: A `*` before a varible returns the value the pointer points to 
```go
var myVar *int
expectValue(*myPtr)
```

----

## Pointers in practice

Not as obtrusive as you'd think

```go
truck := NewTruck()
truck.Color = "Red"
truck.Drive()
```

Note: Ask C/C++ devs if truck is a pointer or a value

----

# Lab time!
https://github.com/zevdg/2018-codemash-go/tree/master/labs/02

---

# Lab 3
## Custom Types

----

## Simple custom types

```go
type farenheight float32
type celcius float32

func isFreezing(temp celcius) bool {
    return temp < 0
}

func main(){
    temp := farenheight(20)
    if isFreezing(temp) { // compile error
        abortMission()
    }
}
```

----

## Complex custom types

```go
type Doggo struct {
    breed string
    name  string
    age   int
}

dog1 := Doggo{}
dog2 := Doggo{"pitt bull", "Suzy", 2}
dog3 := Doggo{
    name: "Thor",
    age: 1,
}
dogPtr := &Doggo{}
```

----

## Methods
```go
type Doggo struct {
    breed string
    age   int
}

func (dog *Doggo)Bark() string {
    if dog.breed == "Shibe" { return "Such wow!" }
    return "Woof"
}

func (dog *Doggo)Grow(){
    dog.Age++
}
```

----

# Lab time!
https://github.com/zevdg/2018-codemash-go/tree/master/labs/03

---

# Lab 4
## Iterfaces, defer, and panic!

----

## Interfaces

```go
type Woofer interface {
    Woof()
}
```
```go
type SingerSongwriter interface {
    Sing(s Song)
    Write() Song
}
```

----

## The error interface

```go
type error interface {
    Error() string
}
```
Creating errors
```go
error1 := errors.New("this is an error")
error2 := fmt.Errorf("error number %d", 2)
```
Error handling
```go
result, err := someCalculation()
if err != nil {
    //handle error
}
```

----

## Panic
```go
result, err := someCalculation()
if err != nil {
    panic(err)
}
```

----

## The IO interfaces

```go
type Reader interface {
    Read(p []byte) (n int, err error)
}
```
```go
type Writer interface {
    Write(p []byte) (n int, err error)
}
```
```go
type Closer interface {
    Close() error
}
```

----

## Defer

```go
func someFunc(filepath string) bool {
    f, err := os.Open(filepath)
    if err != nil {
        panic(err)
    }
    defer f.Close()
    if foo() {
        return false
    }
    if bar() {
        panic()
    }
    return true
}
```

----

## The empty interface

```go
var i interface{}
```

```go
func debug(i interface{}) {
    fmt.Printf("val:%v, type:%T\n", i, i)
}

func main(){
    debug("myString")
    debug(true)
    debug(77)
}
```

----

# Lab time!
https://github.com/zevdg/2018-codemash-go/tree/master/labs/04

---

# Lab 5
## Error handing continued 
## DON'T PANIC!

----

## When panic makes sense

* Debugging during development or
* Programmer error
    * nil pointer dereference
* Unrecoverable error
    * out of memory
* Unrecoverable Programmer error
    * race conditions

----

## When not to panic

Literally, all other times  
But especially across package boundries

----

## Propagating errors
Without context
```go
func foo() (string, error)
    result, err := someCalculation()
    if err != nil {
        return "", err
    }
    return result, nil
```
With context
```go
func foo() (string, error)
    result, err := someCalculation()
    if err != nil {
        return "", fmt.Errorf("calculation failed: %s", err)
    }
    return result, nil
```

----

## Handling Errors

Handle the error in-place
```go
result, err := someCalculation()
if err != nil {
    recoverFrom(err)
}
doNextThing()
```
Intentionally ignore the error
```go
// comment explaining why this error was ignored
result, _ := someCalculation()

```

Note: There's a special place in hell for people who ignore errors without comments

----

# Lab time!
https://github.com/zevdg/2018-codemash-go/tree/master/labs/05

---

# Lab 6
## Packages

----

## Types of packages

* main package - buildable
* other packages - importable

----

## Package declaration

All go files must start with

```go
// optional package comment
package <packageName>
```
Naming Conventions (non-main)
* short and clear
* no under_scores or mixedCaps
* avoid obvious collisions
* examples from the stdlib
    * time, list, http, strconv, syscall, fmt

----

## Importing packages

Packages are imported by their "import paths"
```go
import (
    "fmt"

    "github.com/pkg/errors"
    log "github.com/sirupsen/logrus"

    "corp.intranet.org/pkgs/can/be"
    "corp.intranet.org/pkgs/can/be/nested/inside/eachother"
)
```

Note: generally, the containing folder or repo shoudl match the package name

----

## Exported names

```go
package mylib

import "fmt"

var SocialGopher *Gopher //exported
var shyGopher *Gopher

type lemming struct{}
func newLemming() *lemming { return &lemming{} }

type Gopher struct{} //exported
func NewGopher() *Gopher { return &Gopher{} } //exported
func(g *Gopher) Laugh(){ fmt.Println("haha") } //exported
func(g *Gopher) cry() { fmt.Println("wah!") }

```

----

## Exported names (cont.)

```go
package mylib

import "fmt"

var SocialGopher *Gopher

type Gopher struct{}
func NewGopher() *Gopher
func(g *Gopher) Laugh()

```

----

# Lab time!
https://github.com/zevdg/2018-codemash-go/tree/master/labs/06

---

# Lab 7
## goroutines and channels

----

## The go keyword

```go
func slowPrint(){
    time.Sleep(300 * time.Millisecond)
    fmt.Println("after 300ms")
}

func main(){
    go slowPrint()
    go func(){
        time.Sleep(100 * time.Millisecond)
        fmt.Println("after 100ms")
    }()
    fmt.Println("immediately")
    time.Sleep(500 * time.Millisecond)
}
```

----

## Channels
Don't communicate by sharing memory, share memory by communicating.

```go
var myChan chan int
```
```go
c := make(chan string)
go func(){
    time.Sleep(500 * time.Millisecond)
    c <- "a very important message" //send message to channel
}()
msg := <- c //recieving from channel
fmt.Println("recieved:", msg)
```

----

## Synchronization

```go
func main() {
	ch := make(chan int)
	ch <- 1
	ch <- 2 //this code will block foerever here
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
```

----

## Buffered Channels

```go
func main() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2 //all good now
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
```

----

# Lab time!
https://github.com/zevdg/2018-codemash-go/tree/master/labs/07

---

# Lab 8
## The net/http package

----

## Sending http requests

Get
```go
resp, err := http.Get("http://example.com/")
if err != nil {
	// handle error
}
defer resp.Body.Close()
body, err := ioutil.ReadAll(resp.Body)
// ...
```

Post
```go
resp, err := http.PostForm("http://example.com/form",
	url.Values{"key": {"Value"}, "id": {"123"}})
```

----

## Serving http handlers

```go
func HelloHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "hello, world!\n")
}

func main() {
    http.HandleFunc("/hello", HelloHandler)

    //blocks indefinitly unless there is an error
    err := http.ListenAndServe(":12345", nil)

    fmt.Println("error:", err)
}
```

----

# Lab time!
https://github.com/zevdg/2018-codemash-go/tree/master/labs/08

---

# Lab 9
## Service scoped variables

----

![smelly smell](smelly.jpg)

----

## Real world examples

* database handles
* loggers

----

## Gobal variables

```go
var msg string

func main() {
    msg = "hello, world!\n"
    http.HandleFunc("/hello", HelloHandler)
    fmt.Println("error:", http.ListenAndServe(":12345", nil))
}

func HelloHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, msg)
}
```

----

## Methods as handlers

```go
type Env struct {
    db *sql.DB
}

func main() {
    env := NewEnv() //intitalize svc scoped vars here
    http.HandleFunc("/endpoint", env.Handle)
    fmt.Println("error:", http.ListenAndServe(":12345", nil))
}

func(env *Env)Handle(w http.ResponseWriter,r *http.Request) {
	fmt.Fprint(readRecords(env.db))
}
```

----

# Lab time!
https://github.com/zevdg/2018-codemash-go/tree/master/labs/09

---

# Lab 10
## The encoding/json package

----

## Marshal and Unmarshal

```go
package json
func Marshal(v interface{}) ([]byte, error){ ... }
func Unmarshal(data []byte, v interface{}) error { ... }
```
Usage
```go
type Doggo struct {
    Name string
    Age int
}
b, err := json.Marshal(Doggo{"Alice", 3})
handle(err)
// b should be a byte slice of
// `{"Name":"Alice", "Age":3}`

var myDoge Doggo
err = json.Unmarshal(b, &m)
handle(err)
```

----

## Customize JSON

Capitalization

```go
type Doggo struct {
    name string
    Age int
}
b, err := json.Marshal(Doggo{"Alice", 3})
handle(err)
// b will be a byte slice of
// `{"Age":3}`
```

----

## Struct tags

```go
type Doggo struct {
    Name string `json:"name"`
    Age int
}
b, err := json.Marshal(Doggo{"Alice", 3})
handle(err)
// b will be a byte slice of
// `{"name":"Alice", "Age":3}`
```

----

## Working with streams

```go
package json

func NewDecoder(r io.Reader) *Decoder { ... }
func NewEncoder(w io.Writer) *Encoder { ... }

func (dec *Decoder) Decode(v interface{}) error { ... }
func (enc *Encoder) Encode(v interface{}) error { ... }
```
Usage
```go
dec := json.NewDecoder(os.Stdin)
enc := json.NewEncoder(os.Stdout)

v := map[string]interface{}

err := dec.Decode(&v)
handle(err)

err := enc.Encode(&v)
handle(err)
```

----

# Lab time!
https://github.com/zevdg/2018-codemash-go/tree/master/labs/10

---

# Questions?

----

## Zev Goldstein

zevdg on [github](https://github.com/zevdg), [linkedin](https://www.linkedin.com/in/zevdg/), [reddit](https://www.reddit.com/user/zevdg/), and [twitter](https://twitter.com/zevdg)

Please fill our the session survey on the codemash app  
and/or tweet at me if you liked the presentation
