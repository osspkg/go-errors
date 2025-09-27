/*
 *  Copyright (c) 2024-2025 Mikhail Knyazhev <markus621@yandex.ru>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package errors

func Queue(calls ...func() error) error {
	for _, call := range calls {
		if err := call(); err != nil {
			return err
		}
	}
	return nil
}
