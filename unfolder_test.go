package ics

import (
	"bytes"
	"io"
	"strings"
	"testing"
)

func TestUnfolder(t *testing.T) {
	tests := []struct {
		Input, Output string
	}{
		{"A", "A"},
		{"A\nB", "A\nB"},
		{"A\r\n B", "AB"},
		{"ABCDEFGHIJKL\r\n MNOP\r\n QRSTUV\r\nWXY\r\n Z", "ABCDEFGHIJKLMNOPQRSTUV\r\nWXYZ"},
	}
	var buf bytes.Buffer
	for n, test := range tests {
		io.Copy(&buf, &unfolder{r: strings.NewReader(test.Input)})
		if str := buf.String(); str != test.Output {
			t.Errorf("test %d.1: expecting output %q, got %q", n, test.Output, str)
		}
		buf.Reset()
		var b [1]byte
		u := &unfolder{r: strings.NewReader(test.Input)}
		for {
			if _, err := u.Read(b[:]); err == io.EOF {
				break
			}
			buf.Write(b[:])
		}
		if str := buf.String(); str != test.Output {
			t.Errorf("test %d.2: expecting output %q, got %q", n, test.Output, str)
		}
		buf.Reset()
	}
}
