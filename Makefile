.PHONY: all build test clean run-app deps

GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
BINARY_NAME=hancock

all: deps test build
run: deps build run-app

deps:
	$(GOCMD) mod tidy

test:
	$(GOTEST) -v -race ./...

build:
	$(GOBUILD) -o $(BINARY_NAME) -v

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

run-app:
	./$(BINARY_NAME)
