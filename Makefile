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

all: build/isctl
.PHONY: all

clean:
> rm -Rf build
.PHONY: clean

# Go unit tests
test: $(shell find cmd pkg -name \*.go -type f) $(shell find cmd/extensions -name \*.py -type f) go.mod pkg/oapi/intersight-openapi.json
> $(GO_CMD) test -v $(GO_MODULE)/...
.PHONY: test

# Run functional tests using BATS. These require bats, jq installed and working API keys so only run on dev workstation, not CI
functional-test: build/isctl
> bats tests
.PHONY: functional-test

pkg/oapi/intersight-openapi.json: intersight-sdk-version
> VERSION=`cat intersight-sdk-version | sed -r -e 's/^"v([0-9]+)\.([0-9]+)\.([0-9]+)\.([0-9]+)"$$/\1.\2.\3-\4/'`
> echo "Using API version $${VERSION}"
> curl -o "$@" --location "https://cdn.intersight.com/components/an-apidocs/$${VERSION}/model/intersight-openapi-v3-$${VERSION}.json"


build/isctl: $(shell find cmd pkg -name \*.go -type f) $(shell find cmd/extensions -name \*.py -type f) go.mod pkg/oapi/intersight-openapi.json
> $(GO_BUILD_CMD) -o "$@" $(GO_BUILD_FLAGS) $(GO_MODULE)/cmd

crossarch: $(shell find cmd pkg -name \*.go -type f) $(shell find cmd/extensions -name \*.py -type f) go.mod pkg/oapi/intersight-openapi.json
> GOOS=linux GOARCH=amd64 $(GO_BUILD_CMD) -o "build/isctl-linux_amd64" $(GO_BUILD_FLAGS) $(GO_MODULE)/cmd
> GOOS=windows GOARCH=amd64 $(GO_BUILD_CMD) -o "build/isctl-windows_amd64.exe" $(GO_BUILD_FLAGS) $(GO_MODULE)/cmd
> GOOS=darwin GOARCH=amd64 $(GO_BUILD_CMD) -o "build/isctl-darwin_amd64" $(GO_BUILD_FLAGS) $(GO_MODULE)/cmd
.PHONY: crossarch
