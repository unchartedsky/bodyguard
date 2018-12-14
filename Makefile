NAME    := bodyguard
BIN			:= $(NAME)

SHELL := /bin/bash


# go source files, ignore vendor directory
SRC = $(shell find . -type f -name '*.go' -not -path "./vendor/*")

.PHONY: all
all: test

.PHONY: proto
proto:
	@ if ! which protoc > /dev/null; then \
		echo "error: protoc not installed, you might want to run 'brew install protobuf'." >&2; \
		exit 1; \
	fi
	go get -u -v github.com/golang/protobuf/protoc-gen-go
	go generate ./...

.PHONY: build
build: proto fmt
	@go build -a -o bin/$(BIN) .

.PHONY: test
test: build
	@go test

.PHONY: fmt
fmt: proto
	@ if ! which goimports > /dev/null; then \
		go get -u -v golang.org/x/tools/cmd/goimports; \
	fi

	go mod tidy
	goimports -l -w $(SRC)
	gofmt -l -w -s $(SRC)

.PHONY: run
run: build
	bin/$(BIN)
