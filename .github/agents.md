# AGENTS

This file is guidance for agentic coding tools working in this repo.
It consolidates build, lint, and test commands plus style conventions.
It also includes existing Copilot instructions verbatim where possible.

## Project snapshot
- Language: Go (module `github.com/gopinath-langote/1build`)
- Product: CLI tool that runs project-local command aliases (`1build`)
- CLI framework: Cobra, config via Viper and YAML
- Tests: end-to-end style that builds a temp binary and runs fixtures

## Must-read external rules
### Copilot instructions (.github/copilot-instructions.md)
- Build: `go build`
- Tests (full suite): `go test -v -cover github.com/gopinath-langote/1build/testing -run .`
- Tests with coverage cache off: `go test -v -cover github.com/gopinath-langote/1build/testing -run . GOCACHE=off`
- Lint: `go list ./... | xargs -L1 golint` (requires `golint` install)
- Project uses `1build` itself: `1build test`, `1build lint`, `1build build`
- CI uses Go 1.22, module go 1.12, release via goreleaser

## Build, lint, test
### Build
- `go build` (builds binary in current directory)
- `go build -v .` (matches CI)

### Lint
- `go list ./... | xargs -L1 golint`
- `go get -u golang.org/x/lint/golint` (install lint tool)
- CI also runs golangci-lint with reviewdog: `reviewdog/action-golangci-lint` (enable-all)

### Tests
- Full suite (recommended):
  `go test -v -cover github.com/gopinath-langote/1build/testing -run .`
- With coverage cache off:
  `go test -v -cover github.com/gopinath-langote/1build/testing -run . GOCACHE=off`

### Single test
- Run specific test by name:
  `go test -v github.com/gopinath-langote/1build/testing -run TestAll`
- Run by subtest name (fixture-based):
  `go test -v github.com/gopinath-langote/1build/testing -run "TestAll/\.:\.<feature>\.:\.<name>"`
- If needed, run an individual package test file locally:
  `go test -v ./testing -run TestAll`

### 1build shortcuts (project local)
From `1build.yaml` in repo root:
- `1build build`
- `1build lint`
- `1build test`

## Code style and conventions
### Formatting
- Use `gofmt` for all Go files. Keep formatting stable and idiomatic.
- Keep lines reasonably short; docs lint allows up to 190 chars.

### Imports
- Group imports in standard Go order:
  1) standard library
  2) third party
  3) local module (`github.com/gopinath-langote/1build/...`)
- Avoid unused imports; keep import names short.

### Package structure
- CLI commands live under `cmd/` with Cobra commands per folder.
- Command execution logic in `cmd/exec` and configuration in `cmd/config`.
- Tests in `testing/` package and fixtures in `testing/fixtures`.

### Naming
- Use Go standard naming: PascalCase for exported, camelCase for local.
- Command names are lowercase CLI tokens (`set`, `unset`, `exec`).
- Keep error and log messages user-facing and consistent.

### Types and data flow
- Prefer explicit struct types for command/config models.
- Follow existing `models.CommandContext` usage for execution phases.
- When adding fields, update YAML parsing and validation accordingly.

### Error handling
- Follow existing pattern: print error, then `utils.ExitError()` or `ExitWithCode`.
- For CLI errors, prefer exit code 1 or specific code (127 for missing command).
- Avoid panics; use returned errors and explicit exits.

### Logging and output
- Use `utils.CPrint`/`utils.CPrintln` for styled output.
- Respect `--quiet` flag; when quiet, suppress logs and only print success/failure.
- When executing commands, show phase banners and command labels.

### Tests and fixtures
- Tests build a temporary `1build` binary and run it against fixture dirs.
- Fixtures are defined in `testing/fixtures` and consumed by `fixtures.GetFixtures()`.
- Subtests are named with pattern `".:.<feature>.:.<name>"`.

### Dependencies and modules
- Use Go modules; keep `go.mod` and `go.sum` updated.
- Key deps: cobra, viper, go-sh, yaml.v3, aurora, testify.

## Repo hygiene
- Do not commit build artifacts in `dist/` or binary `1build`.
- Keep `1build.yaml` in repo (explicitly unignored).

## Notes for agents
- CI runs on ubuntu and macOS with Go 1.22.
- Module target is go 1.12; avoid features requiring newer language version.
- If you add docs, keep Markdown consistent with `.remarkrc.yaml`.

## Pointers
- Main entry: `cmd/root.go`
- Execution flow: `cmd/exec/exec.go`
- Config parsing: `cmd/config/parse.go`
- Tests: `testing/cli_test.go`
