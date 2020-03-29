BINARY_NAME=server
ENTRY_POINT=./cmd/main.go

build:
	go build -o $(BINARY_NAME) $(ENTRY_POINT)

linter:
	golangci-lint run ./...

unit-test:
	mkdir results
	go test -coverprofile ./results/cover.out ./...
	go tool cover -html=./results/cover.out -o ./results/cover.html

run:
	go run $(ENTRY_POINT)

.PHONY: build, linter, run, unit-test
