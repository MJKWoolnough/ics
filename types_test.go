package ics

import (
	"bytes"
	"io"
	"reflect"
	"testing"
)

type typeTest struct {
	Params map[string]string
	Data   string
	Input  interface {
		Decode(map[string]string, string) error
	}
	Output interface {
		Encode(io.Writer)
	}
	Error error
}

func testType(t *testing.T, tests []typeTest) {
	var buf bytes.Buffer
	for n, test := range tests {
		err := test.Input.Decode(test.Params, test.Data)
		if err != test.Error {
			t.Errorf("test %d: expecting error %s, got %s", n+1, test.Error, err)
			continue
		}
		if !reflect.DeepEqual(test.Input, test.Output) {
			t.Errorf("test %d: input and output do not match", n+1)
			continue
		}
		test.Output.Encode(&buf)
		if str := buf.String(); str != test.Data {
			t.Errorf("test %d: expecting output string %q, got %q", test.Data, str)
		}
	}
}

func TestBinary(t *testing.T) {
	testType(t, []typeTest{
		{
			Input:  &Binary{},
			Output: &Binary{},
			Error:  ErrInvalidEncoding,
		},
		{
			Params: map[string]string{"ENCODING": "BASE64"},
			Data:   "MTIzNDU=",
			Input:  &Binary{},
			Output: &Binary{'1', '2', '3', '4', '5'},
		},
	})
}
