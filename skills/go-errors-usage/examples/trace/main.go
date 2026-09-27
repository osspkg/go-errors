/*
 *  Copyright (c) 2024-2026 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

// Package main demonstrates capturing a runtime error trace.
package main

import (
	"fmt"

	pkgerrors "go.osspkg.com/errors"
)

func loadConfig() error {
	return pkgerrors.Trace(pkgerrors.New("file not found"), "load application config")
}

func main() {
	if err := loadConfig(); err != nil {
		fmt.Println(err)
	}
}
