# 1build - Copilot Instructions

## Project Overview

**1build** is a Go-based CLI tool that provides project-local command aliases for automation. It allows developers to configure simple, memorable commands for project-specific build tools and tasks via a `1build.yaml` configuration file.

## Build & Test Commands

### Build
```bash
go build
```
Builds the binary in the current directory.

### Run Tests
```bash
# Full test suite
go test -v -cover github.com/gopinath-langote/1build/testing -run .

# With coverage output captured
go test -v -cover github.com/gopinath-langote/1build/testing -run . GOCACHE=off
```

### Lint
```bash
go list ./... | xargs -L1 golint
```
Requires `golint` to be installed first:
```bash
go get -u golang.org/x/lint/golint
```

### Using 1build Commands (for this project)
The project itself uses 1build:
```bash
1build test        # Run tests
1build lint        # Run linter
1build build       # Build the binary
```

## Architecture

### Command Structure (Cobra-based)
- **Root Command** (`cmd/root.go`): Entry point that orchestrates all subcommands. If no args provided, shows command list. Otherwise executes commands via `exec.ExecutePlan()`.
- **Subcommands** in `cmd/`:
  - `exec/` - Executes configured commands with before/after hooks
  - `list/` - Lists available commands from config
  - `initialize/` - Creates new `1build.yaml` file
  - `set/` - Adds or updates command definitions
  - `unset/` - Removes command definitions
  - `del/` - Deletes entire command configurations (force flag available)
  - `config/` - Loads and parses `1build.yaml`
  - `models/` - Data structures for commands and configuration
  - `utils/` - Helper functions (styling, exit codes)

### Configuration System
- **File**: `1build.yaml` (path can be overridden with `-f` flag)
- **Key Fields**:
  - `project` - Project name
  - `commands` - List of command definitions (name → shell command mapping)
  - `before` - Optional hook executed before all commands
  - `after` - Optional hook executed after all commands
  - `beforeAll` / `afterAll` - Project-level hooks (separate from per-command hooks)
- **Loading**: `config.LoadOneBuildConfiguration()` validates existence and parses YAML

### Execution Flow
1. `ExecutePlan()` in `exec/exec.go` orchestrates command execution
2. Global `beforeAll` hook runs if defined
3. For each command argument:
   - Find command definition
   - Execute `before` hook if defined
   - Execute main command (captures output, handles exit codes)
   - Execute `after` hook if defined
4. Stops on first failure (unless `--quiet` flag suppresses output)

### Testing
- **Location**: `testing/` package
- **Test Framework**: Standard Go testing with fixtures-based approach
- **Fixtures**: `testing/fixtures/` contains test cases with setup, command args, assertions, and teardown
- **Execution**: Tests build a temporary `1build` binary and run it against test directories
- **Coverage**: Tests exercise CLI behavior end-to-end by simulating user commands

## Key Conventions

### Command-Line Flags
- `-q` / `--quiet`: Suppress command output, only show SUCCESS/FAILURE
- `-f` / `--file`: Custom path to config file (default: `1build.yaml`)
- `--before`, `--after`: Hooks for individual commands (via `set` command)
- `--beforeAll`, `--afterAll`: Project-level hooks

### Error Handling
- Non-zero exit codes from commands stop execution unless `--quiet` is used
- `utils.ExitError()` and `utils.ExitWithCode()` standardize exit behavior
- Missing commands show config and exit with code 127

### Module & Dependency Management
- Uses Go modules (`go.mod`)
- Key dependencies:
  - `github.com/spf13/cobra` - CLI framework
  - `github.com/spf13/viper` - Config parsing
  - `github.com/codeskyblue/go-sh` - Shell command execution
  - `gopkg.in/yaml.v3` - YAML parsing
  - `github.com/logrusorgru/aurora` - Colored terminal output
  - `github.com/stretchr/testify` - Testing assertions

### Release Process
- Uses `goreleaser` for multi-platform builds (Linux, macOS, Windows)
- Triggered by git tags (semantic versioning): `git tag v1.x.x && git push origin --tags`
- Runs `goreleaser` command to build and create release

### Styling & Output
- `utils.Style` provides colored output (RED, BOLD, etc.)
- `utils.CPrintln()` prints styled messages
- Used to highlight errors and status messages in terminal

## Go Version

Go 1.12+ (specified in go.mod); workflows test with Go 1.22

## Development Setup

```bash
git clone https://github.com/gopinath-langote/1build
cd 1build
go build
```

Ensure `go.sum` is up-to-date after modifying dependencies. CI runs on both ubuntu-latest and macos-latest.
