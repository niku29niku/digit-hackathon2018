# note: call scripts from /scripts
BIN := niku
PKG := github.com/numa08/niku

SRC_DIRS := cmd pkg

ARCH ?= linux-amd64
CLI_PLATFORMS := linux-arm64 darwin-amd64

all: clean rpi darwin
rpi:
	GOOS=linux GOARCH=arm GOARM=7 go build -v -o bin/linux/arm7/$(BIN) cmd/niku/main.go 

darwin:
	GOOS=darwin GOARCH=amd64 go build -v -o bin/darwin/amd64/$(BIN) cmd/niku/main.go

clean:
	@rm -rf bin

test:
	go test -v -short ./...

.PHONY: rpi darwin all clean test