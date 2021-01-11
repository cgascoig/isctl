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

OPENAPI_GENERATOR_CLI_IMAGE_TAG := @sha256:bcc4e88bd375b749b6b2555048f9853e8005829c0baa9394f9028e9bc5c224fe

all: build/isctl build/generator-postprocess openapi/operations.yaml cmd/cli.go
.PHONY: all

generate: cmd/cli.go
.PHONY: generate

clean:
> rm -Rf openapi build
.PHONY: clean

# Go unit tests
test: cmd/cli.go $(shell find cmd -name \*.go -type f) go.mod
> $(GO_CMD) test -v $(GO_MODULE)/cmd
.PHONY: test

# Run functional tests using BATS. These require bats, jq installed and working API keys so only run on dev workstation, not CI
functional-test: build/isctl
> bats tests
.PHONY: functional-test

openapi/operations.yaml: spec/intersight-openapi-v3.json generator-templates/go-experimental-generator-config.json $(shell find generator-templates -type f)
> mkdir -p $(@D)
> docker run --rm -v ${PWD}:/local openapitools/openapi-generator-cli$(OPENAPI_GENERATOR_CLI_IMAGE_TAG) generate  -g go-experimental -c /local/generator-templates/go-experimental-generator-config.json -i /local/spec/intersight-openapi-v3.json -o /local/openapi --template-dir /local/generator-templates/go-experimental
> rm openapi/go.mod openapi/go.sum
> mv openapi/README.md openapi/operations.yaml
> $(GO_PATH)/bin/goimports -l -w openapi

build/generator-postprocess: $(shell find generator-postprocess -name \*.go -type f) go.mod
> mkdir -p $(@D)
> go build -v -o "$@" $(GO_MODULE)/generator-postprocess 

cmd/cli.go: build/generator-postprocess openapi/operations.yaml generator-postprocess/cli.gotmpl generator-postprocess/types.gotmpl
> build/generator-postprocess --operations openapi/operations.yaml --template generator-postprocess/cli.gotmpl --output "$@"
> go fmt "$@"

build/isctl: cmd/cli.go $(shell find cmd -name \*.go -type f) go.mod
> $(GO_BUILD_CMD) -o "$@" $(GO_BUILD_FLAGS) $(GO_MODULE)/cmd

crossarch: cmd/cli.go $(shell find cmd -name \*.go -type f) go.mod
> GOOS=linux GOARCH=amd64 $(GO_BUILD_CMD) -o "build/isctl-linux_amd64" $(GO_BUILD_FLAGS) $(GO_MODULE)/cmd
> GOOS=windows GOARCH=amd64 $(GO_BUILD_CMD) -o "build/isctl-windows_amd64.exe" $(GO_BUILD_FLAGS) $(GO_MODULE)/cmd
> GOOS=darwin GOARCH=amd64 $(GO_BUILD_CMD) -o "build/isctl-darwin_amd64" $(GO_BUILD_FLAGS) $(GO_MODULE)/cmd
.PHONY: crossarch