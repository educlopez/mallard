.PHONY: build test lint vet vuln all

## build: compile all packages
build:
	go build ./...

## test: run all tests with race detector
test:
	go test -race ./...

## vet: run go vet
vet:
	go vet ./...

## lint: run golangci-lint
lint:
	golangci-lint run

## vuln: run govulncheck
vuln:
	go run golang.org/x/vuln/cmd/govulncheck@latest ./...

## all: build, vet, lint, test, vuln
all: build vet lint test vuln
