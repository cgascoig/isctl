[![User Guide](https://img.shields.io/badge/User%20Guide-Netlify-success)](https://isctl.netlify.app/) [![Build status](https://dev.azure.com/cgascoig/isctl/_apis/build/status/Full%20test?branchName=devel)](https://dev.azure.com/cgascoig/isctl/_build/latest?definitionId=2) [![Go Report](https://goreportcard.com/badge/github.com/cgascoig/isctl)](https://goreportcard.com/report/github.com/cgascoig/isctl)
# isctl - CLI for Cisco Intersight
`isctl` is a `kubectl`-insipired CLI for the Cisco Intersight service. 

## Features

* Generated automatically from the Intersight OpenAPI v3 spec - it should be easy to keep up with new Intersight features. 
* Written in Go and distributed as MacOS, Linux and Windows binaries - should work anywhere
* Human, JSON or YAML output
* [JSONpath](https://goessner.net/articles/JsonPath/) support for extraction and transformation of the data returned from the API


# Installation

## MacOS

If you use [Homebrew](https://brew.sh), the easiest way to install `isctl` is:

```
brew install cgascoig/isctl/isctl
```

If you don't use Homebew:

* Download the latest release from the [Releases](https://github.com/cgascoig/isctl/releases/latest) page. 
* Unzip and move the `isctl` binary somewhere that is on your path (e.g. `/usr/local/bin`). 


## Windows

The easiest way is using the [scoop.sh](https://scoop.sh/) installer:

```
scoop install https://github.com/cgascoig/isctl/raw/devel/isctl.json
```

Otherwise:

* Download the latest release from the [Releases](https://github.com/cgascoig/isctl/releases/latest) page. 
* Unzip and move the `isctl.exe` binary somewhere that is on your path. 

## Linux

* Download the latest release from the [Releases](https://github.com/cgascoig/isctl/releases/latest) page. 
* Extract the `.tar.gz` and move the `isctl` binary somewhere that is on your path (e.g. `/usr/local/bin`). 

# Documentation

The [Quick Start](#quick-start) below covers the basics but you should review the [Users Guide](https://isctl.netlify.app/) for complete documentation.

# Quick Start

## Initial configuration

### Obtain your API Key

`isctl` interacts with the Cisco Intersight REST API, so it needs an API key. 

1. Login to the Intersight GUI. 
2. Generate a new API Key (under Settings -> API Keys). Choose "API key for OpenAPI schema version 2" as the API Key Purpose. 
3. Save the key somewhere on your desktop and make a note of the key ID. 

### Configure `isctl`

Run `isctl configure` to configure it to use your API key. Follow the prompts for your key ID and path to the key file. (Note: the path to the key file should be an absolute path and doesn't support shell metacharacters such as "`~`")

```
keyID is currently ''
Enter new keyID or press Enter to keep existing: 123456789012345678901234/123456789012345678901234/123456789012345678901234
key filename is currently ''
Enter new key file name or press Enter to keep existing: /Users/chris/intersight-key.pem
2020/06/25 17:28:22 Writing config file
```

## Querying Intersight

To retrieve resources from Intersight, use the `isctl get ...` commands. 

For example, to query all objects of a certain type:

```
isctl get ntp policy
```
Output:
```
------------  -----------------------------  ------------------------  -----------------------------
    Enabled                           Moid                      Name                     NtpServers
------------  -----------------------------  ------------------------  -----------------------------
      False       5ecb069c6275722d3143e668           test-ntp-policy

       True       5ee1aa076275722d3122a944          cg-tf-ntp-test-1       10.10.10.10, 10.10.10.12
------------  -----------------------------  ------------------------  -----------------------------
```

For detailed documentation, see the [isctl Users Guide](https://isctl.netlify.app/). 

# Development

## How is this built?

There are a number of steps in the build process as the majority of the code is generated automatically:

1. [OpenAPITools OpenAPI Generator](https://github.com/OpenAPITools/openapi-generator) is used to generate the Go models and API client from the Intersight OpenAPI v3 spec. The OpenAPI Generator Go templates are customised to also generate an `operations.yaml` file. 
2. The Go program in `generator-postprocess` is used to parse the `operations.yaml` file and generate the command line structure in `cmd/cli.go`. This postprocess step also cleans up the command naming and creates the help text. 
3. Finally, the `isctl` command is built from the generated `cmd/cli.go` along with some supporting code. 

## Creating a release
The `Makefile` is used for most of the post processing and building `isctl` for the local system but [GoReleaser](https://goreleaser.com) is used for creating releases on GitHub. 

```
# Create tag and push to GitHub
git tag vX.Y.Z
git push origin --tags

# Create the release
goreleaser --rm-dist
```
