.PHONY: build clean ui

VERSION=1.4.1
BIN=answer
DIR_SRC=./cmd/answer
DOCKER_CMD=docker

GO_ENV=CGO_ENABLED=0 GO111MODULE=on
Revision=$(shell git rev-parse --short HEAD 2>/dev/null || echo "")
GO_FLAGS=-ldflags="-X github.com/apache/incubator-answer/cmd.Version=$(VERSION) -X 'github.com/apache/incubator-answer/cmd.Revision=$(Revision)' -X 'github.com/apache/incubator-answer/cmd.Time=`date +%s`' -extldflags -static"
GO=$(GO_ENV) "$(shell which go)"

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
	@$(GO) get github.com/swaggo/swag/cmd/swag@v1.16.3
	@$(GO) install github.com/swaggo/swag/cmd/swag@v1.16.3
	@$(GO) install github.com/google/wire/cmd/wire@v0.5.0
	@$(GO) install github.com/golang/mock/mockgen@v1.6.0
	@$(GO) generate ./...
	@$(GO) mod tidy

image:
	@$(DOCKER_CMD) build . -t answerdev/answer:latest

test:
	@$(GO) test ./internal/repo/repo_test
test-e2e: image
	cd e2e && ./start.sh

# clean all build result
clean:
	@$(GO) clean ./...
	@rm -f $(BIN)

install-ui-packages:
	@corepack enable
	@corepack prepare pnpm@8.9.2 --activate

ui:
	@cd ui && pnpm pre-install && pnpm build && cd -

lint: generate
	@bash ./script/check-asf-header.sh
	@gofmt -w -l .

all: clean build
