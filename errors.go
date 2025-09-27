/*
 *  Copyright (c) 2024-2025 Mikhail Knyazhev <markus621@yandex.ru>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package errors

import "strings"

type errorEntity struct {
	cause   error
	message string
}

func New(message string) error {
	return &errorEntity{message: message}
}

func (v *errorEntity) Error() string {
	if v.cause == nil && len(v.message) == 0 {
		return ""
	}

	var b strings.Builder

	var mw bool
	if len(v.message) > 0 {
		b.WriteString(v.message)
		mw = true
	}

	if v.cause != nil {
		if mw {
			b.WriteString(": ")
		}

		b.WriteString(v.cause.Error())
	}

	return b.String()
}

func (v *errorEntity) Cause() error {
	return v.cause
}

func (v *errorEntity) Unwrap() error {
	return v.cause
}
