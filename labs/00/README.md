# Workspace Setup

1. Verify go is working correctly `go version`
1. Verify git is working correctly `git --version`
1. Verify that your editor's go plugin is installed.
1. "go get" a hello world program `go get github.com/zevdg/codemash-2018-go/labs/0/hello`
1. cd to the directory `cd $(go env GOPATH)/src/github.com/zevdg/codemash-2018-go/labs/0/hello`
1. edit it
1. go run it `go run hello.go`
1. go build it `go build`
1. run the built executable `./hello`
1. go install it `go install`
1. run the installed version `$(go env GOPATH)/bin/hello`

*Bonus: print out an ascii picture*