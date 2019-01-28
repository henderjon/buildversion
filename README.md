# buildversion
a simple set of string vars defined at compile time to for buildVersion and buildTimestamp

## for example

A Makefile might look like...

```makefile
BIN=example
HEAD=$(shell git describe --dirty --long --tags 2> /dev/null  || git rev-parse --short HEAD)
TIMESTAMP=$(shell date '+%Y-%m-%dT%H:%M:%S %z %Z')

LDFLAGS="-X main.buildVersion=$(HEAD) -X 'main.buildTimestamp=$(TIMESTAMP)'"

all: darwin64

.PHONY: dep
dep:
	go mod vendor

.PHONY clean
clean:
	rm -f $(BIN) $(BIN)-*

.PHONY: darwin64
darwin64:
	env GOOS=darwin GOARCH=amd64 go build -ldflags $(LDFLAGS) -o $(BIN)-darwin64-$(HEAD)
```
