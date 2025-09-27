/*
 *  Copyright (c) 2024-2025 Mikhail Knyazhev <markus621@yandex.ru>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package errors

import (
	e "errors"
	"fmt"
	"testing"
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
			err := New(tt.message)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err.Error() != tt.want {
				t.Errorf("New() error = %v, want %v", err.Error(), tt.want)
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
		{
			name:    "Case1",
			msgs:    nil,
			want:    "",
			wantErr: false,
		},
		{
			name:    "Case2",
			msgs:    []error{New("hello"), e.New("world")},
			want:    "hello: world",
			wantErr: true,
		},
		{
			name:    "Case3",
			msgs:    []error{New("err1"), e.New("err2"), nil, e.New("err3")},
			want:    "err1: err2: err3",
			wantErr: true,
		},
		{
			name: "Case4",
			msgs: []error{Wrapf(New("err1"), "err1 message"),
				Wrapf(e.New("err2"), "err2 message"),
				Wrapf(e.New("err3"), "err3 message")},
			want:    "err1 message: err1: err2 message: err2: err3 message: err3",
			wantErr: true,
		},
		{
			name:    "Case5",
			msgs:    []error{nil, nil, nil},
			want:    "",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Wrap(tt.msgs...)

			if tt.wantErr {
				if err == nil {
					t.Errorf("Wrap() error = %v, wantErr %v", err, tt.wantErr)
					return
				}

				if err.Error() != tt.want {
					t.Errorf("Wrap() error = %v, want %v", err.Error(), tt.want)
				}
			} else {
				if err != nil {
					t.Errorf("Wrap() error = %v, wantErr %v", err, tt.wantErr)
				}
			}
		})
	}
}

func TestUnit_Wrapf(t *testing.T) {
	type args struct {
		cause   error
		message string
		args    []interface{}
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
				args:    []interface{}{"127.0.0.1"},
			},
			want:    "bad ip 127.0.0.1: err1",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Wrapf(tt.args.cause, tt.args.message, tt.args.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Wrapf() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr && err.Error() != tt.want {
				t.Errorf("Wrapf() error = %v, want %v", err.Error(), tt.want)
				return
			}
		})
	}
}

func TestUnit_CauseUnwrap(t *testing.T) {
	type fields struct {
		cause   error
		message string
		args    []interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{
			name: "Case1",
			fields: fields{
				cause:   e.New("err1"),
				message: "context",
			},
			want:    "err1",
			wantErr: true,
		},
		{
			name: "Case2",
			fields: fields{
				cause:   nil,
				message: "context",
			},
			want:    "err1",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := Wrapf(tt.fields.cause, tt.fields.message, tt.fields.args...)
			err := Cause(v)
			if (err != nil) != tt.wantErr {
				t.Errorf("Cause() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr && err.Error() != tt.want {
				t.Errorf("Cause() error = %v, want %v", err.Error(), tt.want)
				return
			}
			err = Unwrap(v)
			if (err != nil) != tt.wantErr {
				t.Errorf("Unwrap() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr && err.Error() != tt.want {
				t.Errorf("Unwrap() error = %v, want %v", err.Error(), tt.want)
				return
			}
		})
	}
}

func TestUnit_Is(t *testing.T) {
	err0 := New("test")
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
		{name: "Case2", args: args{err: Wrapf(err0, "ttt"), target: err0}, want: true},
		{name: "Case3", args: args{err: New("hello"), target: err0}, want: false},
		{name: "Case4", args: args{err: nil, target: err0}, want: false},
		{name: "Case5", args: args{err: New("hello"), target: nil}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Is(tt.args.err, tt.args.target); got != tt.want {
				t.Errorf("Is() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnit_As(t *testing.T) {
	err0 := New("err0")
	err1 := fmt.Errorf("err1")

	var err2 *errorEntity
	if !As(err0, &err2) {
		t.Errorf("As() error = %v, wantErr %v", err0, err2)
		return
	}

	if As(err1, &err2) {
		t.Errorf("As() error = %v, wantErr %v", err1, err2)
		return
	}
}
