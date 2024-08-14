#
# Project Makefile
#
export GO111MODULE=on

.PHONY: setup help dep format lint vet build build-docker test test-coverage
.DEFAULT: help

dev-deps: ## Install Development dependencies
	npm i -g conventional-changelog-cli commitizen

dev: ## Setup the Development environment
	pre-commit install

fmt: ## Formats the go code using gofmt
	@gofmt -w -s .

lint: ## Lint code
	@revive -config revive.toml -formatter friendly $(PACKAGE)

vet: ## Run go vet
	@go vet $(PACKAGE)

build: ## Build the app
	@go build -o build/app logger.go

test: ## Run package unit testsS
	@go test -v -race -short ./...

test-coverage: ## Run tests with coverage
	@go test -short -coverprofile cover.out -covermode=atomic ./...

help: ## Displays help menu
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
