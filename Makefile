## needs to generate docs
# include .env

Command := $(firstword $(MAKECMDGOALS))
# Skips the first word
Arguments := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))

PROJECTNAME=$(shell basename "$(PWD)")

# Go related variables.
GOBASE=$(shell pwd)
GOPATH=${HOME}/go
GOBIN=$(GOBASE)/bin
GOFILES=$(wildcard *.go)

# Redirect error output to a file, so we can show it in development mode.
STDERR=/tmp/.$(PROJECTNAME)-stderr.txt

# PID file will keep the process id of the server
PID=/tmp/.$(PROJECTNAME).pid

# Make is verbose in Linux. Make it silent.
MAKEFLAGS += --silent

## install: Install missing dependencies. Runs `go get` internally. e.g; make install get=github.com/foo/bar
install: go-get

## build-exchange-collector: Build the binary for cmd/exchange-scrapers/collector/collector
build-exchange-collector:
	@cd ./cmd/exchange-scrapers/collector && go build

## build-liquidity-scraper: Build the binary for cmd/liquidityScraper
build-liquidity-scraper:
	@cd ./cmd/liquidityScraper && go build

## build-asset-collector: Build the binary for cmd/assetCollectionService
build-asset-collector:
	@cd ./cmd/assetCollectionService && go build

## run-exchange-collector-service: Run exchange collector scraper
run-exchange-collector-service:
	@go run ./cmd/exchange-scrapers/collector/collector.go -exchange $(Arguments)

## run-liquidity-scraper-service: Run liquidity pools scraper
run-liquidity-scraper-service:
	@go run ./cmd/liquidityScraper/main.go -exchange $(Arguments)

## run-asset-collector-service: Run asset collector
run-asset-collector-service:
	@go run ./cmd/assetCollectionService/main.go -source $(Arguments)

## exec: Run given command, wrapped with custom GOPATH. e.g; make exec run="go test ./..."
exec:
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) $(run)

## clean: Clean build files. Runs `go clean` internally.
clean: go-clean

go-compile: go-clean go-get go-build

go-get:
	@echo "  >  Checking if there is any missing dependencies..."
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go get $(get)

go-install:
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go install $(GOFILES)

go-clean:
	@echo "  >  Cleaning build cache"
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go clean

.PHONY: help

all: help

help: Makefile
	@echo
	@echo " Choose a command run in "$(PROJECTNAME)":"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo