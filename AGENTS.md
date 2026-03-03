# AGENTS

Guidance for agentic coding tools working in this repository. Covers build,
lint, test commands, and code style conventions. Derived from
`.github/copilot-instructions.md` and `.github/agents.md`.

## Project Snapshot

- **Language:** Go (module `github.com/gopinath-langote/1build`)
- **Product:** CLI tool that runs project-local command aliases via `1build.yaml`
- **CLI framework:** Cobra (command routing), Viper (flag/config binding)
- **Config format:** YAML (`1build.yaml` in project root)
- **Tests:** End-to-end — builds a temp binary, runs fixtures as subprocesses
- **Go version:** Module targets go 1.12; CI uses Go 1.22 (avoid features newer than 1.12)

---

## Build Commands

```bash
go build          # build binary in current directory
go build -v .     # verbose build (matches CI)
```

The project self-hosts its own shortcuts via `1build.yaml`:

```bash
1build build
```

---

## Lint Commands

```bash
# Install golint (one-time)
go get -u golang.org/x/lint/golint

# Run linter
go list ./... | xargs -L1 golint
```

CI additionally runs `golangci-lint --enable-all --exclude-use-default=false`
via `reviewdog` on pull requests.

```bash
1build lint   # project shortcut
```

---

## Test Commands

### Full test suite

```bash
go test -v -cover github.com/gopinath-langote/1build/testing -run .

# Cache-busting variant (matches CI exactly)
go test -v -cover github.com/gopinath-langote/1build/testing -run . GOCACHE=off
```

### Run a single top-level test

```bash
go test -v github.com/gopinath-langote/1build/testing -run TestAll
```

### Run a single subtest (fixture-based naming)

Subtests follow the pattern `".:.<feature>.:.<name>"`. Examples:

```bash
go test -v github.com/gopinath-langote/1build/testing \
  -run "TestAll/\.:\.exec\.:\.shouldExecuteAvailableCommand"

go test -v github.com/gopinath-langote/1build/testing \
  -run "TestAll/\.:\.flag\.:\.shouldPrintVersion"
```

### Project shortcut

```bash
1build test
```

### How tests work

- `testing/cli_test.go:TestMain` builds a real `1build` binary into a temp dir.
- `TestAll` iterates `fixtures.GetFixtures()` and runs each as a subprocess.
- Each `Test` struct carries: `Feature`, `Name`, `CmdArgs`, `Setup`, `Assertion`, `Teardown`.
- Add new test cases in `testing/fixtures/` following existing fixture files.

---

## Directory Structure

```
1build.go                  # main() — calls cmd.Execute()
1build.yaml                # Project's own 1build config (self-hosted)
cmd/
  root.go                  # Root Cobra command, flag definitions
  config/
    io.go                  # ReadFile, WriteConfigFile, DeleteConfigFile
    parse.go               # OneBuildConfiguration struct, LoadOneBuildConfiguration
  exec/exec.go             # ExecutePlan() — main execution orchestrator
  models/
    onebuild-execution-plan.go  # CommandContext, OneBuildExecutionPlan
  utils/
    utils.go               # ExitWithCode, SliceIndex, LongestString
    printer.go             # CPrint, CPrintln, color constants
  del/ initialize/ list/ set/ unset/   # One subcommand per directory
testing/
  cli_test.go              # TestMain + TestAll
  fixtures/                # All test fixture definitions (one file per feature)
  utils/test_utils.go      # Helpers: CreateConfigFile, AssertContains, temp dirs
  def/definitions.go       # Shared constants (ConfigFileName)
```

---

## Code Style Guidelines

### Formatting

- Run `gofmt` on all Go files before committing.
- Markdown line length limit: 190 characters (enforced by `.remarkrc.yaml`).

### Imports

Group in this order, separated by blank lines:

```go
import (
    "fmt"        // 1. standard library
    "os"

    "github.com/spf13/cobra"   // 2. third-party

    "github.com/gopinath-langote/1build/cmd/utils"  // 3. local module
)
```

Use import aliases only to resolve name collisions (e.g., `configuration "..."`,
`utils2 "..."`). Keep alias names short and obvious.

### Naming Conventions

