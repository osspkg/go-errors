/*
 *  Copyright (c) 2024-2026 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

// Package errors provides error construction, wrapping, and inspection helpers.
package errors

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
	message := v.message
	if v.cause != nil {
		if message != "" {
			message += ": "
		}
		message += v.cause.Error()
	}
	return message + v.trace
}

// Cause returns the underlying cause.
func (v *errorEntity) Cause() error {
	return v.cause
}

func (v *errorEntity) Unwrap() error {
	return v.cause
}
