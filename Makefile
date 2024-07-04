EXECUTABLE=mpl
WINDOWS=./bin/$(EXECUTABLE).exe
LINUX=./bin/$(EXECUTABLE)_l
DARWIN=./bin/$(EXECUTABLE)
VERSION=./bin/$(shell git describe --tags --always --long --dirty)
MAINPATH=./main.go

.PHONY: all test clean

all: test build ## Build and run tests

test: ## Run unit tests
	go test ./...

build: windows linux darwin ## Build binaries
	@echo version: $(VERSION)

windows: $(WINDOWS) ## Build for Windows

linux: $(LINUX) ## Build for Linux

darwin: $(DARWIN) ## Build for Darwin (macOS)

$(WINDOWS):
	env GOOS=windows GOARCH=amd64 go build -o $(WINDOWS) $(MAINPATH)

$(LINUX):
	env GOOS=linux GOARCH=amd64 go build -o $(LINUX) $(MAINPATH)


$(DARWIN):
	env GOOS=darwin GOARCH=amd64 go build -o $(DARWIN) $(MAINPATH)

clean: ## Remove previous build
	rm -f $(WINDOWS) $(LINUX) $(DARWIN)

help: ## Display available commands
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
