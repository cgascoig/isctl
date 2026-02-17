[![User Guide](https://img.shields.io/badge/User%20Guide-Netlify-success)](https://isctl.netlify.app/) [![Build status](https://dev.azure.com/cgascoig/isctl/_apis/build/status/Full%20test?branchName=devel)](https://dev.azure.com/cgascoig/isctl/_build/latest?definitionId=2) [![Go Report](https://goreportcard.com/badge/github.com/cgascoig/isctl)](https://goreportcard.com/report/github.com/cgascoig/isctl)
# isctl - CLI for Cisco Intersight

The ultimate command-line companion for Cisco Intersight. Designed for DevOps professionals, bringing `kubectl`-style semantics to your infrastructure management, `isctl` allows you to manage your Intersight resources with ease and precision. 

## Features

* **OpenAPI-Driven:** Automatically generated from the Intersight OpenAPI v3 spec, ensuring capability with the latest features.
* **Cross-Platform:** Native binaries for macOS, Linux, and Windows.
* **Flexible Authentication:** Supports both API Key and OAuth 2.0 authentication with automatic [token caching](https://isctl.netlify.app/6-configuration/).
* **Developer Friendly:** Human, JSON, or YAML output with [JSONpath](https://goessner.net/articles/JsonPath/) and [Go Template](https://pkg.go.dev/text/template) support for powerful data extraction.
* **Shell Autocompletion:** Native support for [Bash, Zsh, Fish, and PowerShell](https://isctl.netlify.app/7-shell-autocompletion/).
* **Operational Reports:** Built-in reports for [HCL compliance, Contract status](https://isctl.netlify.app/8-reports/), and more.


## Installation

### MacOS

If you use [Homebrew](https://brew.sh), the easiest way to install `isctl` is:

```
brew install cgascoig/isctl/isctl
```

If you don't use Homebrew:

* Download the latest release from the [Releases](https://github.com/cgascoig/isctl/releases/latest) page. 
* Unzip and move the `isctl` binary somewhere that is on your path (e.g. `/usr/local/bin`). 


### Windows

The easiest way is using the [scoop.sh](https://scoop.sh/) installer:

```
scoop install https://github.com/cgascoig/isctl/raw/devel/isctl.json
```

Otherwise:

* Download the latest release from the [Releases](https://github.com/cgascoig/isctl/releases/latest) page. 
* Unzip and move the `isctl.exe` binary somewhere that is on your path. 

### Linux

If you use [Homebrew](https://brew.sh), the easiest way to install `isctl` is:

```
brew install cgascoig/isctl/isctl
```

If you don't use Homebrew:

* Download the latest release from the [Releases](https://github.com/cgascoig/isctl/releases/latest) page. 
* Extract the `.tar.gz` and move the `isctl` binary somewhere that is on your path (e.g. `/usr/local/bin`). 

## Documentation

The [Quick Start](#quick-start) below covers the basics but you should review the [Users Guide](https://isctl.netlify.app/) for complete documentation.

### Guides

* [Basic Queries](https://isctl.netlify.app/1-basic-queries/)
* [Advanced Queries](https://isctl.netlify.app/2-advanced-queries/)
* [Create, Update, Delete](https://isctl.netlify.app/3-create-update-delete/)
* [Bulk Operations](https://isctl.netlify.app/4-bulk-operations/)
* [Example Use Cases](https://isctl.netlify.app/5-example-use-cases/)
* [Configuration](https://isctl.netlify.app/6-configuration/)
* [Shell Autocompletion](https://isctl.netlify.app/7-shell-autocompletion/)
* [Reports](https://isctl.netlify.app/8-reports/)

## Quick Start

### Initial configuration

#### Credentials

`isctl` interacts with the Cisco Intersight REST API, so it needs **either** an API key or OAuth application credentials. 

##### API Key

1. Login to the Intersight GUI. 
2. Generate a new API Key (under Settings -> API Keys). Choose "API key for OpenAPI schema version 2" as the API Key Purpose. 
3. Save the key somewhere on your desktop and make a note of the key ID. 

##### OAuth

1. Login to the Intersight GUI.
2. Generate a new OAuth application (under Settings -> OAuth).
3. Make a note of the Client ID and Client Secret.

#### Configure `isctl`

Run `isctl configure` to configure it. Follow the prompts for your preferred authentication method.

```
Let's setup authentication. You can either configure intersight_api_key_id and intersight_secret_key (for API Key authentication), or you can configure intersight_client_id and intersight_client_secret (for OAuth authentication).
intersight_api_key_id is currently ''
Enter new intersight_api_key_id or press Enter to keep existing: 
...
```

### Querying Intersight

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

### Other Commands

#### Reports

Run operational reports directly from the CLI:

```
# Check Hardware Compatibility List (HCL) status
isctl report hcl

# Check device contract status
isctl report contract-status
```

#### Shell Autocompletion

Generate autocompletion scripts for your shell:

```
# Load bash completion
source <(isctl completion bash)

# For other shells (zsh, fish, powershell), see the documentation
isctl completion --help
``` 

