.PHONY: build clean ui

VERSION=0.0.1
BIN=answer
DIR_SRC=./cmd/answer
DOCKER_CMD=docker

GO_ENV=CGO_ENABLED=0
Revision=$(shell git rev-parse --short HEAD)
GO_FLAGS=-ldflags="-X main.Version=$(VERSION) -X main.Revision=$(Revision) -X 'main.Time=`date`' -extldflags -static"
GO=$(GO_ENV) $(shell which go)

build:
	@$(GO_ENV) $(GO) build $(GO_FLAGS) -o $(BIN) $(DIR_SRC)

generate:
	go get github.com/google/wire/cmd/wire@latest
	go generate ./...
	go mod tidy

test:
	@$(GO) test ./...

# clean all build result
clean:
	@$(GO) clean ./...
	@rm -f $(BIN)

install-ui-packages:
	@corepack enable
	@corepack prepare pnpm@v7.12.2 --activate

ui:
	@npm config set registry https://repo.huaweicloud.com/repository/npm/
	@cd ui && echo "REACT_APP_VERSION=$(VERSION)" >> .env && pnpm install && pnpm build && cd -

all: clean build
