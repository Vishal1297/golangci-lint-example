SHELL := /bin/bash

.PHONY: help
## help: shows this help message
help:
	@ echo "Usage: make [target]"
	@ sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

# ==================================================================================== #
# QUALITY CONTROL
# ==================================================================================== #

## tidy: format code and tidy modfile
.PHONY: tidy
tidy:
	go fmt ./...
	go mod tidy -v

## golangci-lint: runs installed linter for all packages
.PHONY: golangci-lint
golangci-lint:
	@ echo "[$@] Running pre-installed golangci-lint"
	@ echo "[$@] Check installed version"
	golangci-lint --version
	golangci-lint run -v .
	@ echo "[$@] Linting down without any issues"
	
.PHONY: lint
## lint: runs linter for a given directory, specified via PACKAGE variable
lint: 
	@ if [ -z "$(PACKAGE)" ]; then echo >&2 please set directory via variable PACKAGE; exit 2; fi
	@ docker run  --rm -v "`pwd`:/workspace:cached" -w "/workspace/$(PACKAGE)" golangci/golangci-lint:latest golangci-lint run -v .
	@ echo "[$@] Linting down without any issues"

.PHONY: lint-all
## lint-all: runs linter for all packages
lint-all: 
	@ docker run  --rm -v "`pwd`:/workspace:cached" -w "/workspace/." golangci/golangci-lint:latest golangci-lint run -v .
	@ echo "[$@] Linting down without any issues"

# ==================================================================================== #
# DEVELOPMENT
# ==================================================================================== #

## build: build the application
.PHONY: build
build:
## Include additional build steps, like TypeScript, SCSS or Tailwind compilation here...
	go build -o=/tmp/bin/${BINARY_NAME} ${MAIN_PACKAGE_PATH}

## run: run the  application
.PHONY: run
run: build
	/tmp/bin/${BINARY_NAME}
	go run main.go

## test: run all tests
.PHONY: test
test:
	go test -v -race -buildvcs ./...

## test/cover: run all tests and display coverage
.PHONY: test/cover
test/cover:
	go test -v -race -buildvcs -coverprofile=/tmp/coverage.out ./...
	go tool cover -html=/tmp/coverage.out
