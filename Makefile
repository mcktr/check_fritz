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
	GOOS=linux GOARCH=arm64 $(BUILD) -o build/$(BINARY).linux.arm64 ./cmd/check_fritz
	GOOS=linux GOARCH=arm $(BUILD) -o build/$(BINARY).linux.arm ./cmd/check_fritz
	GOOS=windows GOARCH=amd64 $(BUILD) -o build/$(BINARY).windows.amd64.exe ./cmd/check_fritz
	GOOS=darwin GOARCH=amd64 $(BUILD) -o build/$(BINARY).darwin.amd64 ./cmd/check_fritz
	cd build; sha256sum $(BINARY).linux.amd64 > $(BINARY).linux.amd64.sha256
	cd build; sha256sum $(BINARY).linux.arm64 > $(BINARY).linux.arm64.sha256
	cd build; sha256sum $(BINARY).linux.arm > $(BINARY).linux.arm.sha256
	cd build; sha256sum $(BINARY).windows.amd64.exe > $(BINARY).windows.amd64.exe.sha256
	cd build; sha256sum $(BINARY).darwin.amd64 > $(BINARY).darwin.amd64.sha256

test-checksum:
	cd build; sha256sum -c $(BINARY).linux.amd64.sha256
	cd build; sha256sum -c $(BINARY).linux.arm64.sha256
	cd build; sha256sum -c $(BINARY).linux.arm.sha256
	cd build; sha256sum -c $(BINARY).windows.amd64.exe.sha256
	cd build; sha256sum -c $(BINARY).darwin.amd64.sha256