| Element | Convention | Examples |
|---|---|---|
| Exported types/funcs | PascalCase | `OneBuildConfiguration`, `ExecutePlan` |
| Unexported funcs/vars | camelCase | `bashCommand`, `executeAndStopIfFailed` |
| Constants (enum-like) | ALL\_CAPS | `CYAN`, `RED`, `MaxOutputWidth` |
| Package names | lowercase single word | `config`, `exec`, `utils`, `models` |
| File names | `kebab-case.go` for multi-word | `onebuild-execution-plan.go` |
| Test case functions | camelCase descriptive | `shouldExecuteAvailableCommand` |

CLI subcommand tokens are lowercase (`set`, `unset`, `exec`, `list`).

### Error Handling

Follow this pattern consistently — no panics:

```go
// Config/parse layer: return the error, caller decides
configuration, err := config.LoadOneBuildConfiguration()
if err != nil {
    fmt.Println(err)
    return
}

// CLI layer: print then exit with code
if err != nil {
    fmt.Println(err)
    utils.ExitError()   // exit code 1
}

// Command execution failure: styled print + subprocess exit code
utils.CPrintln(text, utils.Style{Color: utils.RED})
utils.ExitWithCode(exitCode)
```

- Use `errors.New("...")` to construct errors in the config layer.
- Exit code 127 for "command not found" (POSIX convention).
- Explicitly discard ignored errors with `_ =`, never silently.

### Types and Data Flow

- Use explicit named struct types for all models (`CommandContext`,
  `OneBuildExecutionPlan`, `CommandDefinition`).
- YAML struct tags use `omitempty` for optional fields.
- Implement `yaml.Unmarshaler` on `CommandDefinition` to support both
  inline (`build: "go build"`) and nested YAML forms.
- Avoid defining new Go interfaces unless necessary; prefer concrete types.
- Use pointer receivers on methods that read or mutate struct state.

### Output and Logging

- Use `utils.CPrint` / `utils.CPrintln` for all styled terminal output.
- Respect the `--quiet` (`-q`) flag: suppress banners and command output,
  show only SUCCESS/FAILURE when quiet.
- Show phase banners and command labels during execution (see `exec/exec.go`).

### CLI Command Pattern

Each subcommand is a package-level Cobra var:

```go
var Cmd = &cobra.Command{
    Use:    "set",
    Short:  "...",
    PreRun: func(cmd *cobra.Command, args []string) { ... },
    Run:    func(cmd *cobra.Command, args []string) { ... },
}
```

Register all subcommands in `cmd/root.go:init()`.

### Comments and Documentation

- All exported symbols must have a GoDoc comment starting with the symbol name.
- Explain "why" in comments, not just "what".
- Use single-line comments for constants; block comments for longer explanations.

---

## Dependencies

Key dependencies (do not introduce replacements without discussion):

| Package | Purpose |
|---|---|
| `github.com/spf13/cobra` | CLI command routing |
| `github.com/spf13/viper` | Global flag/config binding |
| `github.com/codeskyblue/go-sh` | Shell session execution |
| `gopkg.in/yaml.v3` | YAML parsing / marshalling |
| `github.com/logrusorgru/aurora` | Colored terminal output |
| `github.com/stretchr/testify` | Test assertions |

Keep `go.mod` and `go.sum` up to date after any dependency change.

---

## Repo Hygiene

- Do not commit build artifacts (`dist/`, binary `1build` at root).
- Keep `1build.yaml` committed (explicitly unignored in `.gitignore`).
- CI runs on `ubuntu-latest` and `macos-latest` with Go 1.22.
- If adding Markdown files, comply with `.remarkrc.yaml` (max 190-char lines).
- Changelog excludes commits prefixed with `docs:` or `test:` (goreleaser config).

---

## Key Entry Points

| File | Purpose |
|---|---|
| `1build.go` | `main()` — calls `cmd.Execute()` |
| `cmd/root.go` | Root Cobra command, flag registration |
| `cmd/exec/exec.go` | `ExecutePlan()` — full execution orchestrator |
| `cmd/config/parse.go` | `LoadOneBuildConfiguration()`, struct definitions |
| `cmd/config/io.go` | File read/write/delete helpers |
| `testing/cli_test.go` | `TestMain`, `TestAll` |
| `testing/fixtures/fixures.go` | `GetFixtures()` aggregator, `Test` struct |
