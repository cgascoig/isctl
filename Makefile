SHELL := bash
.ONESHELL:
.SHELLFLAGS := -eu -o pipefail -c
.DELETE_ON_ERROR:
MAKEFLAGS += --warn-undefined-variables
MAKEFLAGS += --no-builtin-rules

ifeq ($(origin .RECIPEPREFIX), undefined)
  $(error This Make does not support .RECIPEPREFIX. Please use GNU Make 4.0 or later)
endif
.RECIPEPREFIX = >

GO_MODULE := github.com/cgascoig/isctl
GO_CMD ?= go
GO_BUILD_CMD := $(GO_CMD) build -v 

all: build/isctl build/generator-postprocess openapi/operations.yaml cmd/cli.go
.PHONY: all

clean:
> rm -Rf openapi build
.PHONY: clean

openapi/operations.yaml: spec/intersight-openapi-v3.json generator-templates/go-experimental-generator-config.json $(shell find generator-templates -type f)
> mkdir -p $(@D)
> docker run --rm -v ${PWD}:/local openapitools/openapi-generator-cli:latest generate -g go-experimental -c /local/generator-templates/go-experimental-generator-config.json -i /local/spec/intersight-openapi-v3.json -o /local/openapi --template-dir /local/generator-templates/go-experimental
> rm openapi/go.mod openapi/go.sum
> mv openapi/README.md openapi/operations.yaml
> ${GOPATH}/bin/goimports -l -w openapi

build/generator-postprocess: $(shell find generator-postprocess -name \*.go -type f)
> mkdir -p $(@D)
> go build -v -o "$@" $(GO_MODULE)/generator-postprocess 

cmd/cli.go: build/generator-postprocess openapi/operations.yaml generator-postprocess/cli.gotmpl
> build/generator-postprocess --operations openapi/operations.yaml --template generator-postprocess/cli.gotmpl --output "$@"
> go fmt "$@"

build/isctl: cmd/cli.go $(shell find cmd -name \*.go -type f)
> $(GO_BUILD_CMD) -o "$@" $(GO_MODULE)/cmd