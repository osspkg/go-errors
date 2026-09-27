/*
 *  Copyright (c) 2024-2026 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

// Package errors provides error construction, wrapping, and inspection helpers.
package errors

import "strings"

type errorEntity struct {
	cause   error
	message string
	trace   string
}

// New returns an error with the provided message.
func New(message string) error {
	return &errorEntity{message: message}
}

func (v *errorEntity) Error() string {
	if v.cause == nil {
		return v.message + v.trace
	}

	cause := v.cause.Error()
	var message strings.Builder
	message.Grow(len(v.message) + len(cause) + len(v.trace) + 2)
	message.WriteString(v.message)
	if v.message != "" {
		message.WriteString(": ")
	}
	message.WriteString(cause)
	message.WriteString(v.trace)
	return message.String()
}

// Cause returns the underlying cause.
func (v *errorEntity) Cause() error {
	return v.cause
}

func (v *errorEntity) Unwrap() error {
	return v.cause
}
