/*
 *  Copyright (c) 2024-2025 Mikhail Knyazhev <markus621@yandex.ru>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package errors_test

import (
	"testing"

	"go.osspkg.com/errors"
)

func TestUnit_Queue(t *testing.T) {
	incr1 := 0
	incr2 := 0

	err := errors.Queue(
		func() error {
			incr1++
			return errors.New("1")
		},
		func() error {
			incr2++
			return errors.New("2")
		},
	)

	if err == nil {
		t.Fatalf("err = nil, want err")
	}

	if incr1 != 1 {
		t.Fatalf("incr1 = %v, want %v", incr1, 1)
	}

	if incr2 != 0 {
		t.Fatalf("incr2 = %v, want %v", incr2, 0)
	}
}
