BUILD_DIR=./build

packages=./...

## all: run formatters, tests, check quality control and build the project
.PHONY: all
.DELETE_ON_ERROR:
all: tidy test check build

# ============================================================================ #
# HELPERS
# ============================================================================ #

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' $(MAKEFILE_LIST) | column -t -s ':' | sed -e 's/^/  /'

# ============================================================================ #
# QUALITY CONTROL
# ============================================================================ #

## tidy: format code and tidy modfile
.PHONY: tidy
tidy:
	go fmt $(packages)
	@echo
	go mod tidy -v

## check: run quality control checks (i.e. linters, etc)
.PHONY: check
check:
	go mod verify
	@echo
	go vet $(packages)
	@echo
	# simplify is NOT run by default by `go fmt`
	gofmt -d -s -e .

# ============================================================================ #
# DEVELOPMENT
# ============================================================================ #

## test: run all tests (set packages=<pkg> to run subset of tests)
.PHONY: test
test:
	go test $(addopts) $(packages)

## generate: run code generators
.PHONY: generate
generate:
	go generate $(addopts) $(packages)

## build: build the application
.PHONY: build
build:
	go generate ./...
	go build $(addopts) -o $(BUILD_DIR)/bin/pokedex-cli

## run: run the application
.PHONY: run
run: build
	@$(BUILD_DIR)/bin/pokedex-cli

## clean: clean up build artifacts
.PHONY: clean
clean:
	rm -rf $(BUILD_DIR)
	@echo
	go clean
