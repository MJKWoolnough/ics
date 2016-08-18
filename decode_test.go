package ics

import (
	"io"
	"reflect"
	"strings"
	"testing"
)

func TestDecode(t *testing.T) {
	tests := []struct {
		Input  string
		Output *Calendar
		Error  error
	}{
		{
			Error: io.ErrUnexpectedEOF,
		},
		{
			Input: "BEGIN:VCALENDAR\r\nPRODID:TestDecode\r\nVERSION:2.0\r\nEND:VCALENDAR\r\n",
			Output: &Calendar{
				ProdID:  "TestDecode",
				Version: "2.0",
			},
		},
	}

	for n, test := range tests {
		c, err := Decode(strings.NewReader(test.Input))
		if err != test.Error {
			t.Errorf("test %d: expecting error %q, got %q", n+1, test.Error, err)
		} else if !reflect.DeepEqual(c, test.Output) {
			t.Errorf("test %d: expecting calendar %v, got %v", n+1, test.Output, c)
		}
	}
}
