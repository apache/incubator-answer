.PHONY: build clean ui

VERSION=1.1.0
BIN=answer
DIR_SRC=./cmd/answer
DOCKER_CMD=docker

GO_ENV=CGO_ENABLED=0 GO111MODULE=on
Revision=$(shell git rev-parse --short HEAD)
GO_FLAGS=-ldflags="-X github.com/answerdev/answer/cmd.Version=$(VERSION) -X 'github.com/answerdev/answer/cmd.Revision=$(Revision)' -X 'github.com/answerdev/answer/cmd.Time=`date +%s`' -extldflags -static"
GO=$(GO_ENV) $(shell which go)

build: generate
	@$(GO) build $(GO_FLAGS) -o $(BIN) $(DIR_SRC)

# https://dev.to/thewraven/universal-macos-binaries-with-go-1-16-3mm3
universal: generate
	@GOOS=darwin GOARCH=amd64 $(GO_ENV) $(GO) build $(GO_FLAGS) -o ${BIN}_amd64 $(DIR_SRC)
	@GOOS=darwin GOARCH=arm64 $(GO_ENV) $(GO) build $(GO_FLAGS) -o ${BIN}_arm64 $(DIR_SRC)
	@lipo -create -output ${BIN} ${BIN}_amd64 ${BIN}_arm64
	@rm -f ${BIN}_amd64 ${BIN}_arm64

generate:
	@$(GO) get github.com/google/wire/cmd/wire@v0.5.0
	@$(GO) get github.com/golang/mock/mockgen@v1.6.0
	@$(GO) install github.com/google/wire/cmd/wire@v0.5.0
	@$(GO) install github.com/golang/mock/mockgen@v1.6.0
	@$(GO) generate ./...
	@$(GO) mod tidy

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
