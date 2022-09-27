.PHONY: build clean ui

VERSION=0.0.1
BIN=answer
DIR_SRC=./cmd/answer
DOCKER_CMD=docker

GO_ENV=CGO_ENABLED=0
GO_FLAGS=-ldflags="-X main.Version=$(VERSION) -X 'main.Time=`date`' -extldflags -static"
GO=$(GO_ENV) $(shell which go)

build:
	@$(GO_ENV) $(GO) build $(GO_FLAGS) -o $(BIN) $(DIR_SRC)

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
	@cd ui && pnpm install && pnpm build && cd -

all: clean build
