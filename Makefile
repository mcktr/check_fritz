VERSION := $(shell git describe --tags --always)
BUILD := go build -v -ldflags "-s -w -X main.Version=$(VERSION)"

BINARY = check_fritz

.PHONY : all clean build test

all: build test

test:
	go test -v ./...

clean:
	rm -rf build/

build:
	mkdir -p build
	GOOS=linux GOARCH=amd64 $(BUILD) -o build/$(BINARY).linux.amd64 ./cmd/check_fritz
	sha256sum build/$(BINARY).linux.amd64 > build/$(BINARY).linux.amd64.sha256
	GOOS=linux GOARCH=arm64 $(BUILD) -o build/$(BINARY).linux.arm64 ./cmd/check_fritz
	sha256sum build/$(BINARY).linux.arm64 > build/$(BINARY).linux.arm64.sha256
	GOOS=linux GOARCH=arm $(BUILD) -o build/$(BINARY).linux.arm ./cmd/check_fritz
	sha256sum build/$(BINARY).linux.arm > build/$(BINARY).linux.arm.sha256
	GOOS=windows GOARCH=amd64 $(BUILD) -o build/$(BINARY).windows.amd64.exe ./cmd/check_fritz
	sha256sum build/$(BINARY).windows.amd64.exe > build/$(BINARY).windows.amd64.exe.sha256
	GOOS=darwin GOARCH=amd64 $(BUILD) -o build/$(BINARY).darwin.amd64 ./cmd/check_fritz
	sha256sum build/$(BINARY).darwin.amd64 > build/$(BINARY).darwin.amd64.sha256
