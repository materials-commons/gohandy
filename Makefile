.PHONY: all test fmt

all: fmt test

test:
	rm -rf test_data/t
	-go test -v ./...

fmt:
	-go fmt ./...
