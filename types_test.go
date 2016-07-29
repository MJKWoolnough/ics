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
		if test.Error == nil {
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
			Data:  "HotDog",
			Input: fa,
			Error: ErrInvalidBoolean,
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

func TestDuration(t *testing.T) {
	testType(t, []typeTest{
		{
			Data:   "P1W",
			Input:  &Duration{},
			Match:  &Duration{Weeks: 1},
			Output: "P1W",
		},
		{
			Data:   "P1D",
			Input:  &Duration{},
			Match:  &Duration{Days: 1},
			Output: "P1D",
		},
		{
			Data:   "PT1H",
			Input:  &Duration{},
			Match:  &Duration{Hours: 1},
			Output: "PT1H",
		},
		{
			Data:   "PT1M",
			Input:  &Duration{},
			Match:  &Duration{Minutes: 1},
			Output: "PT1M",
		},
		{
			Data:   "PT1S",
			Input:  &Duration{},
			Match:  &Duration{Seconds: 1},
			Output: "PT1S",
		},
		{
			Data:   "P25W",
			Input:  &Duration{},
			Match:  &Duration{Weeks: 25},
			Output: "P25W",
		},
		{
			Data:  "P25W1D",
			Input: &Duration{},
			Error: ErrInvalidDuration,
		},
		{
			Data:  "P25G",
			Input: &Duration{},
			Error: ErrInvalidDuration,
		},
		{
			Data:   "P1DT2H3M4S",
			Input:  &Duration{},
			Match:  &Duration{Days: 1, Hours: 2, Minutes: 3, Seconds: 4},
			Output: "P1DT2H3M4S",
		},
		{
			Data:   "PT0S",
			Input:  &Duration{},
			Match:  &Duration{},
			Output: "PT0S",
		},
		{
			Data:   "+P1DT2H3M4S",
			Input:  &Duration{},
			Match:  &Duration{Days: 1, Hours: 2, Minutes: 3, Seconds: 4},
			Output: "P1DT2H3M4S",
		},
		{
			Data:   "-P1DT2H3M4S",
			Input:  &Duration{},
			Match:  &Duration{Negative: true, Days: 1, Hours: 2, Minutes: 3, Seconds: 4},
			Output: "-P1DT2H3M4S",
		},
	})
}

func TestPeriod(t *testing.T) {
	testType(t, []typeTest{
		{
			Data:  "19970101T180000Z/19970102T070000Z",
			Input: &Period{},
			Match: &Period{
				Start: DateTime{Time: time.Date(1997, 1, 1, 18, 0, 0, 0, time.UTC)},
				End:   DateTime{Time: time.Date(1997, 1, 2, 7, 0, 0, 0, time.UTC)},
			},
			Output: "19970101T180000Z/19970102T070000Z",
		},
		{
			Data:  "19970101T180000Z/PT5H30M",
			Input: &Period{},
			Match: &Period{
				Start:    DateTime{Time: time.Date(1997, 1, 1, 18, 0, 0, 0, time.UTC)},
				Duration: Duration{Hours: 5, Minutes: 30},
			},
			Output: "19970101T180000Z/PT5H30M",
		},
	})
}

func TestText(t *testing.T) {
	newText := func(s string) *Text {
		nt := Text(s)
		return &nt
	}
	testType(t, []typeTest{
		{
			Data:   "",
			Input:  new(Text),
			Match:  newText(""),
			Output: "",
		},
		{
			Data:   "Jackdaws love my big Sphinx of quartz",
			Input:  new(Text),
			Match:  newText("Jackdaws love my big Sphinx of quartz"),
			Output: "Jackdaws love my big Sphinx of quartz",
		},
		{
			Data:   "Project XYZ Final Review\\nConference Room - 3B\\nCome Prepared.",
			Input:  new(Text),
			Match:  newText("Project XYZ Final Review\nConference Room - 3B\nCome Prepared."),
			Output: "Project XYZ Final Review\\nConference Room - 3B\\nCome Prepared.",
		},
		{
			Data:   "\\\\\\;\\,\\N\\n",
			Input:  new(Text),
			Match:  newText("\\;,\n\n"),
			Output: "\\\\\\;\\,\\n\\n",
		},
		{
			Data:  ";",
			Input: new(Text),
			Error: ErrInvalidText,
		},
		{
			Data:  ",",
			Input: new(Text),
			Error: ErrInvalidText,
		},
		{
			Data:  "\\",
			Input: new(Text),
			Error: ErrInvalidText,
		},
		{
			Data:  "\"",
			Input: new(Text),
			Error: ErrInvalidText,
		},
		{
			Data:  "\\g",
			Input: new(Text),
			Error: ErrInvalidText,
		},
		{
			Data:   "^n",
			Input:  new(Text),
			Match:  newText("\n"),
			Output: "\\n",
		},
		{
			Data:   "^^",
			Input:  new(Text),
			Match:  newText("^"),
			Output: "^^",
		},
		{
			Data:   "^'",
			Input:  new(Text),
			Match:  newText("\""),
			Output: "^'",
		},
	})
}

func TestTime(t *testing.T) {
	l, err := time.LoadLocation("America/New_York")
	if err != nil {
		t.Errorf("unexpected error: %s", err)
		return
	}
	testType(t, []typeTest{
		{
			Data:   "010203",
			Input:  &Time{},
			Match:  &Time{Time: time.Date(0, 1, 1, 1, 2, 3, 0, time.Local)},
			Output: "010203",
		},
		{
			Data:   "010203Z",
			Input:  &Time{},
			Match:  &Time{Time: time.Date(0, 1, 1, 1, 2, 3, 0, time.UTC)},
			Output: "010203Z",
		},
		{
			Params: map[string]string{"TZID": "America/New_York"},
			Data:   "010203",
			Input:  &Time{},
			Match:  &Time{Time: time.Date(0, 1, 1, 1, 2, 3, 0, l)},
			Output: "010203",
		},
	})
}
