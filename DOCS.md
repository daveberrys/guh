# guh CLI — Command Reference

## `guh`

Root command. Prints a welcome message when run without a subcommand.

```
guh
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

Save a GitHub account with username, email, and classic token.

```
guh create account <username> <email> <classicToken>
```

Credentials are stored in the user config directory under `dev.pages.codedave.guh/accounts.json`. All three values are required and must be non-empty.

---

## `diff`

Show working tree changes.

```
guh diff
```

Equivalent to `git diff`.

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
