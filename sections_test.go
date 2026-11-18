package ics

import (
	"errors"
	"io"
	"strings"
	"testing"
)

type sectionTest struct {
	Input, Output string
	Error         error
}

func testSection[T any, PT interface {
	*T
	section
}](t *testing.T, tests []sectionTest) {
	for n, test := range tests {
		var section PT = new(T)

		if err := section.decode(newTokeniser(strings.NewReader(test.Input))); !errors.Is(err, test.Error) {
			t.Errorf("test %d: expecting error %v, got %v", n+1, test.Error, err)
		} else {
			var buf strings.Builder

			section.encode(&buf)

			if out := buf.String(); out != test.Output {
				t.Errorf("test %d: expecting output %q, got %q", n+1, test.Output, out)
			}
		}
	}
}

func TestCalendar(t *testing.T) {
	testSection[Calendar](t, []sectionTest{
		{
			Error:  io.ErrUnexpectedEOF,
			Output: "BEGIN:VCALENDAR\r\nVERSION:\r\nPRODID:\r\nEND:VCALENDAR\r\n",
		},
		{
			Input:  "+",
			Error:  ErrInvalidContentLine,
			Output: "BEGIN:VCALENDAR\r\nVERSION:\r\nPRODID:\r\nEND:VCALENDAR\r\n",
		},
		{
			Input:  "END:VCALENDAR",
			Error:  ErrMissingRequired,
			Output: "BEGIN:VCALENDAR\r\nVERSION:\r\nPRODID:\r\nEND:VCALENDAR\r\n",
		},
		{
			Input:  "VERSION:1.0\r\nEND:VCALENDAR",
			Error:  ErrMissingRequired,
			Output: "BEGIN:VCALENDAR\r\nVERSION:1.0\r\nPRODID:\r\nEND:VCALENDAR\r\n",
		},
		{
			Input:  "PRODID:1\r\nEND:VCALENDAR",
			Error:  ErrMissingRequired,
			Output: "BEGIN:VCALENDAR\r\nVERSION:\r\nPRODID:1\r\nEND:VCALENDAR\r\n",
		},
		{
			Input:  "VERSION:1.0\r\nPRODID:1\r\nEND:VCALENDAR",
			Output: "BEGIN:VCALENDAR\r\nVERSION:1.0\r\nPRODID:1\r\nEND:VCALENDAR\r\n",
		},
		{
			Input:  "VERSION:1.0\r\nPRODID:1\r\nBEGIN:VEVENT\r\nDTSTAMP:19990404T020000\r\nUID:2\r\nEND:VEVENT\r\nEND:VCALENDAR",
			Output: "BEGIN:VCALENDAR\r\nVERSION:1.0\r\nPRODID:1\r\nBEGIN:VEVENT\r\nDTSTAMP:19990404T020000\r\nUID:2\r\nEND:VEVENT\r\nEND:VCALENDAR\r\n",
		},
		{
			Input:  "VERSION:1.0\r\nPRODID:1\r\nBEGIN:VEVENT\r\nDTSTAMP:19990404T020000\r\nEND:VEVENT\r\nEND:VCALENDAR",
			Error:  ErrMissingRequired,
			Output: "BEGIN:VCALENDAR\r\nVERSION:1.0\r\nPRODID:1\r\nEND:VCALENDAR\r\n",
		},
		{
			Input:  "VERSION:1.0\r\nPRODID:1\r\nBEGIN:VEVENT\r\nDTSTAMP:19990404T020000\r\nUID:2\r\nEND:VEVENT\r\nBEGIN:VEVENT\r\nDTSTAMP:19990404T020000\r\nUID:3\r\nEND:VEVENT\r\nEND:VCALENDAR",
			Output: "BEGIN:VCALENDAR\r\nVERSION:1.0\r\nPRODID:1\r\nBEGIN:VEVENT\r\nDTSTAMP:19990404T020000\r\nUID:2\r\nEND:VEVENT\r\nBEGIN:VEVENT\r\nDTSTAMP:19990404T020000\r\nUID:3\r\nEND:VEVENT\r\nEND:VCALENDAR\r\n",
		},
	})
}
