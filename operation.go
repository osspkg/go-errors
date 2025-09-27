/*
 *  Copyright (c) 2024-2025 Mikhail Knyazhev <markus621@yandex.ru>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package errors

import (
	e "errors"
	"fmt"
)

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

func Wrap(messages ...error) error {
	if len(messages) == 0 {
		return nil
	}

	var err error

	for _, msg := range messages {
		if msg == nil {
			continue
		}
		if err == nil {
			err = &errorEntity{cause: msg}
			continue
		}
		err = &errorEntity{
			cause:   msg,
			message: err.Error(),
		}
	}

	return err
}

func Unwrap(err error) error {
	if err == nil {
		return nil
	}

	if v, ok := err.(Unwrapper); ok {
		return v.Unwrap()
	}

	return nil
}

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

func Is(err, target error) bool {
	return e.Is(err, target)
}

func As(err error, target any) bool {
	return e.As(err, target)
}
