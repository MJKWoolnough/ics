package ics

import (
	"bytes"
	"io"
	"reflect"
	"testing"
	"time"
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

var emptyMap = make(map[string]string)

func testType(t *testing.T, tests []typeTest) {
	var buf bytes.Buffer
	for n, test := range tests {
		if test.Params == nil {
			test.Params = emptyMap
		}
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
			Error:  ErrInvalidBoolean,
		},
	})
}

func TestDate(t *testing.T) {
	testType(t, []typeTest{
		{
			Data:   "20011225",
			Input:  &Date{},
			Match:  &Date{time.Date(2001, 12, 25, 0, 0, 0, 0, time.UTC)},
			Output: "20011225",
		},
		{
			Data:   "20081111",
			Input:  &Date{},
			Match:  &Date{time.Date(2008, 11, 11, 0, 0, 0, 0, time.UTC)},
			Output: "20081111",
		},
	})
}

func TestDateTime(t *testing.T) {
	l, err := time.LoadLocation("America/New_York")
	if err != nil {
		t.Errorf("unexpected error: %s", err)
		return
	}
	testType(t, []typeTest{
		{
			Data:   "20011225T131415",
			Input:  &DateTime{},
			Match:  &DateTime{time.Date(2001, 12, 25, 13, 14, 15, 0, time.Local)},
			Output: "20011225T131415",
		},
		{
			Data:   "20011225T131415Z",
			Input:  &DateTime{},
			Match:  &DateTime{time.Date(2001, 12, 25, 13, 14, 15, 0, time.UTC)},
			Output: "20011225T131415Z",
		},
		{
			Params: map[string]string{"TZID": "America/New_York"},
			Data:   "20011225T131415",
			Input:  &DateTime{},
			Match:  &DateTime{time.Date(2001, 12, 25, 13, 14, 15, 0, l)},
			Output: "20011225T131415",
		},
	})
}