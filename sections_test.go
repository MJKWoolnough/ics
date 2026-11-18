package ics

import (
	"errors"
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
