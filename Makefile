SHELL := bash
.ONESHELL:
.SHELLFLAGS := -eu -o pipefail -c
# .DELETE_ON_ERROR:
MAKEFLAGS += --warn-undefined-variables
MAKEFLAGS += --no-builtin-rules

ifeq ($(origin .RECIPEPREFIX), undefined)
  $(error This Make does not support .RECIPEPREFIX. Please use GNU Make 4.0 or later)
endif
.RECIPEPREFIX = >

GO_MODULE := github.com/cgascoig/isctl
GO_CMD ?= go
GO_BUILD_CMD := $(GO_CMD) build -v 
GO_BUILD_FLAGS := -ldflags "-X main.commit=`git rev-parse HEAD`"
GO_PATH ?= $(shell go env GOPATH)

INTERSIGHT_SDK_VERSION := $(shell cat intersight-sdk-version)

all: build/isctl
.PHONY: all

generate: cmd/cli.go
.PHONY: generate

clean:
> rm -Rf build
> rm cmd/cli.go cmd/operations.go cmd/types.go
.PHONY: clean

# Go unit tests
test: cmd/cli.go $(shell find cmd -name \*.go -type f) go.mod
> $(GO_CMD) test -v $(GO_MODULE)/cmd
.PHONY: test

# Run functional tests using BATS. These require bats, jq installed and working API keys so only run on dev workstation, not CI
functional-test: build/isctl
> bats tests
.PHONY: functional-test

spec/openapi.yaml: intersight-sdk-version
> curl -o "$@" --location "https://github.com/CiscoDevNet/intersight-go/raw/$(INTERSIGHT_SDK_VERSION)/api/openapi.yaml"

go.mod: intersight-sdk-version
> go mod tidy
> touch "$@"

build/generator: $(shell find generator -name \*.go -type f) go.mod
> mkdir -p $(@D)
> go build -v -o "$@" $(GO_MODULE)/generator 

cmd/cli.go: build/generator $(shell find generator -name \*.go.tmpl -type f) spec/openapi.yaml
> build/generator --operations spec/openapi.yaml
> go fmt $(shell find cmd -name \*.go -type f)
> $(GO_PATH)/bin/goimports -l -w cmd

build/isctl: cmd/cli.go $(shell find cmd -name \*.go -type f) go.mod
> $(GO_BUILD_CMD) -o "$@" $(GO_BUILD_FLAGS) $(GO_MODULE)/cmd

crossarch: cmd/cli.go $(shell find cmd -name \*.go -type f) go.mod
> GOOS=linux GOARCH=amd64 $(GO_BUILD_CMD) -o "build/isctl-linux_amd64" $(GO_BUILD_FLAGS) $(GO_MODULE)/cmd
> GOOS=windows GOARCH=amd64 $(GO_BUILD_CMD) -o "build/isctl-windows_amd64.exe" $(GO_BUILD_FLAGS) $(GO_MODULE)/cmd
> GOOS=darwin GOARCH=amd64 $(GO_BUILD_CMD) -o "build/isctl-darwin_amd64" $(GO_BUILD_FLAGS) $(GO_MODULE)/cmd
.PHONY: crossarch
