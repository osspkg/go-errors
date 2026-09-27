# Wrapping and inspection patterns

## Add operation context

Use `Wrapf` when an underlying error exists and callers may need to inspect it:

```go
return pkgerrors.Wrapf(err, "read account %q", accountID)
```

A nil cause produces nil, even if a context message was supplied. Do not use `Wrapf` to create a standalone error; use `New` or `fmt.Errorf` for that case.

## Inspect by identity or type

Use `errors.Is` for sentinels and `errors.As` for typed errors. The package also exports `Is` and `As` as direct forwards to the standard library.

```go
if errors.Is(err, ErrNotFound) {
    // Handle absence.
}

var parseErr *ParseError
if errors.As(err, &parseErr) {
    // Use fields from ParseError.
}
```

Do not compare `Error()` strings to identify error kinds.

## Combine independent failures

Use `pkgerrors.Wrap(errA, errB)` when both errors must be retained. Its string is the colon-separated input text; `errors.Is` and `errors.As` inspect every cause. `pkgerrors.Cause` preserves the legacy contract and returns the last non-nil input, so it is not a substitute for inspecting all causes.

For cleanup that should stop at the first failure, use `Queue` instead of collecting every error. `Queue` is sequential and stops immediately after a callback returns an error.

## Legacy Cause versus modern wrapping

`Cause` follows only `Cause() error`. It does not traverse standard `fmt.Errorf("context: %w", err)` wrappers. Prefer `%w`, `errors.Is`, and `errors.As` for new APIs; use `Cause` when maintaining compatibility with errors that implement the legacy interface.

## Capture a trace

`Trace` adds context and a runtime stack trace to a non-nil cause. The trace is captured at the call site. Its output includes multiple `[trace]` entries, so callers should avoid adding it where a large error string is undesirable.
