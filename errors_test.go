/*
 *  Copyright (c) 2024-2026 Mikhail Knyazhev <markus621@yandex.ru>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package errors_test

import (
	e "errors"
	"fmt"
	"strings"
	"testing"

	pkgerrors "go.osspkg.com/errors"
)

func TestUnit_New(t *testing.T) {
	tests := []struct {
		name    string
		message string
		want    string
		wantErr bool
	}{
		{name: "Case1", message: "hello", want: "hello", wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := pkgerrors.New(tt.message)
			if (err != nil) != tt.wantErr {
				t.Errorf("pkgerrors.New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err.Error() != tt.want {
				t.Errorf("pkgerrors.New() error = %v, want %v", err.Error(), tt.want)
				return
			}
		})
	}
}

func TestUnit_Wrap(t *testing.T) {
	tests := []struct {
		name    string
		msgs    []error
		want    string
		wantErr bool
	}{
		{name: "no inputs"},
		{
			name:    "two errors",
			msgs:    []error{pkgerrors.New("hello"), e.New("world")},
			want:    "hello: world",
			wantErr: true,
		},
		{
			name:    "skip nil inputs",
			msgs:    []error{pkgerrors.New("err1"), e.New("err2"), nil, e.New("err3")},
			want:    "err1: err2: err3",
			wantErr: true,
		},
		{
			name: "wrapped errors",
			msgs: []error{
				pkgerrors.Wrapf(pkgerrors.New("err1"), "err1 message"),
				pkgerrors.Wrapf(e.New("err2"), "err2 message"),
				pkgerrors.Wrapf(e.New("err3"), "err3 message"),
			},
			want:    "err1 message: err1: err2 message: err2: err3 message: err3",
			wantErr: true,
		},
		{name: "all nil", msgs: []error{nil, nil, nil}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := pkgerrors.Wrap(tt.msgs...)
			assertErrorResult(t, err, tt.want, tt.wantErr)
		})
	}
}

func TestUnit_Wrapf(t *testing.T) {
	type args struct {
		cause   error
		message string
		args    []any
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Case1",
			args: args{
				cause:   nil,
				message: "err context",
				args:    nil,
			},
			want:    "",
			wantErr: false,
		},
		{
			name: "Case2",
			args: args{
				cause:   e.New("err1"),
				message: "err context",
				args:    nil,
			},
			want:    "err context: err1",
			wantErr: true,
		},
		{
			name: "Case3",
			args: args{
				cause:   e.New("err1"),
				message: "bad ip %s",
				args:    []any{"127.0.0.1"},
			},
			want:    "bad ip 127.0.0.1: err1",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := pkgerrors.Wrapf(tt.args.cause, tt.args.message, tt.args.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("pkgerrors.Wrapf() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr && err.Error() != tt.want {
				t.Errorf("pkgerrors.Wrapf() error = %v, want %v", err.Error(), tt.want)
				return
			}
		})
	}
}

func TestUnit_CauseUnwrap(t *testing.T) {
	tests := []struct {
		name    string
		cause   error
		message string
		want    string
		wantErr bool
	}{
		{
			name:    "with cause",
			cause:   e.New("err1"),
			message: "context",
			want:    "err1",
			wantErr: true,
		},
		{name: "nil cause", message: "context"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := pkgerrors.Wrapf(tt.cause, "%s", tt.message)
			assertErrorResult(t, pkgerrors.Cause(err), tt.want, tt.wantErr)
			assertErrorResult(t, pkgerrors.Unwrap(err), tt.want, tt.wantErr)
		})
	}
}

func assertErrorResult(t *testing.T, got error, want string, wantErr bool) {
	t.Helper()
	if (got != nil) != wantErr {
		t.Errorf("error = %v, wantErr %v", got, wantErr)
		return
	}
	if wantErr && got.Error() != want {
		t.Errorf("error = %v, want %v", got.Error(), want)
	}
}

type typedTestError struct{}

func (*typedTestError) Error() string { return "typed" }

func TestWrapPreservesAllCauses(t *testing.T) {
	first := e.New("first")
	second := e.New("second")

	err := pkgerrors.Wrap(first, second)
	if got, want := err.Error(), "first: second"; got != want {
		t.Fatalf("pkgerrors.Wrap() = %q, want %q", got, want)
	}
	if !e.Is(err, first) {
		t.Error("pkgerrors.Wrap() should retain the first cause")
	}
	if !e.Is(err, second) {
		t.Error("pkgerrors.Wrap() should retain the second cause")
	}

	var got *typedTestError
	if !pkgerrors.As(pkgerrors.Wrap(first, &typedTestError{}), &got) || got == nil {
		t.Error("pkgerrors.As() should find a typed error in a combined error")
	}
}

func TestTrace(t *testing.T) {
	if pkgerrors.Trace(nil, "context") != nil {
		t.Fatal("pkgerrors.Trace(nil, ...) should return nil")
	}

	cause := e.New("cause")
	err := pkgerrors.Trace(cause, "context")
	if err == nil {
		t.Fatal("pkgerrors.Trace() = nil, want error")
	}
	if !strings.Contains(err.Error(), "[trace]") {
		t.Fatalf("pkgerrors.Trace() error %q does not contain a trace", err)
	}
	if !e.Is(err, cause) {
		t.Error("pkgerrors.Trace() should preserve the cause")
	}
}

func TestUnit_Is(t *testing.T) {
	err0 := pkgerrors.New("test")
	type args struct {
		err    error
		target error
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "Case1", args: args{err: err0, target: err0}, want: true},
		{name: "Case2", args: args{err: pkgerrors.Wrapf(err0, "ttt"), target: err0}, want: true},
		{name: "Case3", args: args{err: pkgerrors.New("hello"), target: err0}, want: false},
		{name: "Case4", args: args{err: nil, target: err0}, want: false},
		{name: "Case5", args: args{err: pkgerrors.New("hello"), target: nil}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pkgerrors.Is(tt.args.err, tt.args.target); got != tt.want {
				t.Errorf("pkgerrors.Is() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnit_As(t *testing.T) {
	typed := &typedTestError{}
	err0 := pkgerrors.Wrapf(typed, "context")
	err1 := e.New("err1")

	var got *typedTestError
	if !pkgerrors.As(err0, &got) || got != typed {
		t.Errorf("As() = %v, want %v", got, typed)
	}
	if pkgerrors.As(err1, &got) {
		t.Errorf("As(%v) = true, want false", err1)
	}
}

var benchmarkErrorSink error

func Benchmark_PkgError(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		err := pkgerrors.New("test")
		benchmarkErrorSink = pkgerrors.Wrapf(err, "Hello %d", 1)
	}
}

func Benchmark_SdkError(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		err := e.New("test")
		benchmarkErrorSink = fmt.Errorf("hello %d %w", 1, err)
	}
}
