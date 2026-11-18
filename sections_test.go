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
		{ // 1
			Error:  io.ErrUnexpectedEOF,
			Output: "BEGIN:VCALENDAR\r\nVERSION:\r\nPRODID:\r\nEND:VCALENDAR\r\n",
		},
		{ // 2
			Input:  "+",
			Error:  ErrInvalidContentLine,
			Output: "BEGIN:VCALENDAR\r\nVERSION:\r\nPRODID:\r\nEND:VCALENDAR\r\n",
		},
		{ // 3
			Input:  "END:VCALENDAR",
			Error:  ErrMissingRequired,
			Output: "BEGIN:VCALENDAR\r\nVERSION:\r\nPRODID:\r\nEND:VCALENDAR\r\n",
		},
		{ // 4
			Input:  "VERSION:1.0\r\nEND:VCALENDAR",
			Error:  ErrMissingRequired,
			Output: "BEGIN:VCALENDAR\r\nVERSION:1.0\r\nPRODID:\r\nEND:VCALENDAR\r\n",
		},
		{ // 5
			Input:  "PRODID:1\r\nEND:VCALENDAR",
			Error:  ErrMissingRequired,
			Output: "BEGIN:VCALENDAR\r\nVERSION:\r\nPRODID:1\r\nEND:VCALENDAR\r\n",
		},
		{ // 6
			Input:  "VERSION:1.0\r\nPRODID:1\r\nEND:VCALENDAR",
			Output: "BEGIN:VCALENDAR\r\nVERSION:1.0\r\nPRODID:1\r\nEND:VCALENDAR\r\n",
		},
		{ // 7
			Input:  "VERSION:1.0\r\nPRODID:1\r\nBEGIN:VEVENT\r\nDTSTAMP:19990404T020000\r\nUID:2\r\nEND:VEVENT\r\nEND:VCALENDAR",
			Output: "BEGIN:VCALENDAR\r\nVERSION:1.0\r\nPRODID:1\r\nBEGIN:VEVENT\r\nDTSTAMP:19990404T020000\r\nUID:2\r\nEND:VEVENT\r\nEND:VCALENDAR\r\n",
		},
		{ // 8
			Input:  "VERSION:1.0\r\nPRODID:1\r\nBEGIN:VEVENT\r\nDTSTAMP:19990404T020000\r\nEND:VEVENT\r\nEND:VCALENDAR",
			Error:  ErrMissingRequired,
			Output: "BEGIN:VCALENDAR\r\nVERSION:1.0\r\nPRODID:1\r\nEND:VCALENDAR\r\n",
		},
		{ // 9
			Input:  "VERSION:1.0\r\nPRODID:1\r\nBEGIN:VEVENT\r\nDTSTAMP:19990404T020000\r\nUID:2\r\nEND:VEVENT\r\nBEGIN:VEVENT\r\nDTSTAMP:19990404T020000\r\nUID:3\r\nEND:VEVENT\r\nEND:VCALENDAR",
			Output: "BEGIN:VCALENDAR\r\nVERSION:1.0\r\nPRODID:1\r\nBEGIN:VEVENT\r\nDTSTAMP:19990404T020000\r\nUID:2\r\nEND:VEVENT\r\nBEGIN:VEVENT\r\nDTSTAMP:19990404T020000\r\nUID:3\r\nEND:VEVENT\r\nEND:VCALENDAR\r\n",
		},
		{ // 10
			Input:  "VERSION:1.0\r\nPRODID:1\r\nBEGIN:VTODO\r\nDTSTAMP:19990404T020000\r\nUID:2\r\nEND:VTODO\r\nEND:VCALENDAR",
			Output: "BEGIN:VCALENDAR\r\nVERSION:1.0\r\nPRODID:1\r\nBEGIN:VTODO\r\nDTSTAMP:19990404T020000\r\nUID:2\r\nEND:VTODO\r\nEND:VCALENDAR\r\n",
		},
		{ // 11
			Input:  "VERSION:1.0\r\nPRODID:1\r\nBEGIN:VTODO\r\nDTSTAMP:19990404T020000\r\nEND:VTODO\r\nEND:VCALENDAR",
			Error:  ErrMissingRequired,
			Output: "BEGIN:VCALENDAR\r\nVERSION:1.0\r\nPRODID:1\r\nEND:VCALENDAR\r\n",
		},
		{ // 12
			Input:  "VERSION:1.0\r\nPRODID:1\r\nBEGIN:VTODO\r\nDTSTAMP:19990404T020000\r\nUID:2\r\nEND:VTODO\r\nBEGIN:VTODO\r\nDTSTAMP:19990404T020000\r\nUID:3\r\nEND:VTODO\r\nEND:VCALENDAR",
			Output: "BEGIN:VCALENDAR\r\nVERSION:1.0\r\nPRODID:1\r\nBEGIN:VTODO\r\nDTSTAMP:19990404T020000\r\nUID:2\r\nEND:VTODO\r\nBEGIN:VTODO\r\nDTSTAMP:19990404T020000\r\nUID:3\r\nEND:VTODO\r\nEND:VCALENDAR\r\n",
		},
		{ // 13
			Input:  "VERSION:1.0\r\nPRODID:1\r\nBEGIN:VJOURNAL\r\nDTSTAMP:19990404T020000\r\nUID:2\r\nEND:VJOURNAL\r\nEND:VCALENDAR",
			Output: "BEGIN:VCALENDAR\r\nVERSION:1.0\r\nPRODID:1\r\nBEGIN:VJOURNAL\r\nDTSTAMP:19990404T020000\r\nUID:2\r\nEND:VJOURNAL\r\nEND:VCALENDAR\r\n",
		},
		{ // 14
			Input:  "VERSION:1.0\r\nPRODID:1\r\nBEGIN:VJOURNAL\r\nDTSTAMP:19990404T020000\r\nEND:VJOURNAL\r\nEND:VCALENDAR",
			Error:  ErrMissingRequired,
			Output: "BEGIN:VCALENDAR\r\nVERSION:1.0\r\nPRODID:1\r\nEND:VCALENDAR\r\n",
		},
		{ // 15
			Input:  "VERSION:1.0\r\nPRODID:1\r\nBEGIN:VJOURNAL\r\nDTSTAMP:19990404T020000\r\nUID:2\r\nEND:VJOURNAL\r\nBEGIN:VJOURNAL\r\nDTSTAMP:19990404T020000\r\nUID:3\r\nEND:VJOURNAL\r\nEND:VCALENDAR",
			Output: "BEGIN:VCALENDAR\r\nVERSION:1.0\r\nPRODID:1\r\nBEGIN:VJOURNAL\r\nDTSTAMP:19990404T020000\r\nUID:2\r\nEND:VJOURNAL\r\nBEGIN:VJOURNAL\r\nDTSTAMP:19990404T020000\r\nUID:3\r\nEND:VJOURNAL\r\nEND:VCALENDAR\r\n",
		},
	})
}
