.PHONY: run lint test clean build build-and-run

DEBUG ?= 1
ifeq ($(DEBUG), 1)
	LDFLAGS := '-extldflags="-static"'
else
	GOFLAGS += -trimpath
	LDFLAGS := '-s -w -extldflags="-static"'
endif

GOFLAGS := -ldflags ${LDFLAGS}

run:
	go run main.go

lint:
	golangci-lint run

test:
	go test ./day*

clean:
	rm -rf target

build: clean
	CGO_ENABLED=0 go build ${GOFLAGS} -v -o target/aoc2023

build-and-run: build
	target/aoc2023