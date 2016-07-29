package ics

import (
	"bytes"
	"io"
	"reflect"
	"strconv"
	"testing"
)

type typeTest struct {
	Params map[string]string
	Data   string
	Input  interface {
		Decode(map[string]string, string) error
	}
	Match interface {
		Encode(io.Writer)
	}
	Output string
	Error  error
}

func testType(t *testing.T, tests []typeTest) {
	var buf bytes.Buffer
	for n, test := range tests {
		err := test.Input.Decode(test.Params, test.Data)
		if err != test.Error {
			if !reflect.DeepEqual(err, test.Error) {
				t.Errorf("test %d: expecting error %s, got %s", n+1, test.Error, err)
				continue
			}
		}
		if !reflect.DeepEqual(test.Input, test.Match) {
			t.Errorf("test %d: input does not match expected", n+1)
			continue
		}
		test.Match.Encode(&buf)
		if str := buf.String(); str != test.Output {
			t.Errorf("test %d: expecting output string %q, got %q", n+1, test.Output, str)
		}
		buf.Reset()
	}
}

func TestBinary(t *testing.T) {
	testType(t, []typeTest{
		{
			Input: &Binary{},
			Match: &Binary{},
			Error: ErrInvalidEncoding,
		},
		{
			Params: map[string]string{"ENCODING": "BASE64"},
			Data:   "MTIzNDU=",
			Input:  &Binary{},
			Match:  &Binary{'1', '2', '3', '4', '5'},
			Output: "MTIzNDU=",
		},
	})
}

func TestBoolean(t *testing.T) {
	tr := new(Boolean)
	fa := new(Boolean)
	*tr = true
	testType(t, []typeTest{
		{
			Data:   "False",
			Input:  fa,
			Match:  fa,
			Output: "FALSE",
		},
		{
			Data:   "true",
			Input:  tr,
			Match:  tr,
			Output: "TRUE",
		},
		{
			Data:   "HotDog",
			Input:  fa,
			Match:  fa,
			Output: "FALSE",
			Error: &strconv.NumError{
				Func: "ParseBool",
				Num:  "HotDog",
				Err:  strconv.ErrSyntax,
			},
		},
	})
}
