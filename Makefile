.PHONY: test test-local test-remote

test-local:
	go test github.com/codecrafters-io/shell-starter-go/app

test-remote:
	codecrafters test

test: test-local test-remote

