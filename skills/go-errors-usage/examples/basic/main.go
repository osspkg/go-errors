/*
 *  Copyright (c) 2024-2026 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

// Package main demonstrates wrapping, combining, and inspecting errors.
package main

import (
	"errors"
	"fmt"

	pkgerrors "go.osspkg.com/errors"
)

const (
	accountID        = 42
	statusBadGateway = 502
)

var errUnavailable = errors.New("service unavailable")

type responseError struct {
	status int
}

func (e *responseError) Error() string {
	return fmt.Sprintf("unexpected status %d", e.status)
}

func fetch() error {
	return pkgerrors.Wrapf(errUnavailable, "fetch account %d", accountID)
}

func main() {
	err := fetch()
	if errors.Is(err, errUnavailable) {
		fmt.Println("fetch failed:", err)
	}

	statusErr := &responseError{status: statusBadGateway}
	combined := pkgerrors.Wrap(err, statusErr)
	var target *responseError
	fmt.Println("contains status error:", errors.As(combined, &target))
	fmt.Println("contains unavailable error:", errors.Is(combined, errUnavailable))
}
