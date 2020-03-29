BINARY_NAME=server
ENTRY_POINT=./cmd/main.go

build:
	go build -o $(BINARY_NAME) $(ENTRY_POINT)

run:
	go run $(ENTRY_POINT)

.PHONY: build, run, compile