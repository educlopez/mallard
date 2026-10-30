.PHONY: build test lint lint-skills vet vuln all

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

## lint-skills: validate skill, command and agent frontmatter and README coverage
lint-skills:
	./scripts/lint-skills.sh

## vuln: run govulncheck
vuln:
	go run golang.org/x/vuln/cmd/govulncheck@latest ./...

## all: build, vet, lint, test, vuln
all: build vet lint lint-skills test vuln
