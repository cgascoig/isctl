![Build](https://github.com/cgascoig/isctl/workflows/Build/badge.svg)

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

* Download the latest release from the [Releases](releases/) page. 
* Unzip and move the `isctl` binary somewhere that is on your path (e.g. `/usr/local/bin`). 


## Windows

* Download the latest release from the [Releases](releases/) page. 
* Unzip and move the `isctl.exe` binary somewhere that is on your path. 

## Linux

* Download the latest release from the [Releases](releases/) page. 
* Extract the `.tar.gz` and move the `isctl` binary somewhere that is on your path (e.g. `/usr/local/bin`). 


# Usage

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

To retrieve resources from Intersight, use the `isctl get ...` commands. For example:
```
isctl get ntp policy
```
>```
------------  -----------------------------  ------------------------  -----------------------------
    Enabled                           Moid                      Name                     NtpServers
------------  -----------------------------  ------------------------  -----------------------------
       True       5a9a0337736d74787a6093db                   NTP_ESL              ntp.esl.cisco.com

       True       5b21faf1666c663463a8341f                  CiscoNTP              ntp.esl.cisco.com

       True       5cc6f2426275722d30fbd461                    BA-NTP              ntp.esl.cisco.com

      False       5df87ad06275722d31e0dfe5       se-cimc-6ntp-policy

      False       5ecb069c6275722d3143e668           test-ntp-policy

       True       5ee1aa076275722d3122a944          cg-tf-ntp-test-1       10.10.10.10, 10.10.10.12
------------  -----------------------------  ------------------------  -----------------------------
```

# Developing

## Creating a release
```
# Create tag and push to GitHub
git tag vX.Y.Z
git push origin --tags

# Create the release (using a specific go installation built from source until Go 1.15 is released)
PATH=/Users/cgascoig/dev/gosrc/bin:$PATH goreleaser --rm-dist
```