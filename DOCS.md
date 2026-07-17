# guh CLI — Command Reference

## `guh`

Root command. Prints a welcome message when run without a subcommand.

```
guh
```

---

## `browse`

Open the repository's remote origin in the default browser.

```
guh browse
```

Equivalent to opening the repo URL (without `.git` suffix) in a browser.

### `issues`

Open the issues page.

```
guh browse issues
```

### `pr`

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

## `commit`

Add files and commit changes. Shows `git status --short` before committing.

```
guh commit                                       # show edited files only
guh commit "." "message"                         # git add . && commit
guh commit '["file1", "file2"]' "message"        # add specific files && commit
guh commit '["file1"]' "message" "description"   # commit with message + description
```

**Arguments**

| # | Name | Required | Description |
|---|------|----------|-------------|
| 1 | files | yes* | JSON array of files, or `"."` to add everything |
| 2 | message | yes* | Commit message |
| 3 | description | no | Optional commit description (uses `-m` twice) |

\*When no arguments are given, shows `git status --short` instead.

---

## `create`

### `branch`

Create and switch to a new branch.

```
guh create branch <branch>
```

Equivalent to `git switch -c <branch>`.

### `account`

Save a GitHub account with username, email, classic token, and platform.

```
guh create account <username> <email> <classicToken> <platform>
```

Credentials are stored in the user config directory under `dev.pages.codedave.guh/accounts.json`. All four values are required and must be non-empty. `platform` is the Git provider domain (e.g. `github.com`).

### `repo`

Link or view the repository remote origin.

```
guh create repo <url>     # set origin URL
guh create repo what      # show current origin URL
```

---

## `diff`

Show working tree changes.

```
guh diff
```

Equivalent to `git diff`.

---

## `init`

Initialize a new Git repository in the current directory.

```
guh init
```

Equivalent to `git init`.

---

## `pull`

Fetch and pull upstream changes.

```
guh pull
```

Runs `git fetch` followed by `git pull`.

---

## `push`

Push local commits.

```
guh push
```

Equivalent to `git push`.

---

## `stash`

Stash current working directory changes.

```
guh stash
```

Equivalent to `git stash`.

---

## `switch`

### `branch`

Switch to an existing branch.

```
guh switch branch <name>
```

Equivalent to `git switch <name>`.

### `account`

Switch to a saved GitHub account (updates global git config and credentials).

```
guh switch account <username>
```

Updates `user.name`, `user.email` in global git config, and overwrites `~/.git-credentials` with the stored token. The account must have been created with `guh create account` first.

---

## `undo`

Undo local commits using `git reset`.

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

## `logs`
Show commit logs.

```
guh logs [commits]
```

**Arguments**

| # | Name | Required | Description |
|---|------|----------|-------------|
| 1 | commits | no | Number of commits to show (default: all) |

Examples:
```
guh logs 5
```

Equivalent to `git log`.
