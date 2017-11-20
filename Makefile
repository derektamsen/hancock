.PHONY: all build test clean run deps build-macos

GOCMD=go
DEPCMD=dep
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
DEP_ENSURE=$(DEPCMD) ensure
BINARY_NAME=hancock
VENDOR_DIR=./vendor

all: deps test build
run: deps build run-app

deps:
	$(DEP_ENSURE) -v

test:
	$(GOTEST) -v ./...

build:
	$(GOBUILD) -o $(BINARY_NAME) -v

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -r $(VENDOR_DIR)

run-app:
	./$(BINARY_NAME)
