.PHONY: help clean fmt fmt-check build run test check

help:
	@echo 'clean -- remove binary'
	@echo 'fmt -- gofmt'
	@echo 'fmt -- gofmt but no write'
	@echo 'vet -- go vet'
	@echo 'build -- go build'
	@echo 'check -- format and static check'
	@echo 'run -- execute'
	@echo 'test -- go test'

clean:
	go clean

fmt:
	gofmt -s -w -l ./

fmt-check:
	(! gofmt -s -d -e . | grep '^')

vet:
	go tool vet ./

lint:
	golint -set_exit_status ./

check:fmt-check vet lint

build:clean
	go build -x -v github.com/k-saka/lvlogger

run:build
	./slack-haven

install:clean
	go install github.com/k-saka/lvlogger

test:
	go test
