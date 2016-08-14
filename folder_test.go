package ics

import (
	"bytes"
	"io"
	"strings"
	"testing"
)

func TestFolder(t *testing.T) {
	tests := []struct {
		Input, Output string
	}{
		{"A", "A"},
		{"AB", "AB"},
		{"ABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZ", "ABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVW\r\n XYZ"},
		{"ABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZ", "ABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVW\r\n XYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRST\r\n UVWXYZ"},
		{"ABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZ\r\nABCDEFGHIJKLMNOPQRSTUVWXYZ", "ABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZ\r\nABCDEFGHIJKLMNOPQRSTUVWXYZ"},
	}
	var buf bytes.Buffer
	for n, test := range tests {
		io.Copy(&folder{w: &buf}, strings.NewReader(test.Input))
		if str := buf.String(); str != test.Output {
			t.Errorf("test %d.1: expecting output %q, got %q", n+1, test.Output, str)
		}
		buf.Reset()
		var b [1]byte
		f := &folder{w: &buf}
		for i := 0; i < len(test.Input); i++ {
			b[0] = test.Input[i]
			f.Write(b[:])
		}
		if str := buf.String(); str != test.Output {
			t.Errorf("test %d.2: expecting output %q, got %q", n+1, test.Output, str)
		}
		buf.Reset()
	}
}
