# CLAUDE.md — boynoiz/dst

Implementation rules for Claude Code sessions working on this repository.

---

## Project Overview

This is a maintained fork of [`dave/dst`](https://github.com/dave/dst) (Decorated Syntax Tree),
renamed to `github.com/boynoiz/dst`. The fork exists because upstream is unmaintained and
incompatible with modern Go versions (Go 1.21+).

**Module path:** `github.com/boynoiz/dst`
**Original upstream:** `github.com/dave/dst`
**Local upstream clone (for cherry-picks):** `/tmp/dst` (branch `pr-86`)

---

## What This Library Does

`dst` is an alternative to `go/ast` that preserves comments and whitespace decorations
when transforming Go source code. Unlike `go/ast`, round-tripping source through `dst`
does not lose comments attached to nodes.

**Core packages:**
- `github.com/boynoiz/dst` — DST node types (parallel to `go/ast`)
- `github.com/boynoiz/dst/decorator` — converts `go/ast` ↔ `dst`, restores source
- `github.com/boynoiz/dst/dstutil` — `dst`-aware tree walker (parallel to `go/ast/astutil`)
- `github.com/boynoiz/dst/gendst` — code generator for generated files

---

## Fork Changes vs Upstream

All changes are logged in `CHANGELOG.md` (generated via `git-cliff`).

Current fork additions over `dave/dst@master`:

| Commit | Change |
|--------|--------|
| `3a30d7d` | Module rename: `dave/dst` → `boynoiz/dst` (60+ files) |
| `8096f70` | Fix struct literal field order, missing `errors` imports |
| `f01cd06` | Fix: only update space type if new value is larger (PR #86) |
| `26becc9` | Feat: add `GoVersion` field to `dst.File` (Go 1.21+) |
| `3a21510` | Feat: support Go 1.24 generic type aliases (`TypeSpec.Assign` ordering) |
| *(pending)* | Modernize: `ioutil` → `os`/`io`, octal literals, `os.DirEntry` conversion |
| *(pending)* | Test: skip known upstream bugs with documented justification |

---

## Generated Files

Several files are **generated** by `gendst`. Do NOT hand-edit these:

```
clone-generated.go
decorations-types-generated.go
decorator/decorator-fragment-generated.go
decorator/decorator-node-generated.go
decorator/restorer-generated.go
dstutil/decorations-generated.go
```

To regenerate: `cd gendst && go run . && cd ..`

The source of truth for generated files is `gendst/data/data.go` and `gendst/data/positions.go`.

---

## Known Upstream Bugs (Do Not Fix Without Discussion)

These are pre-existing bugs in `dave/dst`, not regressions from the fork:

### 1. Blank import alias dropped by restorer

**Symptom:** `TestLoadStdLibAll` fails on files containing `_ "unsafe"` — the blank
identifier alias is dropped during restore, producing `"unsafe"` instead of `_ "unsafe"`.

**Affected files (Go 1.21+ stdlib):**
- `runtime/rand.go`
- `reflect/badlinkname.go`
- `internal/godebug/godebug.go`
- `crypto/rand/rand.go`

**Status:** Skipped with `t.Skip(...)` — not a fork regression.

### 2. `crypto/internal/fips140test` has no syntax

**Symptom:** `TestLoadStdLibAll` reports "Package crypto/internal/fips140test has no syntax".

**Cause:** This internal test package only exists under specific build tags; the test's
package loader finds it but there are no parseable Go files in the current build context.

**Status:** Skipped with `t.Skip(...)` — not a fork regression.

### 3. `TestPositions` — `SliceExpr(2)` fixture

**Symptom:** `decorator_pos_test.go:81: missed *dst.SliceExpr 2`

**Cause:** The fixture in `gendst/data/positions.go` requires an expression in a
position that `go vet` flags as unreachable code when wrapped naively.

**Status:** Skipped or fixed per session — see commit history.

---

## Build & QA

Always use the Makefile. Never run raw `go test` for a full check.

```bash
# Full QA: fmt, fix, lint, vet, vuln, test
make qa

# Run tests only
make test

# Run a specific test
go test ./decorator/ -run TestPositions -v

# Regenerate generated files
cd gendst && go run . && cd ..

# Update changelog
git-cliff -o CHANGELOG.md
```

---

## Definition of Done

Every session must pass all of these before committing:

```
go fix ./...          ✓
go build ./...        ✓
go vet ./...          ✓
golangci-lint run     ✓
make test             ✓ (known upstream skips are acceptable)
```

No exceptions. Do not commit if any of the above fails.

---

## Commit Style

Use Conventional Commits:

```
feat:     new capability added to the fork
fix:      bug fix (distinguish fork fix vs upstream cherry-pick)
test:     test additions or skip documentation
chore:    maintenance (module rename, dep updates)
refactor: code restructuring without behavior change
docs:     documentation only
```

Examples:
```
fix(decorator): only update space type if new value is larger
feat: add GoVersion field to dst.File (Go 1.21+)
test: skip known upstream bugs in TestPositions and TestLoadStdLibAll
chore: rename module path from dave/dst to boynoiz/dst
```

---

## Code Rules

### Output format
Always wrap all code in markdown code blocks. Claude Code tends to omit these — always include them.

### Import ordering
Use `gci` grouping:
1. Standard library
2. Third-party
3. Internal (`github.com/boynoiz/dst/...`)

### Do not use `ioutil`
All `ioutil.*` calls have been replaced with `os.*` / `io.*`. If you add new code,
use the modern equivalents:

| Old | New |
|-----|-----|
| `ioutil.ReadFile` | `os.ReadFile` |
| `ioutil.WriteFile` | `os.WriteFile` |
| `ioutil.ReadAll` | `io.ReadAll` |
| `ioutil.NopCloser` | `io.NopCloser` |
| `ioutil.ReadDir` | `os.ReadDir` (returns `[]os.DirEntry`, not `[]os.FileInfo`) |

### `os.ReadDir` return type

`os.ReadDir` returns `[]os.DirEntry`, **not** `[]os.FileInfo`. If you need `[]os.FileInfo`,
convert explicitly:

```go
entries, err := os.ReadDir(path)
if err != nil {
return nil, err
}
infos := make([]os.FileInfo, len(entries))
for i, e := range entries {
infos[i], err = e.Info()
if err != nil {
return nil, err
}
}
```

### Octal literals
Use `0o755` not `0755`. `go fix` handles this, but write it correctly in new code.

---

## Cherry-picking from Upstream

The local upstream clone at `/tmp/dst` has the `pr-86` branch fetched.
It is registered as a remote:

```bash
git remote add tmp-pr86 /tmp/dst   # if not already added
git fetch tmp-pr86
git cherry-pick <commit-sha>
```

Always verify cherry-picks compile and pass `make qa` before committing.

---

## Tooling Preferences

Use these tools — do not suggest traditional alternatives:

| Task | Tool |
|------|------|
| Search | `rg` (ripgrep), not `grep` |
| Find files | `fd`, not `find` |
| String replace | `sd`, not `sed` |
| View files | `bat`, not `cat` |
| Database | `usql`, not `psql` |

---

## Session Workflow

1. Read this file first
2. Read relevant source files before making changes
3. Make changes
4. Run `make qa` — fix everything until clean
5. Commit with a meaningful Conventional Commit message
6. Update `CHANGELOG.md`: `git-cliff -o CHANGELOG.md`
