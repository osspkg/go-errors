/*
 *  Copyright (c) 2024-2026 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package errors

// Causer is implemented by errors that expose a legacy cause.
type Causer interface {
	Cause() error
}

// Unwrapper is implemented by errors that expose one underlying error.
type Unwrapper interface {
	Unwrap() error
}
