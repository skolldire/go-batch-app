PROJECT_NAME := $(shell basename $(shell git rev-parse --show-toplevel 2>/dev/null || pwd))

.PHONY: all init build clean test run

all: init build test

init:
	@chmod +x init.sh && ./init.sh

build:
	go build -o $(PROJECT_NAME) ./cmd/cli/main.go

clean:
	rm -f $(PROJECT_NAME)

test:
	go test ./... -v

run:
	go run ./cmd/cli/main.go
