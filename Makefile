GO ?= go
GOPATH := $(CURDIR)/_vendor:$(GOPATH)

all: build

build:
	$(GO) build

clean:
	$(GO) clean
