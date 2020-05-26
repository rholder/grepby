.PHONY: all build coverage clean fmt test

NAME=grepby
VERSION=1.3.0-dev

BUILD_DIR=build

all: clean build

## clean: remove the build directory
clean:
	rm -rfv $(BUILD_DIR)

## build: assemble the project and place a binary in build/
build:
	mkdir -p $(BUILD_DIR)
	cd cmd/$(NAME)/; go build -ldflags "-w -s -X main.Version=$(VERSION)" -o ../../$(BUILD_DIR)/$(NAME)
	@echo Build successful.

## fmt: run gofmt for the project
fmt:
	gofmt -w ./cmd/

## test: run the unit tests
test:
	mkdir -p $(BUILD_DIR)
	go test -v -coverprofile $(BUILD_DIR)/coverage.out ./cmd/...

## coverage: run the unit tests with test coverage output to build/coverage.html
coverage: test
	go tool cover -html=$(BUILD_DIR)/coverage.out -o $(BUILD_DIR)/coverage.html

## help: prints this help message
help:
	@echo "Usage: \n"
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'
