# guh CLI — Command Reference

## Table of Contents

- [`guh`](#guh)
- [`account`](#account)
  - [`account save`](#account-save)
  - [`account switch`](#account-switch)
  - [`account edit`](#account-edit)
- [`branch`](#branch)
  - [`branch create`](#branch-create)
  - [`branch switch`](#branch-switch)
  - [`branch rename`](#branch-rename)
  - [`branch delete`](#branch-delete)
- [`browse`](#browse)
  - [`browse issues`](#browse-issues)
  - [`browse pr`](#browse-pr)
- [`cli`](#cli)
  - [`cli install`](#cli-install)
  - [`cli remove`](#cli-remove)
  - [`cli update`](#cli-update)
  - [`cli version`](#cli-version)
- [`clone`](#clone)
- [`commit`](#commit)
- [`diff`](#diff)
- [`init`](#init)
- [`link`](#link)
  - [`link add`](#link-add)
  - [`link remove`](#link-remove)
- [`logs`](#logs)
- [`pull`](#pull)
- [`push`](#push)
- [`stash`](#stash)
- [`undo`](#undo)

---

## `guh`

Root command. Prints a welcome message when run without a subcommand.

```
guh
> Welcome to guh! Use --help to see available commands.
```

---

## `account`

### `account save`

Save a new account with username, email, classic token, and platform.

```
guh account save <username> <email> <classicToken> <platform>
```

All four values are required and must be non-empty. `platform` is the Git provider domain (e.g. `github.com`). Credentials are stored in the user config directory under `dev.pages.codedave.guh/accounts.json`.

### `account switch`

Switch to a saved account (updates global git config and credentials).

```
guh account switch <username>
```

Sets `credential.helper` to `store`, updates `user.name` and `user.email` in global git config, and overwrites `~/.git-credentials` with the stored token. The account must have been saved with `guh account save` first.

### `account edit`

Edit accounts.json via the supported code editor

```
guh account edit
```

Opens the supported code editor; "nano", "code", "zed", "zeditor", "vi", "vim", "emacs", "notepad.exe"

---

## `branch`

No args given, lists all branches.

```
guh branch
```

Equivalent to `git branch`.

### `branch create`

Create and switch to a new branch.

```
guh branch create <name>
```

Equivalent to `git switch -c <name>`.

### `branch switch`

Switch to an existing branch.

```
guh branch switch <name>
```

Equivalent to `git switch <name>`.

### `branch rename`

Rename the current branch.

```
guh branch rename <name>
```

Equivalent to `git branch -m <name>`.

### `branch delete`

Delete the selected branch.

```
guh branch delete <branch>
```

Equivalent to `git branch -D <branch>`.

---

## `browse`

Open the repository's remote origin in the default browser.

```
guh browse
```

Equivalent to opening the repo URL (without `.git` suffix) in a browser.

### `browse issues`

Open the issues page.

```
guh browse issues
```

### `browse pr`

Open the pull requests page.

```
guh browse pr
```

Optionally, add `print` to print the URL instead of opening it.

```
guh browse issues print
guh browse pr print
guh browse print
```

---

## `cli`

### `cli install`
Installs `guh` binary to your system directly with env variable.

```
guh cli install
```

### `cli remove`
Deletes `guh` binary in your system with the env variable.

```
guh cli remove
```

> [!WARNING]
> Windows' permission issue may cause it to not delete. We have provided a command for you to delete through the terminal for you to run. We can't do anything about it because of *cough* windows

### `cli update`
Updates `guh` binary in your system.

```
guh cli update
```

> [!WARNING]
> Windows' permission issue may cause it to not update. We have provided a command for you to update through the terminal for you to run. We can't do anything about it because of *cough* windows

### `cli version`
Shows your `guh` binary hash[:7] version through git's commit.

```
guh cli version
```

---

## `clone`

Clones a github repository and spawns a nested shell inside the directory, automatically putting you in there.

```
guh clone <url>
```

---

## `commit`

Add files and commit changes. Shows `git status --short` before committing. Optionally pushes after commit.

```
guh commit                                                      # show edited files only
guh commit "." "message"                                        # git add . && commit
guh commit '["file1", "file2"]' "message"                       # add specific files && commit
guh commit '["file1"]' "message" "description"                  # commit with message + description
guh commit '["file1"]' "message" push                           # commit + push to origin
guh commit '["file1"]' "message" "description" push             # commit + push to origin
guh commit '["file1"]' "message" push "upstream"                # commit + push to named remote
guh commit '["file1"]' "message" "description" push "all"       # commit + push to all remotes
```

**Arguments**

| # | Name | Required | Description |
|---|------|----------|-------------|
| 1 | files | yes* | JSON array of files, or `"."` to add everything |
| 2 | message | yes* | Commit message |
| 3 | description | no | Optional commit description (uses `-m` twice) |
| 4 | push | no | Literal `"push"` to trigger a push after commit |
| 5 | remote | no | Remote name (default `"origin"`; use `"all"` to push to all remotes) |

\*When no arguments are given, shows `git status --short` instead.

---

## `diff`

Show working tree changes.

```
guh diff            # show all working tree changes
guh diff <file>     # show changes for a specific file
```

Equivalent to `git diff [file]`.

---

## `init`

Initialize a new Git repository in the current directory.

```
guh init
```

Equivalent to `git init`.

---

## `link`

Show remote URLs or add/remove remotes.

```
guh link                        # show all remotes (git remote -v)
guh link add <name> <url>       # add (or replace) a remote
guh link remove <name>          # remove a remote
```

### `link add`

Set a remote URL. Removes any existing remote with the same name first, then adds it. Requires exactly 2 arguments: remote name and URL.

```
guh link add <name> <url>
```

### `link remove`

Remove a remote by name. Requires exactly 1 argument: the remote name.

```
guh link remove <name>
```

---

## `logs`

Show commit logs.

```
guh logs [commits]
```

**Arguments**

| # | Name | Required | Description |
|---|------|----------|-------------|
| 1 | commits | no | Number of commits to show (default: all commits) |

Examples:

```
guh logs        # git log (full history)
guh logs 5      # git log -5 (last 5 commits)
```

---

## `pull`

Fetch and pull upstream changes.

```
guh pull
```

Runs `git fetch` followed by `git pull`.

---

## `push`

Push the current branch to a remote. Uses `-u` to set upstream tracking automatically.

```
guh push                    # git push -u origin <current-branch>
guh push <remote>           # git push -u <remote> <current-branch>
guh push all                # push to all remotes
```

When `all` is given, pushes to every remote returned by `git remote`. The first remote uses `-u` and the current branch name to set upstream tracking; subsequent remotes are pushed without the `-u` flag.

**Arguments**

| # | Name | Required | Description |
|---|------|----------|-------------|
| 1 | remote | no | Remote name, or `"all"` to push to every remote (default: `"origin"`) |

---

## `stash`

Stash current working directory changes.

```
guh stash
```

Equivalent to `git stash`.

---

## `undo`

Undo local commits using `git reset`. Flavour defaults to `soft` for unrecognized values.

```
guh undo <commits> [flavour]
```

**Arguments**

| # | Name | Required | Description |
|---|------|----------|-------------|
| 1 | commits | yes | Number of commits to go back (e.g. `1`) |
| 2 | flavour | no | Reset mode: `hard`, `mixed` (default: `soft`) |

Examples:

```
guh undo 1             # git reset --soft HEAD~1
guh undo 2 hard        # git reset --hard HEAD~2
guh undo 1 mixed       # git reset --mixed HEAD~1
```
