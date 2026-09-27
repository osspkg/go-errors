/*
 *  Copyright (c) 2024-2026 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package errors

import (
	"errors"
	"fmt"
	"runtime"
)

const (
	traceDepth       = 10
	traceCallersSkip = 4
)

// Trace wraps cause with a formatted message and appends a stack trace.
// It returns nil when cause is nil.
func Trace(cause error, message string, args ...any) error {
	err := Wrapf(cause, message, args...)
	if err == nil {
		return nil
	}

	wrapped := func() *errorEntity {
		target := &errorEntity{}
		_ = errors.As(err, &target)
		return target
	}()
	wrapped.trace = runtimeTrace(traceDepth)
	return wrapped
}

func runtimeTrace(depth int) string {
	pcs := make([]uintptr, depth)
	n := runtime.Callers(traceCallersSkip, pcs)
	frames := runtime.CallersFrames(pcs[:n])

	var result string
	for {
		frame, more := frames.Next()
		if !more {
			break
		}
		result += fmt.Sprintf("\n\t[trace] %s:%d", frame.Function, frame.Line)
	}
	return result
}
