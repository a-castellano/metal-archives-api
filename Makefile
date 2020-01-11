PROJECT_NAME := "metal-archives-wrapper"
PKG := "github.com/a-castellano/$(PROJECT_NAME)"
PKG_LIST := $(shell go list ${PKG}/... | grep -v /vendor/)
GO_FILES := $(shell find . -name '*.go' | grep -v /vendor/ | grep -v _test.go)

.PHONY: all dep build clean test coverage coverhtml lint

all: build

lint: prepare ## Lint the files
	@golint -set_exit_status ${PKG_LIST}

test: prepare ## Run unittests
	@go test -short ./...

race: dep ## Run data race detector
	@go test -race -short ${PKG_LIST}

msan: dep ## Run memory sanitizer
	@go test -msan -short ${PKG_LIST}

coverage: prepare ## Generate global code coverage report
	./scripts/coverage.sh;

coverhtml: prepare ## Generate global code coverage report in HTML
	./scripts/coverage.sh html;

dep: prepare ## Get the dependencies
	@dep ensure -v -vendor-only

build: dep ## Build the binary file
	@go build -i -v $(PKG)

clean: ## Remove previous build
	@rm -f $(PROJECT_NAME)

prepare: ## sets GOPATH and prepares CI
	./scripts/setup.sh;

help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
