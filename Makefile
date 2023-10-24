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

generate: pkg/gen/cli.go
.PHONY: generate

clean:
> rm -Rf build
> rm pkg/gen/cli.go pkg/gen/operations.go pkg/gen/types.go
.PHONY: clean

# Go unit tests
test: pkg/gen/cli.go $(shell find cmd pkg -name \*.go -type f) $(shell find cmd/extensions -name \*.py -type f) go.mod
> $(GO_CMD) test -v $(GO_MODULE)/...
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

pkg/gen/cli.go: build/generator $(shell find generator -name \*.go.tmpl -type f) spec/openapi.yaml
> build/generator --operations spec/openapi.yaml
> go fmt $(shell find pkg/gen -name \*.go -type f)
> $(GO_PATH)/bin/goimports -l -w pkg/gen

build/isctl: pkg/gen/cli.go $(shell find cmd pkg -name \*.go -type f) $(shell find cmd/extensions -name \*.py -type f) go.mod
> $(GO_BUILD_CMD) -o "$@" $(GO_BUILD_FLAGS) $(GO_MODULE)/cmd

crossarch: pkg/gen/cli.go $(shell find cmd pkg -name \*.go -type f) $(shell find cmd/extensions -name \*.py -type f) go.mod
> GOOS=linux GOARCH=amd64 $(GO_BUILD_CMD) -o "build/isctl-linux_amd64" $(GO_BUILD_FLAGS) $(GO_MODULE)/cmd
> GOOS=windows GOARCH=amd64 $(GO_BUILD_CMD) -o "build/isctl-windows_amd64.exe" $(GO_BUILD_FLAGS) $(GO_MODULE)/cmd
> GOOS=darwin GOARCH=amd64 $(GO_BUILD_CMD) -o "build/isctl-darwin_amd64" $(GO_BUILD_FLAGS) $(GO_MODULE)/cmd
.PHONY: crossarch
