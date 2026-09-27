/*
 *  Copyright (c) 2024-2026 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

// Package main demonstrates sequential callbacks with Queue.
package main

import (
	"fmt"

	pkgerrors "go.osspkg.com/errors"
)

func main() {
	if err := pkgerrors.Queue(
		func() error {
			fmt.Println("validate configuration")
			return nil
		},
		func() error {
			fmt.Println("connect to database")
			return pkgerrors.New("connection refused")
		},
		func() error {
			fmt.Println("this callback is not called")
			return nil
		},
	); err != nil {
		fmt.Println("startup stopped:", err)
	}
}
