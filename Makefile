.PHONY: build clean ui

VERSION=1.0.1
BIN=answer
DIR_SRC=./cmd/answer
DOCKER_CMD=docker

#GO_ENV=CGO_ENABLED=0
Revision=$(shell git rev-parse --short HEAD)
GO_FLAGS=-ldflags="-X answercmd.Version=$(VERSION) -X 'answercmd.Revision=$(Revision)' -X 'answercmd.Time=`date`' -extldflags -static"
GO=$(GO_ENV) $(shell which go)

build:
	@$(GO_ENV) $(GO) build $(GO_FLAGS) -o $(BIN) $(DIR_SRC)

# https://dev.to/thewraven/universal-macos-binaries-with-go-1-16-3mm3
universal:
	@GOOS=darwin GOARCH=amd64 $(GO_ENV) $(GO) build $(GO_FLAGS) -o ${BIN}_amd64 $(DIR_SRC)
	@GOOS=darwin GOARCH=arm64 $(GO_ENV) $(GO) build $(GO_FLAGS) -o ${BIN}_arm64 $(DIR_SRC)
	@lipo -create -output ${BIN} ${BIN}_amd64 ${BIN}_arm64
	@rm -f ${BIN}_amd64 ${BIN}_arm64

generate:
	go get github.com/google/wire/cmd/wire@latest
	go install github.com/golang/mock/mockgen@v1.6.0
	go generate ./...
	go mod tidy

test:
	@$(GO) test ./internal/repo/repo_test

# clean all build result
clean:
	@$(GO) clean ./...
	@rm -f $(BIN)

install-ui-packages:
	@corepack enable
	@corepack prepare pnpm@v7.12.2 --activate

ui:
	@cd ui && pnpm install && pnpm build && cd -

all: clean build
