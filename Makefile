GO ?= go

all: build install

install:
	$(GO) install ./...

build:
	$(GO) build jcr/main.go
