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

func testOptions[T params](t *testing.T, outputPrefix string, options map[string]T) {
	tests := []paramTests[T]{
		{ // 1
			Input: []parser.Token{},
			Error: ErrInvalidParam,
		},
		{ // 2
			Input: []parser.Token{{}, {}},
			Error: ErrInvalidParam,
		},
	}

	for opt, v := range options {
		for i := range 1 << len(opt) {
			var sb strings.Builder

			for n, c := range opt {
				if c < 'A' || c > 'Z' {
					sb.WriteByte(byte(c))

					continue
				}

				if i&(1<<n) == 0 {
					sb.WriteByte(byte(c) | 0x20)
				} else {
					sb.WriteByte(byte(c) & 0xDF)
				}
			}

			tests = append(tests, paramTests[T]{
				Input:  []parser.Token{{Data: sb.String()}},
				Param:  v,
				Output: outputPrefix + string(opt),
			})
		}
	}

	if _, ok := options["UNKNOWN"]; !ok {
		tests = append(tests, paramTests[T]{
			Input: []parser.Token{{Data: "UnKnOwN"}},
			Error: ErrInvalidParam,
		})
	}

	testParams(t, tests)
}

func TestParamAlternativeRepresentation(t *testing.T) {
	u, _ := url.Parse("http://www.example.com")

	testParams(t, []paramTests[*ParamAlternativeRepresentation]{
		{ // 1
			Input: []parser.Token{},
			Error: ErrInvalidParam,
		},
		{ // 2
			Input: []parser.Token{
				{Type: tokenParamName, Data: "http://www.example.com"},
			},
			Error: ErrInvalidParam,
		},
		{ // 3
			Input: []parser.Token{
				{Type: tokenParamQuotedValue, Data: "http://www.example.com"},
			},
			Param:  &ParamAlternativeRepresentation{*u},
			Output: ";ALTREP=http://www.example.com",
		},
		{ // 4
			Input: []parser.Token{
				{Type: tokenParamQuotedValue, Data: "http://www.example.com"},
				{Type: tokenParamQuotedValue, Data: "http://www.example2.com"},
			},
			Error: ErrInvalidParam,
		},
		{ // 5
			Input: []parser.Token{
				{Type: tokenParamQuotedValue, Data: "http://a%:b@example.com"},
			},
			Error: url.EscapeError("%"),
		},
	})
}

func TestParamCommonName(t *testing.T) {
	testParams(t, []paramTests[*ParamCommonName]{
		{ // 1
			Input: []parser.Token{},
			Error: ErrInvalidParam,
		},
		{ // 2
			Input: []parser.Token{
				{Type: tokenParamValue, Data: ""},
			},
			Param: NewCommonName(""),
		},
		{ // 3
			Input: []parser.Token{
				{Type: tokenParamValue, Data: "abc"},
				{Type: tokenParamValue, Data: "def"},
			},
			Error: ErrInvalidParam,
		},
		{ // 4
			Input: []parser.Token{
				{Type: tokenParamValue, Data: "abc"},
			},
			Param:  NewCommonName("abc"),
			Output: ";CN=abc",
		},
		{ // 5
			Input: []parser.Token{
				{Type: tokenParamValue, Data: "abc;"},
			},
			Param:  NewCommonName("abc;"),
			Output: ";CN=\"abc;\"",
		},
	})
}

func TestParamCalendarUserType(t *testing.T) {
	testOptions(t, ";CUTYPE=", map[string]*ParamCalendarUserType{
		"INDIVIDUAL": new(CalendarUserTypeIndividual),
		"GROUP":      new(CalendarUserTypeGroup),
		"RESOURCE":   new(CalendarUserTypeResource),
		"ROOM":       new(CalendarUserTypeRoom),
		"UNKNOWN":    new(CalendarUserTypeUnknown),
	})
}

func TestParamDelegator(t *testing.T) {
	u1, _ := url.Parse("http://www.example.com")
	u2, _ := url.Parse("http://www.another.com")

	testParams(t, []paramTests[*ParamDelegator]{
		{ // 1
			Input: []parser.Token{},
			Param: new(ParamDelegator),
		},
		{ // 2
			Input: []parser.Token{
				{Type: tokenParamValue, Data: ""},
			},
			Error: ErrInvalidParam,
		},
		{ // 3
			Input: []parser.Token{
				{Type: tokenParamQuotedValue, Data: "http://www.example.com"},
				{Type: tokenParamValue, Data: ""},
			},
			Error: ErrInvalidParam,
		},
		{ // 4
			Input: []parser.Token{
				{Type: tokenParamQuotedValue, Data: "http://www.example.com"},
			},
			Param: &ParamDelegator{
				{*u1},
			},
			Output: ";DELEGATED-FROM=http://www.example.com",
		},
		{ // 5
			Input: []parser.Token{
				{Type: tokenParamQuotedValue, Data: "http://a%:@www.example.com"},
			},
			Error: url.EscapeError("%"),
		},
		{ // 6
			Input: []parser.Token{
				{Type: tokenParamQuotedValue, Data: "http://www.example.com"},
				{Type: tokenParamQuotedValue, Data: "http://www.another.com"},
			},
			Param: &ParamDelegator{
				{*u1},
				{*u2},
			},
			Output: ";DELEGATED-FROM=http://www.example.com,http://www.another.com",
		},
	})
}
