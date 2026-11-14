package ics

import (
	"errors"
	"net/url"
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

func TestParamAlternativeRepresentation(t *testing.T) {
	u, _ := url.Parse("http://www.example.com")

	testParams(t, []paramTests[*ParamAlternativeRepresentation]{
		{
			Input: []parser.Token{},
			Error: ErrInvalidParam,
		},
		{
			Input: []parser.Token{
				{Type: tokenParamName, Data: "http://www.example.com"},
			},
			Error: ErrInvalidParam,
		},
		{
			Input: []parser.Token{
				{Type: tokenParamQuotedValue, Data: "http://www.example.com"},
			},
			Param:  &ParamAlternativeRepresentation{*u},
			Output: ";ALTREP=http://www.example.com",
		},
		{
			Input: []parser.Token{
				{Type: tokenParamQuotedValue, Data: "http://www.example.com"},
				{Type: tokenParamQuotedValue, Data: "http://www.example2.com"},
			},
			Error: ErrInvalidParam,
		},
		{
			Input: []parser.Token{
				{Type: tokenParamQuotedValue, Data: "http://a%:b@example.com"},
			},
			Error: url.EscapeError("%"),
		},
	})
}

func TestParamCommonName(t *testing.T) {
	testParams(t, []paramTests[*ParamCommonName]{
		{
			Input: []parser.Token{},
			Error: ErrInvalidParam,
		},
		{
			Input: []parser.Token{
				{Type: tokenParamValue, Data: ""},
			},
			Param: NewCommonName(""),
		},
		{
			Input: []parser.Token{
				{Type: tokenParamValue, Data: "abc"},
				{Type: tokenParamValue, Data: "def"},
			},
			Error: ErrInvalidParam,
		},
		{
			Input: []parser.Token{
				{Type: tokenParamValue, Data: "abc"},
			},
			Param:  NewCommonName("abc"),
			Output: ";CN=abc",
		},
		{
			Input: []parser.Token{
				{Type: tokenParamValue, Data: "abc;"},
			},
			Param:  NewCommonName("abc;"),
			Output: ";CN=\"abc;\"",
		},
	})
}
