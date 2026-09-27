# Agent instructions

## Project

- This repository is the Go module `go.osspkg.com/errors`, a single package at the repository root. It targets Go 1.26.
- Keep the public error API compatible. `Trace` is exported; `Wrap` with multiple inputs must preserve each input for standard `errors.Is` and `errors.As` while retaining the colon-separated error text and legacy `Cause` behavior.
- Tests for the public API use the external package `errors_test`.

## Commands

Run commands from the repository root.

- `make lint` runs `goppy lint`, which performs Go module tidy/download and formatting before `golangci-lint` and `govulncheck`. It can change files; inspect `git status` and the diff afterward.
- `make tests` runs `goppy test`.
- `make build` runs `goppy build --arch=amd64`.
- CI runs `make ci`. This target also installs `goppy@latest`, runs `goppy setup-lib` and the license target, then lint, tests, and build. Use it when the full CI sequence and local setup are intended.

## Project memory

For non-trivial work, use the Chroma collection `chat_go-errors_memory`:

1. Ensure the collection exists before reading or writing. If missing, list collections and create this exact collection with the default embedding configuration.
2. Query it with a concise semantic description of the current task before making design decisions.
3. After the work, add a concise document only for a durable project decision or lesson. Query related memories first and update an existing document instead of duplicating it.
4. Never store secrets, transcripts, or temporary command output. Memory supplements the current source and tests; it does not override them.
