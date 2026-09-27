/*
 *  Copyright (c) 2024-2026 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package errors

import (
	"errors"
	"runtime"
	"strconv"
	"strings"
)

const (
	traceDepth       = 10
	traceCallersSkip = 4
	decimalBase      = 10
)

// Trace wraps cause with a formatted message and appends a stack trace.
// It returns nil when cause is nil.
func Trace(cause error, message string, args ...any) error {
	err := Wrapf(cause, message, args...)
	if err == nil {
		return nil
	}

	wrapped := &errorEntity{}
	if !errors.As(err, &wrapped) {
		wrapped.cause = err
	}

	wrapped.trace = runtimeTrace(traceDepth)
	return wrapped
}

func runtimeTrace(depth int) string {
	pcs := make([]uintptr, depth)
	n := runtime.Callers(traceCallersSkip, pcs)
	frames := runtime.CallersFrames(pcs[:n])

	var result strings.Builder
	var lineBuffer [20]byte
	for {
		frame, more := frames.Next()
		if !more {
			break
		}
		_, _ = result.WriteString("\n\t[trace] ")
		_, _ = result.WriteString(frame.Function)
		_ = result.WriteByte(':')
		_, _ = result.Write(strconv.AppendInt(lineBuffer[:0], int64(frame.Line), decimalBase))
	}
	return result.String()
}
