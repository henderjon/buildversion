# built
A simple set of string vars defined at compile time for the `Version` of the build (a git tag and commit), `Timestamp` of the build (a timestamp), and `CompiledBy` (go version and platform).

As this package simply exports three getter funcs for three simple vars and setting them happens at compile time, usage is probably better if you copy and paste this code into your `package main`. This is in keeping with the proverb: "[A little copying is better than a little dependency.](https://go-proverbs.github.io/)"

## for example

A Makefile might look like...

```makefile
BIN=example
HEAD=$(shell git describe --dirty --long --tags 2> /dev/null  || git rev-parse --short HEAD)
TIMESTAMP=$(shell date '+%Y-%m-%dT%H:%M:%S %z %Z')

LDFLAGS="-X 'github.com/henderjon/built.version=$(HEAD)' -X 'github.com/henderjon/built.timestamp=$(TIMESTAMP)' -X 'github.com/henderjon/built.compiledBy=$(shell go version)'"

all: darwin64

.PHONY: dep
dep:
	go mod vendor

.PHONY: clean
clean:
	rm -f $(BIN) $(BIN)-*

.PHONY: darwin64
darwin64: clean
	env GOOS=darwin GOARCH=amd64 go build -ldflags $(LDFLAGS) -o $(BIN)-darwin64-$(HEAD)
```
