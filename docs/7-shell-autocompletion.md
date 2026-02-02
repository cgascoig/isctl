# Shell Autocompletion

`isctl` supports autocompletion for Bash, Zsh, Fish, and PowerShell. This feature can save you a lot of typing and help you explore available commands and flags.

You can generate the autocompletion script for your shell using the `isctl completion <SHELL>` command.

## Bash

This script depends on the `bash-completion` package. If it is not installed already, you can install it via your OS's package manager.

To load completions in your current shell session:

```bash
source <(isctl completion bash)
```

To load completions for every new session, execute once:

**Linux:**

```bash
isctl completion bash > /etc/bash_completion.d/isctl
```

**macOS:**

```bash
isctl completion bash > /usr/local/etc/bash_completion.d/isctl
```

You will need to start a new shell for this setup to take effect.

## Zsh

If shell completion is not already enabled in your environment you will need to enable it. You can execute the following once:

```zsh
echo "autoload -U compinit; compinit" >> ~/.zshrc
```

To load completions for every new session, execute once:

**Linux:**

```zsh
isctl completion zsh > "${fpath[1]}/_isctl"
```

**macOS:**

```zsh
isctl completion zsh > /usr/local/share/zsh/site-functions/_isctl
```

You will need to start a new shell for this setup to take effect.

## Fish

To load completions in your current shell session:

```fish
isctl completion fish | source
```

To load completions for every new session, execute once:

```fish
isctl completion fish > ~/.config/fish/completions/isctl.fish
```

You will need to start a new shell for this setup to take effect.

## PowerShell

To load completions in your current shell session:

```powershell
isctl completion powershell | Out-String | Invoke-Expression
```

To load completions for every new session, add the output of the above command to your PowerShell profile.
