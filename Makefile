.PHONY: lint test fmt tidy vet staticcheck

lint:
	golangci-lint run

fmt:
	golangci-lint run --fix

test:
	go test -v ./...


all: fmt lint test
