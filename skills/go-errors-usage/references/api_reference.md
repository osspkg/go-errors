# API reference

The import path is `go.osspkg.com/errors`; the package name is `errors`. The module targets Go 1.26 and has no third-party runtime dependencies.

| API | Behavior |
| --- | --- |
| `New(message string) error` | Creates an error containing `message`. The returned error is non-nil even if the message is empty. |
| `Wrapf(cause error, message string, args ...any) error` | Returns nil for a nil cause. Otherwise stores the literal message when there are no args, or formats it with `fmt.Sprintf` when args are present, then wraps the cause. |
| `Wrap(causes ...error) error` | Skips nil interface values, returns nil if none remain, and joins their messages with `: `. With multiple causes, each remains visible to standard `errors.Is` and `errors.As`. |
| `Trace(cause error, message string, args ...any) error` | Like `Wrapf`, and appends captured runtime frames in `[trace] function:line` form. Returns nil for a nil cause. |
| `Queue(calls ...func() error) error` | Executes callbacks sequentially, returning the first non-nil error. Empty input returns nil. A nil callback panics when reached. |
| `Unwrap(err error) error` | Calls `Unwrap() error` once when implemented. Returns nil for nil input, unsupported errors, and multi-cause errors. |
| `Cause(err error) error` | Follows the legacy `Cause() error` interface until reaching an error without it. For a multi-error returned by `Wrap`, this remains the last non-nil input. |
| `Is(err, target error) bool` | Forwards to the standard `errors.Is`. |
| `As(err error, target any) bool` | Forwards to the standard `errors.As`. |
| `Causer` | Interface containing `Cause() error`. |
| `Unwrapper` | Interface containing `Unwrap() error` for single-cause errors. |

## Interoperation notes

- `Wrapf` and `Trace` create errors that implement both `Cause() error` and `Unwrap() error`.
- Multi-error `Wrap` implements `Unwrap() []error`; standard `errors.Is` and `errors.As` traverse all causes. The package helper `Unwrap` intentionally supports only the single-error interface, like the standard `errors.Unwrap` helper.
- `Cause` does not follow `Unwrap() error`. For standard `%w` wrappers and joined errors, inspect with `Is` or `As`.
- `Trace` captures frames during the call, so invoke it at the point where the trace should begin.
