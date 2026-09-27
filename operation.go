/*
 *  Copyright (c) 2024-2026 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package errors

import (
	e "errors"
	"fmt"
)

// Wrapf returns nil when cause is nil. Otherwise it prefixes cause with a
// formatted message and preserves cause for errors.Is and errors.As.
func Wrapf(cause error, message string, args ...interface{}) error {
	if cause == nil {
		return nil
	}

	err := &errorEntity{
		cause:   cause,
		message: message,
	}

	if len(args) > 0 {
		err.message = fmt.Sprintf(message, args...)
	}

	return err
}

// Wrap combines non-nil errors into one error. Its message joins the input
// messages with ": ", and errors.Is and errors.As can inspect every input.
// Wrap returns nil when all inputs are nil.
func Wrap(messages ...error) error {
	causes := make([]error, 0, len(messages))
	for _, msg := range messages {
		if msg == nil {
			continue
		}
		causes = append(causes, msg)
	}

	switch len(causes) {
	case 0:
		return nil
	case 1:
		return &errorEntity{cause: causes[0]}
	default:
		return &joinedError{causes: causes}
	}
}

// Unwrap returns the single wrapped error when err implements Unwrapper.
// It returns nil for nil errors and errors with multiple causes.
func Unwrap(err error) error {
	if err == nil {
		return nil
	}

	if v, ok := err.(Unwrapper); ok {
		return v.Unwrap()
	}

	return nil
}

// Cause follows legacy Cause() methods and returns the first error that does
// not implement Causer. For errors wrapped with %w, use Is or As to inspect
// the chain.
func Cause(err error) error {
	if err == nil {
		return nil
	}

	for {
		if c, ok := err.(Causer); ok && c != nil {
			err = c.Cause()
			continue
		}

		return err
	}
}

// Is reports whether err or any error in its chain matches target.
func Is(err, target error) bool {
	return e.Is(err, target)
}

// As finds the first error in err's chain assignable to target and stores it
// there, following the standard errors.As contract.
func As(err error, target any) bool {
	return e.As(err, target)
}

type joinedError struct {
	causes []error
}

func (v *joinedError) Error() string {
	message := ""
	for _, cause := range v.causes {
		if message != "" {
			message += ": "
		}
		message += cause.Error()
	}
	return message
}

func (v *joinedError) Unwrap() []error {
	return append([]error(nil), v.causes...)
}

// Cause returns the final error passed to Wrap.
func (v *joinedError) Cause() error {
	return v.causes[len(v.causes)-1]
}
