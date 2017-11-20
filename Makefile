GO ?= go
DEP ?= dep
GOPATH := $(CURDIR)/_vendor:$(GOPATH)

all: build

build:
	$(DEP) ensure
	$(GO) build

clean:
	$(GO) clean
	rm -r ./vendor/
