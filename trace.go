/*
 *  Copyright (c) 2024 Mikhail Knyazhev <markus621@yandex.ru>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package errors

import (
	"fmt"
	"runtime"
)

func runtimeTrace(c int) string {
	list := make([]uintptr, c)

	n := runtime.Callers(4, list)
	frame := runtime.CallersFrames(list[:n])

	result := ""
	for {
		v, ok := frame.Next()
		if !ok {
			break
		}
		result += fmt.Sprintf("\n\t[trace] %s:%d", v.Function, v.Line)
	}
	return result
}
