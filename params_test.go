package ics

import (
	"errors"
	"reflect"
	"strings"
	"testing"

	"vimagination.zapto.org/parser"
)

type params interface {
	decode([]parser.Token) error
	encode(writer)
}

type paramTests[T params] struct {
	Input  []parser.Token
	Param  T
	Error  error
	Output string
}

func testParams[T params](t *testing.T, tests []paramTests[T]) {
	for n, test := range tests {
		v := reflect.New(reflect.TypeFor[T]().Elem()).Interface().(params)

		if err := v.decode(test.Input); !errors.Is(err, test.Error) {
			if errv := reflect.New(reflect.TypeOf(test.Error)); !errors.As(err, errv.Interface()) || !reflect.DeepEqual(test.Error, errv.Elem().Interface().(error)) {
				t.Errorf("test %d: expecting error %v, got %v", n+1, test.Error, err)

				continue
			}
		}

		if test.Error == nil {
			if !reflect.DeepEqual(v, test.Param) {
				t.Errorf("test %d: expecting param %#v, got %#v", n+1, test.Param, v)
			} else if test.Output != "" {
				var buf strings.Builder

				v.encode(&buf)

				if out := buf.String(); out != test.Output {
					t.Errorf("test %d: expecting output %q, got %q", n+1, test.Output, out)
				}
			}
		}
	}
}
