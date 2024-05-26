
.DEFAULT_GOAL := help
.PHONY: cleanup build run-server-compiled

cleanup: ## Cleanup the project
	@rm -rf bin

test: ## Run the tests
	@go test -v ./pkg/...

run-server: ## Run the famcache server
	@go run cmd/server/main.go

run-server-compiled: ## Run the compiled famcache server
	@./bin/server

build: ## Build the famcache server
	@go build -o bin/server cmd/server/main.go

help:
	@echo "Usage: make [target]"
	@echo "Targets:"
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  %-20s %s\n", $$1, $$2}' $(MAKEFILE_LIST)