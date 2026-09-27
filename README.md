# go-errors

[![CI](https://github.com/osspkg/go-errors/actions/workflows/ci.yml/badge.svg)](https://github.com/osspkg/go-errors/actions/workflows/ci.yml)

A small Go library for creating errors, adding context, retaining causes, and inspecting error chains. It has no third-party runtime dependencies and requires Go 1.26 or newer.

## Installation

```sh
go get go.osspkg.com/errors
```

## Quick start

```go
package main

import (
	"errors"
	"fmt"

	errpkg "go.osspkg.com/errors"
)

var errConnectionRefused = errors.New("connection refused")

func main() {
	err := errpkg.Wrapf(errConnectionRefused, "dial %s", "db.example:5432")
	if errors.Is(err, errConnectionRefused) {
		fmt.Println(err)
	}
}
```

## API

| API | Purpose |
| --- | --- |
| `New(message string) error` | Creates an error with a message. |
| `Wrapf(cause error, message string, args ...any) error` | Adds formatted context to one cause; returns `nil` when `cause` is `nil`. |
| `Wrap(errors ...error) error` | Combines non-nil errors, joining their messages with `: `. Every input remains discoverable through `errors.Is` and `errors.As`; returns `nil` if there are no non-nil inputs. |
| `Trace(cause error, message string, args ...any) error` | Adds formatted context and a runtime stack trace; returns `nil` when `cause` is `nil`. |
| `Queue(calls ...func() error) error` | Calls functions in order and returns the first error, or `nil` if all succeed. Callbacks must be non-nil. |
| `Unwrap(err error) error` | Returns one underlying error from an `Unwrapper`; returns `nil` for nil errors and multi-cause errors. |
| `Cause(err error) error` | Follows the legacy `Cause() error` chain and returns its terminal error. |
| `Is(err, target error) bool` | Reports whether `err` or an error in its chain matches `target`. |
| `As(err error, target any) bool` | Finds the first error in the chain assignable to `target`, following the standard `errors.As` contract. |
| `Causer` | Interface for errors exposing `Cause() error`. |
| `Unwrapper` | Interface for errors exposing `Unwrap() error`. |

`Cause` follows legacy `Cause()` methods; it does not walk standard-library `%w` wrappers. Use `Is` or `As` to inspect those chains. For a `Wrap` call with several errors, `Cause` retains its legacy behavior and returns the final input, while `Is` and `As` inspect all inputs.

## Development

The repository uses Go 1.26 and `goppy` for its Makefile targets.

```sh
make lint
make tests
make build
```

CI runs `make ci`, which installs `goppy@latest`, runs `goppy setup-lib` and license generation, then runs lint, tests, and build.

## License

BSD 3-Clause. See [LICENSE](LICENSE).
