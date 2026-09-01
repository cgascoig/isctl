# Getting Started

## Installation

### MacOS

If you use [Homebrew](https://brew.sh), the easiest way to install `isctl` is:

```
brew install cgascoig/isctl/isctl
```

> Newer versions of Homebrew (6.0+) require third-party taps and formulae to be explicitly trusted before use. Installing with the fully-qualified name above automatically trusts just the `isctl` formula, so a fresh install needs no further action.
>
> **Upgrading:** If you already have `isctl` installed from this tap, upgrading with `brew upgrade isctl` (or `brew install isctl`) may fail with:
>
> ```
> Error: Refusing to load formula cgascoig/isctl/isctl from untrusted tap cgascoig/isctl.
> ```
>
> Fix this by trusting the formula (or the whole tap) once, then upgrade as usual:
>
> ```
> brew trust --formula cgascoig/isctl/isctl
> brew upgrade isctl
> ```
>
> Alternatively, `brew trust cgascoig/isctl` trusts the whole tap (all current and future formulae in it). If you prefer to install by the short name, trust the formula first:
>
> ```
> brew tap cgascoig/isctl
> brew trust --formula cgascoig/isctl/isctl
> brew install isctl
> ```

If you don't use Homebrew:

* Download the latest release from the [Releases](https://github.com/cgascoig/isctl/releases/latest) page. 
* Unzip and move the `isctl` binary somewhere that is on your path (e.g. `/usr/local/bin`). 

### Linux

If you use [Homebrew](https://brew.sh), the easiest way to install `isctl` is:

```
brew install cgascoig/isctl/isctl
```

> Newer versions of Homebrew (6.0+) require third-party taps and formulae to be explicitly trusted before use. Installing with the fully-qualified name above automatically trusts just the `isctl` formula, so a fresh install needs no further action.
>
> **Upgrading:** If you already have `isctl` installed from this tap, upgrading with `brew upgrade isctl` (or `brew install isctl`) may fail with:
>
> ```
> Error: Refusing to load formula cgascoig/isctl/isctl from untrusted tap cgascoig/isctl.
> ```
>
> Fix this by trusting the formula (or the whole tap) once, then upgrade as usual:
>
> ```
> brew trust --formula cgascoig/isctl/isctl
> brew upgrade isctl
> ```
>
> Alternatively, `brew trust cgascoig/isctl` trusts the whole tap (all current and future formulae in it). If you prefer to install by the short name, trust the formula first:
>
> ```
> brew tap cgascoig/isctl
> brew trust --formula cgascoig/isctl/isctl
> brew install isctl
> ```

If you don't use Homebrew:

* Download the latest release from the [Releases](https://github.com/cgascoig/isctl/releases/latest) page. 
* Extract the `.tar.gz` and move the `isctl` binary somewhere that is on your path (e.g. `/usr/local/bin`). 

### Windows

The easiest way is using the [scoop.sh](https://scoop.sh/) installer:

```
scoop install https://github.com/cgascoig/isctl/raw/devel/isctl.json
```

Otherwise:

* Download the latest release from the [Releases](https://github.com/cgascoig/isctl/releases/latest) page. 
* Unzip and move the `isctl.exe` binary somewhere that is on your path. 

## Obtain your API Key

`isctl` interacts with the Cisco Intersight REST API, so it needs an API key and key ID. 

1. Login to the Intersight GUI. 
2. Generate a new API Key (under Settings -> API Keys). Choose "API key for OpenAPI schema version 2" as the API Key Purpose. 
3. Save the key somewhere on your desktop and make a note of the key ID. 

## Configure `isctl`

Run `isctl configure` to configure it to use your API key. Follow the prompts for your key ID and path to the key file. 

```
keyID is currently ''
Enter new keyID or press Enter to keep existing: 123456789012345678901234/123456789012345678901234/123456789012345678901234
key filename is currently ''
Enter new key file name or press Enter to keep existing: /Users/chris/intersight-key.pem
2020/06/25 17:28:22 Writing config file
```

`isctl` is now configured and ready to use. 

## Execute your first query
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

<!-- 
# Welcome to MkDocs

For full documentation visit [mkdocs.org](https://www.mkdocs.org).

## Commands

* `mkdocs new [dir-name]` - Create a new project.
* `mkdocs serve` - Start the live-reloading docs server.
* `mkdocs build` - Build the documentation site.
* `mkdocs -h` - Print help message and exit.

## Project layout

    mkdocs.yml    # The configuration file.
    docs/
        index.md  # The documentation homepage.
        ...       # Other markdown pages, images and other files. -->
