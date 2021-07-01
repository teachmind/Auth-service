SHELL=/bin/bash
APP=auth-service
APP_EXECUTABLE="./build/$(APP)"
APP_COMMIT=$(shell git rev-parse HEAD)
ALL_PACKAGES=$(shell go list ./... | grep -v "vendor")
SOURCE_DIRS=$(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))
COVERAGE_MIN=70

.PHONY: build

all: clean test

clean:
	@echo "> cleaning up the mess"
	@rm -rf build && mkdir -p build

lint:
	@echo "> running linter $(SOURCE_DIRS)/..."
	@golangci-lint run -v --timeout 5m $(SOURCE_DIRS)/...

build:
	@echo "> building binary"
	@go build -o $(APP_EXECUTABLE) -ldflags "-X main.commit=$(APP_COMMIT)"

migrate: build
	@echo "> running database migration"
	@${APP_EXECUTABLE} migrate

test:
	@echo "> running test and creating coverage report"
	go test -race -p=1 -cover -coverprofile=coverage.out $(ALL_PACKAGES)
	@go tool cover -html=coverage.out -o coverage.html
	@go tool cover -func=coverage.out | grep -i total:
	@go tool cover -func=coverage.out | gawk '/total:.*statements/ {if (strtonum($$3) < $(COVERAGE_MIN)) {print "ERR: coverage is lower than $(COVERAGE_MIN)"; exit 1}}'
	@rm *.out