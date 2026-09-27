---
name: go-errors-usage
description: Use the go.osspkg.com/errors library when creating, wrapping, combining, tracing, or inspecting errors in Go code. Apply when this repository's error API or its compatibility behavior matters; use general Go error guidance for code that does not depend on this package.
---

# go-errors usage

Use this skill when implementing or reviewing call sites of `go.osspkg.com/errors`.
Start with the closest API reference, then use a runnable example for the relevant pattern.

- Read [API reference](references/api-reference.md) for signatures, nil behavior, and compatibility details.
- Read [wrapping and inspection](references/wrapping-and-inspection.md) when choosing how to preserve causes or inspect an error chain.
- Browse [`examples/`](examples/) for complete Go programs. Run one from the repository root with `go run ./skills/go-errors-usage/examples/<name>`.

Keep standard error inspection semantics intact. Prefer `errors.Is` and `errors.As` (standard library or this package's forwarding helpers) over comparing error strings. Add context with `Wrapf` only when a non-nil cause exists; it returns nil for a nil cause. Use `Wrap` to combine independent errors when every cause must remain discoverable.

Preserve these API contracts:

- `Trace` is public and captures a runtime stack trace when called.
- Multi-error `Wrap` keeps the existing colon-separated message and exposes each input to `errors.Is` and `errors.As`; legacy `Cause` returns its final input.
- The package-level `Unwrap` handles a single `Unwrap() error`; it returns nil for a multi-error `Unwrap() []error`.
- `Cause` follows legacy `Cause() error` methods. It does not walk standard `%w` wrappers.
- `Queue` calls non-nil callbacks in order and stops at the first error.

The package name conflicts with the standard library package name. In examples that import both, alias the project package (for example, `pkgerrors`) and leave the standard package as `errors`.
